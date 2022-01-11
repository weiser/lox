package interpreter

import "github.com/weiser/lox/expr"

type Interpreter struct {
}

func (i *Interpreter) VisitLiteral(exp *expr.Literal) interface{} {
	return exp.Value
}

func (i *Interpreter) VisitGrouping(exp *expr.Grouping) interface{} {
	return i.evaluate(exp.Expression)
}

func (i *Interpreter) VisitBinary(exp *expr.Binary) interface{} {
	// TODO
	return nil
}

func (i *Interpreter) VisitUnary(exp *expr.Unary) interface{} {
	// TODO
	return nil
}

func (i *Interpreter) VisitExpr(exp *expr.Expr) interface{} {
	// TODO
	return nil
}

func (i *Interpreter) evaluate(exp expr.ExprInterface) interface{} {
	return exp.Accept(i)
}

// TODO: start at pg 100, 7.2.2
