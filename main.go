package main

import (
	"fmt"
	"log"
	"os"
)

const code = "./sample.nn"

func main() {
	codeFile, err := os.ReadFile(code)
	if err != nil {
		log.Fatal(err)
	}
	l := NewLexer(string(codeFile))
	for {
		tok := l.ReadToken()
		if tok.Type == "WHSPC" {
			continue
		}
		fmt.Println(tok.String())
		if tok.Type == EOF {
			break
		}
	}
}
