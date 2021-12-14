package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		runFile(os.Args[0])
	case 0:
		runPrompt()
	default:
		println("Usage: lox [path to script]")
		panic(1)
	}

}

func runFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("couldn't read file: %v", err)
		panic(1)
	}

	run(string(data))
}

func run(data string) {
	// TOTO: start on page 41, implement run fxn.

}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "q" || line == "" {
			fmt.Println("Exiting. goodbye")
			break
		}
		run(line)
		fmt.Print("> ")
	}
}
