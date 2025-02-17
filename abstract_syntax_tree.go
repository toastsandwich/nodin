package main

type Statement interface {
	Accept(StatementVisitor) any
}

type Expression interface {
	Accept(ExpressionVisitor) any
}

type Identifier struct {
	Value string
}

type Declaration struct {
	Identifier *Identifier
	Value      string
}

type Keyword struct {
	Value string
}

func (k *Keyword) Accept(s StatementVisitor) any {
	return s.VisitKeyword(k)
}

type Block struct {
	Statements []Statement
}

func (b *Block) Accept(s StatementVisitor) any {
	return s.VisitBlock(b)
}

func (d *Declaration) Accept(s StatementVisitor) any {
	return s.VisitDeclaration(d)
}

type If struct {
	Condition Expression
	Block     Block
}

func (i *If) Accept(s StatementVisitor) any {
	return s.VisitIf(i)
}

type For struct {
	Init      Statement
	Condition Expression
	Update    Expression
	Block     Block
}

func (f *For) Accept(s StatementVisitor) any {
	return s.VisitFor(f)
}

type LiteralExpression struct {
	Value string
}

func (l *LiteralExpression) Accept(e ExpressionVisitor) any {
	return e.VisitLiteralExpression(l)
}

type UnaryExpression struct {
	Expression Expression
	Operator   string
}

func (u *UnaryExpression) Accept(e ExpressionVisitor) any {
	return e.VisitUnaryExpression(u)
}

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryExpression) Accept(e ExpressionVisitor) any {
	return e.VisitBinaryExpression(b)
}
