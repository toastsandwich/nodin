package main

import "fmt"

type Compiler struct {
	Lex *Lexer
	Gen *Generator
}

func (c *Compiler) Compile(file string) {
	c.Lex = NewLexer(file)
	c.Gen = NewGenerator()
loop:
	for {
		t := c.Lex.ReadToken()
		switch t.Type {
		case EOF:
			break loop
		case KWRD:
			c.CKeyword(t)
		}
	}
}

func (c *Compiler) CKeyword(t *Token) {
	switch t.Value {
	case "pkg":
		n := c.Lex.ReadToken()
		fmt.Println(n.String())
		pkg := &Package{Value: &Identifier{
			Value: n.Value,
		}}
		fmt.Println(pkg.Accept(c.Gen))
	}
}
