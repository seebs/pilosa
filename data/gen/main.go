package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type dataType struct {
	name      string
	nameLower string
	bitsIn    string
	bitsOut   string
}

type opData struct {
	original   string
	viewUpdate string
	operand    string
	plural     bool
}

var opTypeParser = regexp.MustCompile("^(View|Update)(.*)$")

func extractOpNames(ops []string) ([]opData, error) {
	opDatas := make([]opData, len(ops))
	for i, op := range ops {
		details := opTypeParser.FindStringSubmatch(op)
		if details == nil || details[0] == "" {
			return nil, fmt.Errorf("name '%s' does not match expected pattern", op)
		}
		plural := false
		if strings.HasSuffix(details[2], "s") {
			plural = true
			details[2] = details[2][:len(details[2])-1]
		}
		opDatas[i] = opData{original: "OpType" + op, viewUpdate: details[1], operand: details[2], plural: plural}

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
		bits, err := strconv.Atoi(out[1])
		if err != nil || (bits != 16 && bits != 64) {
			return nil, fmt.Errorf("type spec '%s' does not have valid bit width ('%s'), should be 16/64",
				spec, out[1])
		}
		if out[2] != "" {
			bits, err = strconv.Atoi(out[2])
			if err != nil || (bits != 16 && bits != 64) {
				return nil, fmt.Errorf("type spec '%s' does not have valid bit width ('%s'), should be 16/64",
					spec, out[2])
			}
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
	err = writeOps(dt, opDatas, file)
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

func writeOps(dt dataType, opDatas []opData, w io.Writer) error {
	var Printf = func(format string, args ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, args...)
	}
	Printf(`package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for %s

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunction%s interface {
	Type(%s) OpType
}

`, dt.name, dt.name, dt.name)
	newline := false
	for _, op := range opDatas {
		readOnly := ""
		aAn := "an"
		if op.viewUpdate == "View" {
			readOnly = "ReadOnly"
			aAn = "a"
		}
		dotDotDot := ""
		operand := op.operand
		other := ""
		if operand == "Other" {
			operand = dt.name
			other = " other"
		}
		pluralOperand := operand
		if op.plural {
			pluralOperand += "s"
			dotDotDot = "..."
		}
		if newline {
			Printf("\n")
		}
		opName := fmt.Sprintf("Op%s%s%s", dt.name, op.viewUpdate, pluralOperand)
		Printf("// %s is %s %s operation on a %s%s ", opName, aAn, op.viewUpdate, readOnly, dt.name)
		switch {
		case op.operand == "":
			Printf("with no other parameters.\n")
		case op.plural:
			Printf("and one or more%s %ss.\n", other, operand)
		default:
			Printf("and one%s %s.\n", other, operand)
		}

		Printf("type %s func(%s%s", opName, readOnly, dt.name)
		switch op.operand {
		case "": // no operands
			Printf(")")
		case "Bit":
			Printf(", %suint%s)", dotDotDot, dt.bitsIn)
		case "Other":
			Printf(", %sReadOnly%s)", dotDotDot, dt.name)
		}
		// return value
		Printf(" (bool, int%s, %s%s)\n",
			dt.bitsOut, readOnly, dt.name)
		Printf("func (%s) Type(%s) OpType { return %s }\n", opName, dt.name, op.original)
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
	typeNames := flag.String("types", "", "types, specified as comma-separated typename:bitsIn:bitsOut, such as Container:16:,Bitmap:64:64")
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
	opsGiven := flag.Args()
	if len(opsGiven) < 1 {
		fmt.Fprintf(os.Stderr, "usage: gen -types <typespecs> <ops>\n")
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
}
