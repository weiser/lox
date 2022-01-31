package expr

// DO NOT MODIFY. GENERATED VIA `go run cmd/tool/generateAst.go expr`
// TODO:  MAKE `cmd/tool/generateAst.go` format code
import . "github.com/weiser/lox/token"

type Stmt struct {
}
type StmtInterface interface {
	Accept(evi StmtVisitorInterface) interface{}
}
type StmtVisitorInterface interface {
	VisitStmt(e *Stmt) interface{}
	VisitBlock(e *Block) interface{}
	VisitExpression(e *Expression) interface{}
	VisitIf(e *If) interface{}
	VisitPrint(e *Print) interface{}
	VisitVar(e *Var) interface{}
}

func (o *Stmt) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitStmt(o)
}

type Block struct {
	*Stmt
	Statements []StmtInterface
}

func (o *Block) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitBlock(o)
}

type Expression struct {
	*Stmt
	Expression ExprInterface
}

func (o *Expression) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitExpression(o)
}

type If struct {
	*Stmt
	Condition  ExprInterface
	ThenBranch StmtInterface
	ElseBranch StmtInterface
}

func (o *If) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitIf(o)
}

type Print struct {
	*Stmt
	Expression ExprInterface
}

func (o *Print) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitPrint(o)
}

type Var struct {
	*Stmt
	Name        Token
	Initializer ExprInterface
}

func (o *Var) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitVar(o)
}
