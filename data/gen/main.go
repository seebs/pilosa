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
	names:     []string{"", "", ""},
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
	for i, typ := range iot.typeNames {
		name := iot.names[i]
		typeName := iot.typeNames[i]
		switch typ {
		case "bit":
			typeName = dtBits
		case "other":
			name = dtName
			typeName = dtName
		default:
			name = iot.names[i]
			typeName = iot.typeNames[i]
		}
		roName := name
		roTypeName := typeName
		if typ == "other" {
			roName = "ReadOnly" + name
			roTypeName = "ReadOnly" + typeName
		}
		if newIot.plurals[i] {
			name += "s"
			typeName = "[]" + typeName
			roTypeName = "[]" + roTypeName
		}
		newIot.names[i] = name
		newIot.typeNames[i] = typeName
		newIot.readOnlyNames[i] = roName
		newIot.readOnlyTypeNames[i] = roTypeName
	}
	return newIot
}

type opData struct {
	base       string
	full       string
	viewUpdate string
	operands   *ioTypes
	results    *ioTypes
}

func (opd *opData) String() string {
	return fmt.Sprintf("%s [%s]: %s op, inputs %s, outputs %s",
		opd.base, opd.full, opd.viewUpdate, opd.operands, opd.results)
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
		details := opTypeParser.FindStringSubmatch(op)
		if details == nil || details[0] == "" {
			return nil, fmt.Errorf("name '%s' does not match expected pattern", op)
		}
		operands := extractTypes(details[2])
		results := extractTypes(details[4])
		opDatas[i] = opData{base: op, full: "OpType" + op, viewUpdate: details[1], operands: operands, results: results}
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

var standardOpNames = [OpTypeMax]string{
`)
	for _, op := range opDatas {
		Printf("\t%s: \"%s\",\n", op.full, op.base)
	}
	Printf(`}

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

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunction%s interface {
	%sOpType() OpType
}

`, dt.name, dt.name, dt.name, dt.name)
	newline := false
	opNames := make(map[string]string, len(opDatas))
	for _, op := range opDatas {
		if newline {
			Printf("\n")
		}
		operands := op.operands.Interpolate(dt.name, dt.bitsIn)
		var results *ioTypes
		if op.results != nil {
			results = op.results.Interpolate(dt.name, dt.bitsOut)
			opNames[op.full] = fmt.Sprintf("Op%s%s%sGives%s", dt.name, op.viewUpdate, operands.Names(), results.Names())
		} else {
			results = defaultResultIoTypes.Interpolate(dt.name, dt.bitsOut)
			opNames[op.full] = fmt.Sprintf("Op%s%s%s", dt.name, op.viewUpdate, operands.Names())
		}
		var resTypes string
		if op.viewUpdate == "View" {
			resTypes = results.ReadOnlyTypes()
		} else {
			resTypes = results.Types()
		}
		Printf("type %s func(%s) (%s)\n", opNames[op.full], operands.ReadOnlyTypes(), resTypes)

		Printf("func (%s) %sOpType() OpType { return %s }\n", opNames[op.full], dt.name, op.full)

		// dummy var to allow us to build a type-table.
		Printf("var zero%s %s\n", opNames[op.full], opNames[op.full])
		newline = true
	}
	Printf(`
// OpType to reflect.Type lookup table
var lookup%sFunctionTypes = [OpTypeMax]OpFunction%s {
`, dt.name, dt.name)
	for _, op := range opDatas {
		Printf("\t%s: zero%s,\n", op.full, opNames[op.full])
	}
	Printf("}\n")
	return nil
}

var nonAlpha = regexp.MustCompile("[^[:alnum:]]")

// Gen creates the op function declarations, and corresponding methods
// to make them satisfy OpFunction, for a given type and series of ops.
// We could also parse the ops from optypes.go, but this seems like a lot
// of work.
func main() {
	typeNames := flag.String("types", "Container:uint16:int,Bitmap:uint64:int64", "types, specified as comma-separated typename:bitIn:bitOut, such as Container:uint16:,Bitmap:uint64:int64")
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
