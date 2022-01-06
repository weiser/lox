package interpreter

import "github.com/weiser/lox/expr"

type Interpreter struct{}

func (i *Interpreter) visitLiteralExpr(exp expr.Literal) interface{} {
	return exp.Value
}

// TODO: start at pg 100, 7.2.2
