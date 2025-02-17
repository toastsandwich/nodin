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
}

type ExpressionVisitor interface {
	VisitUnaryExpression(*UnaryExpression) any
	VisitBinaryExpression(*BinaryExpression) any
	VisitLiteralExpression(*LiteralExpression) any
}

type Generator struct{}

func (g *Generator) VisitBlock(b *Block) any {
	builder := strings.Builder{}
	builder.WriteString("{\n")
	for _, s := range b.Statements {
		builder.WriteString(s.Accept(g).(string))
	}
	builder.WriteString("\n}")
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
