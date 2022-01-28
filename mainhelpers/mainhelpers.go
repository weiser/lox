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
	toks := scanner.ScanTokens()
	p := parser.Parser{Tokens: toks}
	stmts, _ := p.Parse()
	i := interpreter.MakeInterpreter()
	interpret = &i

	if p.ParsingErr != nil {
		// try to parse as an expression
		p1 := parser.Parser{Tokens: toks}
		fmt.Println(interpret.Evaluate(p1.Expression()))
	} else {
		interpret.Interpret(stmts)
	}

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
