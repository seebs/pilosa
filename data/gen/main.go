package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type dataType struct {
	name      string
	nameLower string
	bitsIn    string
	bitsOut   string
}

// ioType represents a function input or output, which has a name
// we describe it as, and a type name, which may not be the same.
// for instance, "Bits" gets "[]int" plus a bit width if applicable,
// and "Others" might turn into "[]Bitmap" or "[]Container".
type ioTypes struct {
	names             []string
	readOnlyNames     []string
	typeNames         []string
	readOnlyTypeNames []string
	plurals           []bool
}

var defaultResultIoTypes = &ioTypes{
	names:     []string{"Bool", "Bit", "Other"},
	typeNames: []string{"bool", "bit", "other"},
}

func newIoTypes(n int) *ioTypes {
	return &ioTypes{names: make([]string, n), typeNames: make([]string, n), readOnlyNames: make([]string, n), readOnlyTypeNames: make([]string, n), plurals: make([]bool, n)}
}

func (iot *ioTypes) Names() string {
	return strings.Join(iot.names, "")
}

func (iot *ioTypes) Types() string {
	return strings.Join(iot.typeNames, ", ")
}

func (iot *ioTypes) ReadOnlyNames() string {
	return strings.Join(iot.readOnlyNames, "")
}

func (iot *ioTypes) ReadOnlyTypes() string {
	return strings.Join(iot.readOnlyTypeNames, ", ")
}

func (iot *ioTypes) Any() bool {
	return len(iot.names) > 0
}

// NamedArgs gives "in1 type1, in2 type2, ..." for the arg list.
func (iot *ioTypes) NamedArgs() string {
	args := make([]string, len(iot.typeNames))
	for i := range iot.typeNames {
		args[i] = fmt.Sprintf("in%d %s", i+1, iot.readOnlyTypeNames[i])
	}
	return strings.Join(args, ", ")
}

// ArgNames gives the names used by NamedArgs.
func (iot *ioTypes) ArgNames() string {
	args := make([]string, len(iot.typeNames))
	for i := range iot.typeNames {
		args[i] = fmt.Sprintf("in%d", i+1)
	}
	return strings.Join(args, ", ")
}

func (iot *ioTypes) String() string {
	if iot == nil {
		return "[nil ioTypes]"
	}
	return fmt.Sprintf("%s/%s", iot.Names(), iot.Types())
}

// Extrapolate, replacing "other" with the provided name, and
// "bit" with the corresponding int type. It also creates parallel
// names prefixed with "ReadOnly".
func (iot *ioTypes) Interpolate(dtName, dtBits string) *ioTypes {
	if iot == nil {
		return newIoTypes(0)
	}
	newIot := newIoTypes(len(iot.names))
	copy(newIot.plurals, iot.plurals)
	newIot.typeNames = newIot.typeNames[:0]
	newIot.readOnlyTypeNames = newIot.readOnlyTypeNames[:0]
	for i, typ := range iot.names {
		name := iot.names[i]
		typeNames := []string{iot.typeNames[i]}
		switch typ {
		case "Bit":
			typeNames[0] = dtBits
		case "Other":
			name = dtName
			typeNames[0] = dtName
		case "Range":
			// a range is done as two bits, specifying an
			// inclusive range. it has to be inclusive because
			// you can't represent a value greater than the
			// maximum value...
			typeNames = []string{dtBits, dtBits}
		case "Writer":
			typeNames[0] = "io.Writer"
		}
		roName := name
		roTypeNames := make([]string, len(typeNames))
		copy(roTypeNames, typeNames)
		if typ == "Other" {
			roName = "ReadOnly" + name
			for j, typeName := range typeNames {
				roTypeNames[j] = "ReadOnly" + typeName
			}
		}
		if newIot.plurals[i] {
			name += "s"
			for j, typeName := range typeNames {
				typeNames[j] = "[]" + typeName
				roTypeNames[j] = "[]" + roTypeNames[j]
			}
		}
		newIot.names[i] = name
		newIot.typeNames = append(newIot.typeNames, typeNames...)
		newIot.readOnlyNames[i] = roName
		newIot.readOnlyTypeNames = append(newIot.readOnlyTypeNames, roTypeNames...)
		if typ == "range" {
			fmt.Printf("new typeNames: %s\n", newIot.typeNames)
		}
	}
	return newIot
}

