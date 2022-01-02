package expr

// DO NOT MODIFY. GENERATED VIA `go run cmd/tool/generateAst.go expr`
import . "github.com/weiser/lox/token"

type Expr struct {
}
type ExprInterface interface {
	Accept(evi ExprVisitorInterface)
}
type ExprVisitorInterface interface {
	VisitExpr(e *Expr)
	VisitBinary(e *Binary)
	VisitGrouping(e *Grouping)
	VisitLiteral(e *Literal)
	VisitUnary(e *Unary)
}

func (o *Expr) Accept(evi ExprVisitorInterface) {
	evi.VisitExpr(o)
}

type Binary struct {
	*Expr
	Left     ExprInterface
	Operator Token
	Right    ExprInterface
}

func (o *Binary) Accept(evi ExprVisitorInterface) {
	evi.VisitBinary(o)
}

type Grouping struct {
	*Expr
	Expression ExprInterface
}

func (o *Grouping) Accept(evi ExprVisitorInterface) {
	evi.VisitGrouping(o)
}

type Literal struct {
	*Expr
	Value interface{}
}

func (o *Literal) Accept(evi ExprVisitorInterface) {
	evi.VisitLiteral(o)
}

type Unary struct {
	*Expr
	Operator Token
	Right    ExprInterface
}

func (o *Unary) Accept(evi ExprVisitorInterface) {
	evi.VisitUnary(o)
}
