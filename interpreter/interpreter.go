package interpreter

import "github.com/weiser/lox/expr"

type Interpreter struct{}

func (i *Interpreter) visitLiteralExpr(exp expr.Literal) interface{} {
	return exp.Value
}