type opData struct {
	base            string
	full            string
	viewUpdate      string
	operands        *ioTypes
	results         *ioTypes
	typedDefaultOps map[string][]string
}

func (opd *opData) String() string {
	var defaults []string
	for k, v := range opd.typedDefaultOps {
		defaults = append(defaults, fmt.Sprintf("%s:%s", k, strings.Join(v, ",")))
	}
	return fmt.Sprintf("%s [%s]: %s op, inputs %s, outputs %s, defaults %s",
		opd.base, opd.full, opd.viewUpdate, opd.operands, opd.results,
		strings.Join(defaults, ", "))
}

var opTypeParser = regexp.MustCompile("^(View|Update)(.*?)(Gives([[:alpha:]]*))?$")
var wordParser = regexp.MustCompile("[A-Z][a-z]*")

// extractTypes takes "BytesInt" and makes a list of ioType objects.
func extractTypes(typeList string) *ioTypes {
	if typeList == "" {
		return nil
	}
	words := wordParser.FindAllString(typeList, -1)
	ioTypes := newIoTypes(len(words))
	for i, word := range words {
		if strings.HasSuffix(word, "s") {
			word = word[:len(word)-1]
			ioTypes.plurals[i] = true
		}
		ioTypes.names[i] = word
		word = strings.ToLower(word)
		ioTypes.typeNames[i] = word
	}
	return ioTypes
}

func extractOpNames(ops []string) ([]opData, error) {
	opDatas := make([]opData, len(ops))
	for i, op := range ops {
		op_and_defaults := strings.Split(op, ",")
		op, defaults := op_and_defaults[0], op_and_defaults[1:]
		details := opTypeParser.FindStringSubmatch(op)
		if details == nil || details[0] == "" {
			return nil, fmt.Errorf("name '%s' does not match expected pattern", op)
		}
		operands := extractTypes(details[2])
		results := extractTypes(details[4])
		opd := opData{base: op, full: "OpType" + op, viewUpdate: details[1], operands: operands, results: results, typedDefaultOps: make(map[string][]string)}
		if len(defaults) > 0 {
			for _, def := range defaults {
				typ_and_name := strings.Split(def, ":")
				if len(typ_and_name) != 2 {
					return nil, fmt.Errorf("%s should be type:name", def)
				}
				typ, name := typ_and_name[0], typ_and_name[1]
				opd.typedDefaultOps[typ] = append(opd.typedDefaultOps[typ], name)
			}
		}
		opDatas[i] = opd
	}
	return opDatas, nil
}

func extractDataTypes(typeSpecs []string) ([]dataType, error) {
	dataTypes := make([]dataType, len(typeSpecs))
	for i, spec := range typeSpecs {
		out := strings.Split(spec, ":")
		if len(out) != 3 {
			return nil, fmt.Errorf("datatype ('%s') must be name:bitwidth:bitwidth", spec)
		}
		name := out[0]
		nameLower := strings.ToLower(name)
		if oops := nonAlpha.FindString(nameLower); oops != "" {
			return nil, fmt.Errorf("type name '%s' does not lowercase to all alphanumeric characters ('%s')", name, oops)
		}
		dataTypes[i] = dataType{name: name, nameLower: nameLower, bitsIn: out[1], bitsOut: out[2]}
	}
	return dataTypes, nil
}

