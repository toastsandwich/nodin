package main

type Node interface {
	GenerateGo() string
}

type Statement interface {
	Node
	statement()
}

func (i *IfStatement) statement() {

}

type IfStatement struct {
	Condition  Expression
	Statements []Statement
}

func (i *IfStatement) GenerateGo() string {
	return ""
}

type ForStatement struct {
	Initialization Statement
	Assignment     Statement
	Condition      Expression
}

type Expression interface {
	Node
	expression()
}
