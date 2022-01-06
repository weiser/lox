package main

import (
	"fmt"
	"os"

	"github.com/weiser/lox/interpreter"
)

func main() {
	// [0] is the program name
	args := os.Args[1:]
	switch len(args) {
	case 1:
		interpreter.RunFile(args[0])
	case 0:
		fmt.Println("Starting interpreter...")
		// TPDO INTERPRETER NOT WORK. START ON PG 94
		interpreter.RunPrompt()
	default:
		println("Usage: lox [path to script]")
		os.Exit(1)
	}

}
