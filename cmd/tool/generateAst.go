package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage generateAst <output_directory>")
		os.Exit(1)
	}

	outputDir := args[0]

	defineAst(outputDir, "Expr", []string{
		"Binary : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal : Object value",
		"Unary : Token operator, Expr right",
	})
}

// TODO: consider making a type that embeds *os.File that implements this, but also lets us format the golang code before we write it out
func writeWithNewline(f *os.File, s string) {
	f.WriteString(fmt.Sprintln(s))
}

func defineAst(outputDir string, baseName string, types []string) {
	path := fmt.Sprintf("%v/%v.go", outputDir, baseName)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	writeWithNewline(f, "package expr")
	writeWithNewline(f, fmt.Sprintf("type %v struct {", baseName))
	writeWithNewline(f, "}")
	for _, typ := range types {
		splits := strings.Split(typ, ":")

		exprType := strings.TrimSpace(splits[0])
		fields := strings.TrimSpace(splits[1])

		defineType(f, baseName, exprType, fields)
	}

}

func defineType(f *os.File, baseName string, exprType string, fields string) {
	writeWithNewline(f, fmt.Sprintf("type %v struct {", exprType))
	writeWithNewline(f, fmt.Sprintf("*%v", baseName))
	for _, field := range strings.Split(strings.TrimSpace(fields), ",") {
		// field is 'Token operator'. needs to be "operator Token" in struct
		fmt.Println("field is: '", field, "'")
		fs := strings.Split(strings.TrimSpace(field), " ")
		writeWithNewline(f, fmt.Sprintf("%v %v", fs[1], fs[0]))
	}
	writeWithNewline(f, "}")

}
