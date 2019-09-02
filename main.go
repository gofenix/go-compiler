package main

import (
	"compilers/simple_lexer"
	"fmt"
)

func main() {
	lexer := simple_lexer.NewSimpleLexer()
	script := "int age = 61;"
	fmt.Println("parse: ", script)
	tokenReader := lexer.Tokenize(script)
	simple_lexer.Dump(tokenReader)
}
