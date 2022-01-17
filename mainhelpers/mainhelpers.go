package mainhelpers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/weiser/lox/interpreter"
	"github.com/weiser/lox/parser"
	"github.com/weiser/lox/scanner"
	"github.com/weiser/lox/token"
)

var hadError bool
var interpret *interpreter.Interpreter

func RunFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("couldn't read file: ", err)
		os.Exit(1)
	}

	Run(string(data))
	if hadError {
		os.Exit(1)
	}
}

func Run(data string) {
	scanner := scanner.MakeScanner(data)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	interpret = &interpreter.Interpreter{}

	if err != nil {

		//return  todo this is commented out b/c the "nil" error isn't nil
	}

	interpret.Interpret(stmts)

}

func RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "q" || line == "" {
			fmt.Println("Exiting. goodbye")
			break
		}
		Run(line)
		hadError = false
		fmt.Print("> ")
	}
}

func ReportError(line int, message string) {
	fmt.Println("Error! '", message, "'\n  line: ", line)
	hadError = true
}

func ReportErrorParser(tok token.Token, err string) {
	if tok.TokenType == token.EOF {
		ReportError(tok.Line, "at end "+err)
	} else {
		ReportError(tok.Line, "at '"+tok.Lexeme+"' "+err)
	}
}
