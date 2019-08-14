package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type dataType struct {
	Name      string
	NameLower string
	BitsIn    string
	BitsOut   string
}

// ioType represents a function input or output, which has a name
// we describe it as, and a type name, which may not be the same.
// for instance, "Bits" gets "[]int" plus a bit width if applicable,
// and "Others" might turn into "[]Bitmap" or "[]Container".
type ioTypes struct {
	Names             []string
	ReadOnlyNames     []string
	TypeNames         []string
	ReadOnlyTypeNames []string
	Plurals           []bool
}

var defaultResultIoTypes = &ioTypes{
	Names:     []string{"Bool", "Bit", "Other"},
	TypeNames: []string{"bool", "bit", "other"},
}

func newIoTypes(n int) *ioTypes {
	return &ioTypes{Names: make([]string, n), TypeNames: make([]string, n), ReadOnlyNames: make([]string, n), ReadOnlyTypeNames: make([]string, n), Plurals: make([]bool, n)}
}

func (iot *ioTypes) NamesJoined() string {
	return strings.Join(iot.Names, "")
}

func (iot *ioTypes) Types() string {
	return strings.Join(iot.TypeNames, ", ")
}

func (iot *ioTypes) ReadOnlyNamesJoined() string {
	return strings.Join(iot.ReadOnlyNames, "")
}

func (iot *ioTypes) ReadOnlyTypes() string {
	return strings.Join(iot.ReadOnlyTypeNames, ", ")
}

func (iot *ioTypes) Any() bool {
	return len(iot.Names) > 0
}

// NamedArgs gives "prefix1 type1, prefix2 type2, ..." for a non-ReadOnly list.
func (iot *ioTypes) NamedArgs(prefix string) string {
	args := make([]string, len(iot.TypeNames))
	for i := range iot.TypeNames {
		args[i] = fmt.Sprintf("%s%d %s", prefix, i+1, iot.TypeNames[i])
	}
	return strings.Join(args, ", ")
}

// ReadOnlyNamedArgs gives "prefix1 type1, prefix2 type2, ..." for the arg list.
func (iot *ioTypes) ReadOnlyNamedArgs(prefix string) string {
	args := make([]string, len(iot.TypeNames))
	for i := range iot.TypeNames {
		args[i] = fmt.Sprintf("%s%d %s", prefix, i+1, iot.ReadOnlyTypeNames[i])
	}
	return strings.Join(args, ", ")
}

// ArgNames gives the Names used by NamedArgs.
func (iot *ioTypes) ArgNames(prefix string) string {
	args := make([]string, len(iot.TypeNames))
	for i := range iot.TypeNames {
		args[i] = fmt.Sprintf("%s%d", prefix, i+1)
	}
	return strings.Join(args, ", ")
}

func (iot *ioTypes) String() string {
	if iot == nil {
		return "[nil ioTypes]"
	}
	return fmt.Sprintf("%s/%s", iot.NamesJoined(), iot.Types())
}

// Extrapolate, replacing "other" with the provided name, and
// "bit" with the corresponding int type. It also creates parallel
// Names prefixed with "ReadOnly".
func (iot *ioTypes) Interpolate(dtName, dtBits string) *ioTypes {
	if iot == nil {
		return newIoTypes(0)
	}
	newIot := newIoTypes(len(iot.Names))
	copy(newIot.Plurals, iot.Plurals)
	newIot.TypeNames = newIot.TypeNames[:0]
	newIot.ReadOnlyTypeNames = newIot.ReadOnlyTypeNames[:0]
	for i, typ := range iot.Names {
		name := iot.Names[i]
		TypeNames := []string{iot.TypeNames[i]}
		switch typ {
		case "Bit":
			TypeNames[0] = dtBits
		case "Other":
			name = dtName
			TypeNames[0] = dtName
		case "Range":
			// a range is done as two bits, specifying an
			// inclusive range. it has to be inclusive because
			// you can't represent a value greater than the
			// maximum value...
			TypeNames = []string{dtBits, dtBits}
		case "Writer":
			TypeNames[0] = "io.Writer"
		}
		roName := name
		roTypeNames := make([]string, len(TypeNames))
		copy(roTypeNames, TypeNames)
		if typ == "Other" {
			roName = "ReadOnly" + name
			for j, typeName := range TypeNames {
				roTypeNames[j] = "ReadOnly" + typeName
			}
		}
		if newIot.Plurals[i] {
			name += "s"
			for j, typeName := range TypeNames {
				TypeNames[j] = "[]" + typeName
				roTypeNames[j] = "[]" + roTypeNames[j]
			}
		}
		newIot.Names[i] = name
		newIot.TypeNames = append(newIot.TypeNames, TypeNames...)
		newIot.ReadOnlyNames[i] = roName
		newIot.ReadOnlyTypeNames = append(newIot.ReadOnlyTypeNames, roTypeNames...)
		if typ == "range" {
			fmt.Printf("new TypeNames: %s\n", newIot.TypeNames)
		}
	}
	return newIot
}

