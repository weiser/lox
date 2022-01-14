package main

import (
	"fmt"
	"os"

	"github.com/weiser/lox/mainhelpers"
)

//TODO start page 111, chapter 8
func main() {
	// [0] is the program name
	args := os.Args[1:]
	switch len(args) {
	case 1:
		mainhelpers.RunFile(args[0])
	case 0:
		fmt.Println("Starting mainHelper...")
		// TPDO mainHelper NOT WORK. START ON PG 94
		mainhelpers.RunPrompt()
	default:
		println("Usage: lox [path to script]")
		os.Exit(1)
	}

}
