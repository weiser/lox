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
	VisitClass(e *Class) interface{}
	VisitExpression(e *Expression) interface{}
	VisitFunction(e *Function) interface{}
	VisitIf(e *If) interface{}
	VisitPrint(e *Print) interface{}
	VisitWhile(e *While) interface{}
	VisitVar(e *Var) interface{}
	VisitReturn(e *Return) interface{}
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

type Class struct {
	*Stmt
	Name    Token
	Methods []StmtInterface
}

func (o *Class) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitClass(o)
}

type Expression struct {
	*Stmt
	Expression ExprInterface
}

func (o *Expression) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitExpression(o)
}

type Function struct {
	*Stmt
	Name   Token
	Params []Token
	Body   []StmtInterface
}

func (o *Function) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitFunction(o)
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

type While struct {
	*Stmt
	Condition ExprInterface
	Body      StmtInterface
}

func (o *While) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitWhile(o)
}

type Var struct {
	*Stmt
	Name        Token
	Initializer ExprInterface
}

func (o *Var) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitVar(o)
}

type Return struct {
	*Stmt
	Keyword Token
	Value   ExprInterface
}

func (o *Return) Accept(evi StmtVisitorInterface) interface{} {
	return evi.VisitReturn(o)
}
