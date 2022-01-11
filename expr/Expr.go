package expr

// DO NOT MODIFY. GENERATED VIA `go run cmd/tool/generateAst.go expr`
import . "github.com/weiser/lox/token"

type Expr struct {
}
type ExprInterface interface {
	Accept(evi ExprVisitorInterface) interface{}
}
type ExprVisitorInterface interface {
	VisitExpr(e *Expr) interface{}
	VisitBinary(e *Binary) interface{}
	VisitGrouping(e *Grouping) interface{}
	VisitLiteral(e *Literal) interface{}
	VisitUnary(e *Unary) interface{}
}

func (o *Expr) Accept(evi ExprVisitorInterface) interface{} {
	return evi.VisitExpr(o)
}

type Binary struct {
	*Expr
	Left     ExprInterface
	Operator Token
	Right    ExprInterface
}

func (o *Binary) Accept(evi ExprVisitorInterface) interface{} {
	return evi.VisitBinary(o)
}

type Grouping struct {
	*Expr
	Expression ExprInterface
}

func (o *Grouping) Accept(evi ExprVisitorInterface) interface{} {
	return evi.VisitGrouping(o)
}

type Literal struct {
	*Expr
	Value interface{}
}

func (o *Literal) Accept(evi ExprVisitorInterface) interface{} {
	return evi.VisitLiteral(o)
}

type Unary struct {
	*Expr
	Operator Token
	Right    ExprInterface
}

func (o *Unary) Accept(evi ExprVisitorInterface) interface{} {
	return evi.VisitUnary(o)
}
