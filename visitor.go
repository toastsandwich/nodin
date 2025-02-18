package main

import (
	"fmt"
	"strings"
)

type StatementVisitor interface {
	VisitBlock(*Block) any
	VisitDeclaration(*Declaration) any
	VisitIf(*If) any
	VisitFor(*For) any
	VisitKeyword(*Keyword) any
	WithIndent(any) any
}

type ExpressionVisitor interface {
	VisitUnaryExpression(*UnaryExpression) any
	VisitBinaryExpression(*BinaryExpression) any
	VisitLiteralExpression(*LiteralExpression) any
	WithIndent(any) any
}

type Generator struct {
	Indent int
}

func NewGenerator() *Generator {
	return &Generator{
		Indent: 0,
	}
}

func (g *Generator) indent() string {
	return strings.Repeat("\t", g.Indent)
}

func (g *Generator) WithIndent(comp any) any {
	g.Indent++
	str := g.indent() + g.Visit(comp).(string)
	g.Indent--
	return str
}

func (g *Generator) Visit(comp any) any {
	switch c := comp.(type) {
	case *Block:
		return g.VisitBlock(c)
	case *Declaration:
		return g.VisitDeclaration(c)
	case *If:
		return g.VisitIf(c)
	case *For:
		return g.VisitFor(c)
	case *Keyword:
		return g.VisitKeyword(c)
	case *UnaryExpression:
		return g.VisitUnaryExpression(c)
	case *BinaryExpression:
		return g.VisitBinaryExpression(c)
	case *LiteralExpression:
		return g.VisitLiteralExpression(c)
	default:
		return "not found"
	}
}

func (g *Generator) VisitBlock(b *Block) any {
	builder := strings.Builder{}
	builder.WriteString("{\n")
	g.Indent++
	for _, s := range b.Statements {
		builder.WriteString(g.indent() + s.Accept(g).(string) + "\n")
	}
	g.Indent--
	builder.WriteString(g.indent() + "}")
	return builder.String()
}

func (g *Generator) VisitDeclaration(d *Declaration) any {
	return fmt.Sprintf("%s := %s", d.Identifier.Value, d.Value)
}

func (g *Generator) VisitIf(i *If) any {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("if %s ", i.Condition.Accept(g)))
	builder.WriteString(i.Block.Accept(g).(string))
	return builder.String()
}

func (g *Generator) VisitFor(f *For) any {
	builder := strings.Builder{}
	builder.WriteString("for ")
	builder.WriteString(f.Init.Accept(g).(string) + "; ")
	builder.WriteString(f.Condition.Accept(g).(string) + "; ")
	builder.WriteString(f.Update.Accept(g).(string) + " ")
	builder.WriteString(f.Block.Accept(g).(string))
	return builder.String()
}

func (g *Generator) VisitKeyword(k *Keyword) any {
	return k.Value
}

func (g *Generator) VisitUnaryExpression(u *UnaryExpression) any {
	if u.Operator == "!" || u.Operator == "-" {
		return fmt.Sprintf("%s%s",
			u.Operator,
			u.Expression.Accept(g).(string),
		)
	}
	return fmt.Sprintf("%s%s", u.Expression.Accept(g).(string), u.Operator)
}

func (g *Generator) VisitBinaryExpression(b *BinaryExpression) any {
	return fmt.Sprintf("%s %s %s",
		b.Left.Accept(g).(string),
		b.Operator,
		b.Right.Accept(g).(string),
	)
}

func (g *Generator) VisitLiteralExpression(l *LiteralExpression) any {
	return l.Value
}