func writeData(dt dataType, opDatas []opData) error {
	fileName := fmt.Sprintf("%s_gen.go", dt.nameLower)
	tempName := fileName + ".tmp"
	file, err := os.Create(tempName)
	if err != nil {
		return fmt.Errorf("creating temp file '%s': %v\n", tempName, err)
	}
	defer file.Close()
	err = writeTypeOps(dt, opDatas, file)
	if err != nil {
		retErr := fmt.Errorf("writing ops: %v", err)
		err = os.Remove(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't remove '%s': %v\n", tempName, err)
		}
		return retErr
	}
	err = os.Rename(tempName, fileName)
	if err != nil {
		return fmt.Errorf("renaming '%s' to '%s': %v", tempName, fileName, err)
	}
	return nil
}

func writeOpData(opDatas []opData) error {
	fileName := "optype_gen.go"
	tempName := fileName + ".tmp"
	file, err := os.Create(tempName)
	if err != nil {
		return fmt.Errorf("creating temp file '%s': %v\n", tempName, err)
	}
	defer file.Close()
	err = writeOps(opDatas, file)
	if err != nil {
		retErr := fmt.Errorf("writing ops: %v", err)
		err = os.Remove(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't remove '%s': %v\n", tempName, err)
		}
		return retErr
	}
	err = os.Rename(tempName, fileName)
	if err != nil {
		return fmt.Errorf("renaming '%s' to '%s': %v", tempName, fileName, err)
	}
	return nil
}

func writeOps(opDatas []opData, w io.Writer) error {
	var Printf = func(format string, args ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, args...)
	}
	Printf(`package data

// GENERATED CODE, DO NOT EDIT
// Generated things from the OpTypes list (see gen/main.go)

const (
	%s = OpType(iota)
`, opDatas[0].full)

	for _, op := range opDatas[1:] {
		Printf("\t%s\n", op.full)
	}
	Printf(`
	OpTypeMax
)
`)
	return nil
}

func writeTypeOps(dt dataType, opDatas []opData, w io.Writer) error {
	var Printf = func(format string, args ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, args...)
	}
	Printf(`package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for %s (see gen/main.go). These are expressed
// as method signatures -- the %s they operate on is an implicit
// receiver not shown in the signature.

import (
	"io"
	"reflect"
)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunction%s interface {
	%sOpType() OpType
}

`, dt.name, dt.name, dt.name, dt.name)
	newline := false
	opNames := make(map[string]string, len(opDatas))
	operandLists := make(map[string]string, len(opDatas))
	for _, op := range opDatas {
		if newline {
			Printf("\n")
		}
		operands := op.operands.Interpolate(dt.name, dt.bitsIn)
		operandLists[op.full] = operands.NamedArgs()
		var results *ioTypes
		// Function Suffix: e.g., "ViewBitmap" for a view function taking
		// a Bitmap parameter
		var opFuncSuffix string
		if op.results != nil {
			results = op.results.Interpolate(dt.name, dt.bitsOut)
			opFuncSuffix = fmt.Sprintf("%s%sGives%s", op.viewUpdate, operands.Names(), results.Names())
		} else {
			results = defaultResultIoTypes.Interpolate(dt.name, dt.bitsOut)
			opFuncSuffix = fmt.Sprintf("%s%s", op.viewUpdate, operands.Names())
		}
		// opName, such as "BitmapViewBitmap", disambiguated with receiver type.
		opName := dt.name + opFuncSuffix
		opNames[op.full] = opName
		var resTypes string
		if op.viewUpdate == "View" {
			resTypes = results.ReadOnlyTypes()
		} else {
			resTypes = results.Types()
		}
		// guard the result types with () if they're plural
		if strings.Contains(resTypes, ",") {
			resTypes = "(" + resTypes + ")"
		}
		if resTypes != "" {
			resTypes = " " + resTypes
		}
		literalFuncType := fmt.Sprintf("func(%s)%s", operands.ReadOnlyTypes(), resTypes)
		Printf("type Op%s %s\n\n", opName, literalFuncType)

		Printf("func (Op%s) %sOpType() OpType { return %s }\n", opName, dt.name, op.full)

		Printf(`
func LookupOp%s(target ReadOnly%s, name string) Op%s {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "%s")
	if method.IsValid() {
		fn, _ := method.Interface().(%s)
		return Op%s(fn)
	}
	return nil
}
`, opName, dt.name, opName, opFuncSuffix, literalFuncType, opName)

		var readOnly string
		if op.viewUpdate == "View" {
			readOnly = "ReadOnly"
		}
		if defaults, ok := op.typedDefaultOps[dt.name]; ok {
			var typeList, argList, argNames, argComma string
			if operands.Any() {
				typeList = operands.ReadOnlyTypes()
				argList = operands.NamedArgs()
				argNames = operands.ArgNames()
				argComma = ", "
			}
			// create default op implementations
			for _, genericOpName := range defaults {
				opPrefix := ""
				// strip leading "Container" from
				if strings.HasPrefix(genericOpName, dt.name) {
					genericOpName = genericOpName[len(dt.name):]
					opPrefix = dt.name
				}
				Printf("\n// %s performs a default %s on a %s.\n", genericOpName, opName, dt.name)
				Printf("type interface%sHas%s interface {\n\t%s(%s)%s\n}\n\n", dt.name, genericOpName, genericOpName+opFuncSuffix, typeList, resTypes)
				Printf("func %s%s(target %s%s%s%s)%s {\n", opPrefix, genericOpName, readOnly, dt.name, argComma, argList, resTypes)
				Printf("	if target, ok := target.(interface%sHas%s); ok {\n", dt.name, genericOpName)
				Printf("		return target.%s(%s)\n", genericOpName+opFuncSuffix, argNames)
				Printf("	}\n")
				Printf("	return generic%s%s(target%s%s)\n", opPrefix, genericOpName, argComma, argNames)
				Printf(`}
`)
			}

		}
		newline = true
	}
	return nil
}

