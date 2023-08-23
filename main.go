package main

import (
	"log"
	"n0s4/goporth/interpreter"
	"n0s4/goporth/lexer"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Need file to interpret")
	}

	programPath := args[1]
	file, err := os.ReadFile(programPath)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	sourceCode := string(file)

	tokens, err := lexer.Lex(sourceCode)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("tokens:\n%v\n\n", tokens)

	interpreter.Interpret(tokens)
}
