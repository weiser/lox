package expr

// DO NOT MODIFY. GENERATED VIA `go run cmd/tool/generateAst.go expr`
// TODO:  MAKE `cmd/tool/generateAst.go` format code
type Stmt struct {
}
type StmtInterface interface {
	Accept(evi StmtVisitorInterface) interface{}
}
type StmtVisitorInterface interface {
	VisitStmt(e *Stmt) interface{}
	VisitExpression(e *Expression) interface{}
	VisitPrint(e *Print) interface{}
}

func (o *Stmt) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitStmt(o)
}

type Expression struct {
	*Stmt
	Expression ExprInterface
}

func (o *Expression) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitExpression(o)
}

type Print struct {
	*Stmt
	Expression ExprInterface
}

func (o *Print) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitPrint(o)
}
