package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("give me file")
		os.Exit(1)
	}
	codeFile, err := os.ReadFile(args[1])
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

	// gen := NewGenerator()

}
