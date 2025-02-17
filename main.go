package main

import "fmt"

const code = "./sample.nn"

func main() {
	// codeFile, err := os.ReadFile(code)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// l := NewLexer(string(codeFile))
	// for {
	// 	tok := l.ReadToken()
	// 	if tok.Type == "WHSPC" {
	// 		continue
	// 	}
	// 	fmt.Println(tok.String())
	// 	if tok.Type == EOF {
	// 		break
	// 	}
	// }

	gen := &Generator{}

	i := &Identifier{
		Value: "i",
	}

	dec := &Declaration{
		Identifier: i,
		Value:      "100",
	}
	s := dec.Accept(gen)
	fmt.Println(s.(string))

	f := For{
		Init: dec,
		Condition: &BinaryExpression{
			Left:     &LiteralExpression{Value: i.Value},
			Operator: "<",
			Right:    &LiteralExpression{Value: "200"},
		},
		Update: &UnaryExpression{
			Expression: &LiteralExpression{
				Value: i.Value,
			},
			Operator: "++",
		},
		Block: Block{
			Statements: []Statement{
				&If{
					Condition: &LiteralExpression{
						Value: "true",
					},
					Block: Block{
						Statements: []Statement{
							&Keyword{
								Value: "break",
							},
						},
					},
				},
			},
		},
	}
	fcode := f.Accept(gen)
	fmt.Println(fcode)
}