var nonAlpha = regexp.MustCompile("[^[:alnum:]]")

// Gen creates the op function declarations, and corresponding methods
// to make them satisfy OpFunction, for a given type and series of ops.
// We could also parse the ops from optypes.go, but this seems like a lot
// of work.
func main() {
	typeNames := flag.String("types", "Container:uint16:int,Bitmap:uint64:uint64", "types, specified as comma-separated typename:bitIn:bitOut, such as Container:uint16:,Bitmap:uint64:uint64")
	typeSource := flag.String("typelist", "optypes_list.txt", "file containing a list of optypes (specify \"\" to ignore)")
	flag.Parse()
	types := strings.Split(*typeNames, ",")
	if len(types) < 1 {
		fmt.Fprintf(os.Stderr, "types must be specified.\n")
		flag.Usage()
		os.Exit(1)
	}
	dataTypes, err := extractDataTypes(types)
	if err != nil {
		fmt.Fprintf(os.Stderr, "type specifications invalid: %v\n", err)
		os.Exit(1)
	}
	var opsGiven []string
	if *typeSource != "" {
		bits, err := ioutil.ReadFile(*typeSource)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading ops list: %v\n", err)
			os.Exit(1)
		}
		opsGiven = strings.Split(string(bits), "\n")
		j := 0
		// drop blank lines
		for i := range opsGiven {
			if opsGiven[i] != "" && !strings.HasPrefix(opsGiven[i], "//") {
				if i != j {
					opsGiven[j] = opsGiven[i]
				}
				j++
			}
		}
		opsGiven = opsGiven[:j]
	}
	opsGiven = append(opsGiven, flag.Args()...)
	if len(opsGiven) < 1 {
		fmt.Fprintf(os.Stderr, "usage: gen [-types=<typespecs>] [-typelist=filename] [<ops>]\n(typelist or ops list should be non-empty)\n")
		os.Exit(1)
	}
	opDatas, err := extractOpNames(opsGiven)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing op names: %v\n", err)
		os.Exit(1)
	}
	for _, data := range dataTypes {
		err = writeData(data, opDatas)
		if err != nil {
			fmt.Fprintf(os.Stderr, "writing ops for '%s': %v\n", data.name, err)
		}
	}
	err = writeOpData(opDatas)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing generic ops: %v\n", err)
	}
}