type opData struct {
	Base            string
	Full            string
	ViewUpdate      string
	Operands        *ioTypes
	Results         *ioTypes
	TypedDefaultOps map[string][]string
}

func (opd *opData) String() string {
	var defaults []string
	for k, v := range opd.TypedDefaultOps {
		defaults = append(defaults, fmt.Sprintf("%s:%s", k, strings.Join(v, ",")))
	}
	return fmt.Sprintf("%s [%s]: %s op, inputs %s, outputs %s, defaults %s",
		opd.Base, opd.Full, opd.ViewUpdate, opd.Operands, opd.Results,
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
			ioTypes.Plurals[i] = true
		}
		ioTypes.Names[i] = word
		word = strings.ToLower(word)
		ioTypes.TypeNames[i] = word
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
		Operands := extractTypes(details[2])
		Results := extractTypes(details[4])
		opd := opData{Base: op, Full: "OpType" + op, ViewUpdate: details[1], Operands: Operands, Results: Results, TypedDefaultOps: make(map[string][]string)}
		if len(defaults) > 0 {
			for _, def := range defaults {
				typ_and_name := strings.Split(def, ":")
				if len(typ_and_name) != 2 {
					return nil, fmt.Errorf("%s should be type:name", def)
				}
				typ, name := typ_and_name[0], typ_and_name[1]
				opd.TypedDefaultOps[typ] = append(opd.TypedDefaultOps[typ], name)
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
		NameLower := strings.ToLower(name)
		if oops := nonAlpha.FindString(NameLower); oops != "" {
			return nil, fmt.Errorf("type name '%s' does not lowercase to all alphanumeric characters ('%s')", name, oops)
		}
		dataTypes[i] = dataType{Name: name, NameLower: NameLower, BitsIn: out[1], BitsOut: out[2]}
	}
	return dataTypes, nil
}

// sometimes we want to split out a possibly-optional prefix of a name,
// like when looking for functions named `AnyViewGivesBool` to satisfy
// `ContainerAnyViewGivesBool`
type prefixedName struct {
	Prefix string
	Name   string
}

type interpolatedOp struct {
	Op              opData
	ReadOnly        string
	OpName          string
	Results         *ioTypes
	Operands        *ioTypes
	ResTypes        string
	OpFuncSuffix    string
	LiteralFuncType string
	// if there's default-implementations to check for, we build a list
	// of function args, etc, to pass around.
	Defaults []prefixedName
	TypeList string
	ArgNames string
	ArgList  string
	ArgComma string
	ResList  string
	ResNames string
}

func writeData(dt dataType, opDatas []opData, tmpl *template.Template) error {
	fileName := fmt.Sprintf("%s_gen.go", dt.NameLower)
	tempName := fileName + ".tmp"
	file, err := os.Create(tempName)
	if err != nil {
		return fmt.Errorf("creating temp file '%s': %v\n", tempName, err)
	}
	defer file.Close()
	forTemplate := struct {
		Data dataType
		Ops  []interpolatedOp
	}{
		Data: dt,
		Ops:  make([]interpolatedOp, len(opDatas)),
	}
	// Populate the list and interpolate the data type into op names in some places,
	// generate argument lists, etcetera.
	for i, op := range opDatas {
		interp := interpolatedOp{
			Op:       op,
			Operands: op.Operands.Interpolate(dt.Name, dt.BitsIn),
		}
		if op.Results != nil {
			interp.Results = op.Results.Interpolate(dt.Name, dt.BitsOut)
			interp.OpFuncSuffix = fmt.Sprintf("%s%sGives%s", op.ViewUpdate, interp.Operands.NamesJoined(), interp.Results.NamesJoined())
		} else {
			interp.Results = defaultResultIoTypes.Interpolate(dt.Name, dt.BitsOut)
			interp.OpFuncSuffix = fmt.Sprintf("%s%s", op.ViewUpdate, interp.Operands.NamesJoined())
		}
		var resTypes string
		if op.ViewUpdate == "View" {
			resTypes = interp.Results.ReadOnlyTypes()
			interp.ReadOnly = "ReadOnly"
			if interp.Results.Any() {
				interp.ResList = interp.Results.ReadOnlyNamedArgs("out")
				if interp.ResList != "" {
					interp.ResList = " (" + interp.ResList + ")"
				}
				interp.ResNames = interp.Results.ArgNames("out")
			}
		} else {
			resTypes = interp.Results.Types()
			if interp.Results.Any() {
				interp.ResList = interp.Results.NamedArgs("out")
				if interp.ResList != "" {
					interp.ResList = " (" + interp.ResList + ")"
				}
				interp.ResNames = interp.Results.ArgNames("out")
			}
		}
		// guard the result types with () if they're plural
		if strings.Contains(resTypes, ",") {
			resTypes = "(" + resTypes + ")"
		}
		if resTypes != "" {
			resTypes = " " + resTypes
		}
		interp.ResTypes = resTypes
		interp.OpName = dt.Name + interp.OpFuncSuffix
		interp.LiteralFuncType = fmt.Sprintf("func(%s)%s", interp.Operands.ReadOnlyTypes(), interp.ResTypes)

		defaults := op.TypedDefaultOps[dt.Name]
		if interp.Operands.Any() {
			interp.TypeList = interp.Operands.ReadOnlyTypes()
			interp.ArgList = interp.Operands.ReadOnlyNamedArgs("in")
			interp.ArgNames = interp.Operands.ArgNames("in")
			interp.ArgComma = ", "
		}

		if len(defaults) > 0 {
			interp.Defaults = make([]prefixedName, len(defaults))
			for i, name := range defaults {
				if strings.HasPrefix(name, dt.Name) {
					interp.Defaults[i] = prefixedName{Prefix: dt.Name, Name: name[len(dt.Name):]}
				} else {
					interp.Defaults[i] = prefixedName{Name: name}
				}
			}
		}

		forTemplate.Ops[i] = interp
	}
	err = tmpl.Execute(file, forTemplate)
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

func writeOpData(opDatas []opData, tmpl *template.Template) error {
	fileName := "optype_gen.go"
	tempName := fileName + ".tmp"
	file, err := os.Create(tempName)
	if err != nil {
		return fmt.Errorf("creating temp file '%s': %v\n", tempName, err)
	}
	defer file.Close()
	forTemplate := struct {
		First      opData
		Additional []opData
	}{First: opDatas[0], Additional: opDatas[1:]}
	err = tmpl.Execute(file, forTemplate)
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

var nonAlpha = regexp.MustCompile("[^[:alnum:]]")

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func usage(format string, args ...interface{}) {
	flag.Usage()
	fatal(format, args...)
}

// Gen creates the op function declarations, and corresponding methods
// to make them satisfy OpFunction, for a given type and series of ops.
// We could also parse the ops from optypes.go, but this seems like a lot
// of work.
func main() {
	TypeNames := flag.String("types", "Container:uint16:int,Bitmap:uint64:uint64", "types, specified as comma-separated typename:bitIn:bitOut, such as Container:uint16:,Bitmap:uint64:uint64")
	typeSource := flag.String("typelist", "optypes_list.txt", "file containing a list of optypes (specify \"\" to ignore)")
	dataTemplateFile := flag.String("datatemplate", "data.tmpl", "per-datatype code template")
	opsTemplateFile := flag.String("opstemplate", "ops.tmpl", "op-generic code template")
	flag.Parse()

	if dataTemplateFile == nil || opsTemplateFile == nil || *dataTemplateFile == "" || *opsTemplateFile == "" {
		usage("data and ops template have to be specified")
	}
	dataTemplate, err := template.ParseFiles(*dataTemplateFile)
	if err != nil {
		fatal("parsing data template: %v", err)
	}
	opsTemplate, err := template.ParseFiles(*opsTemplateFile)
	if err != nil {
		fatal("parsing ops template: %v", err)
	}
	types := strings.Split(*TypeNames, ",")
	if len(types) < 1 {
		usage("types must be specified.\n")
	}
	dataTypes, err := extractDataTypes(types)
	if err != nil {
		fatal("type specifications invalid: %v\n", err)
		os.Exit(1)
	}
	var opsGiven []string
	if *typeSource != "" {
		bits, err := ioutil.ReadFile(*typeSource)
		if err != nil {
			fatal("reading ops list: %v\n", err)
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
		fatal("usage: gen [-types=<typespecs>] [-typelist=filename] [<ops>]\n(typelist or ops list should be non-empty)\n")
	}
	opDatas, err := extractOpNames(opsGiven)
	if err != nil {
		fatal("parsing op Names: %v\n", err)
	}
	for _, data := range dataTypes {
		err = writeData(data, opDatas, dataTemplate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "writing ops for '%s': %v\n", data.Name, err)
		}
	}
	err = writeOpData(opDatas, opsTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing generic ops: %v\n", err)
	}
}
