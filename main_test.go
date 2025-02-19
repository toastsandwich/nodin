package main

import "testing"

func TestGeneratePackage(t *testing.T) {
	test := []string{
		"pkg main",
		"pkg abv",
		"pkg net",
		"pkg a1b2",
	}
	result := []string{
		"package main",
		"package abv",
		"package net",
		"package a1b2",
	}
	c := &Compiler{}
	for i, te := range test {
		c.Lex = NewLexer(te)
		c.Gen = NewGenerator()
	inner:
		for {
			tok := c.Lex.ReadToken()
			if tok.Type == EOF {
				break inner
			}
			if tok.Type != KWRD {
				t.Fatal("tok type not keyword")
			}
			if tok.Value != "pkg" {
				t.Fatal("tok value not pkg")
			}
			pkg := &Package{
				Value: &Identifier{
					Value: c.Lex.ReadToken().Value,
				},
			}
			if pkg.Accept(c.Gen).(string) != result[i] {
				t.Fatalf("result not match got %s expected %s", pkg.Accept(c.Gen).(string), result[i])
			}
		}
	}
}
