package main

import (
	"fmt"
	"strings"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/playground/astPrinter"
	"github.com/weiser/lox/token"
)

func main() {
	expr := &expr.Binary{
		Left: &expr.Unary{
			Operator: token.MakeToken(token.MINUS, "-", nil, 1),
			Right:    &expr.Literal{Value: 123},
		},
		Operator: token.MakeToken(token.STAR, "*", nil, 1),
		Right: &expr.Grouping{
			Expression: &expr.Literal{Value: 45.67},
		},
	}
	astp := &astPrinter.AstPrinter{E: expr, Sb: strings.Builder{}}

	fmt.Println(astp.Print())

}
