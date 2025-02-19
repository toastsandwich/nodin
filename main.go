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
	cmplr := &Compiler{}
	cmplr.Compile(string(codeFile))
}
