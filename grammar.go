package main

import (
	"fmt"
	"grammar/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	exp := "a=2;b=3; (1 + 2) * ( 3 + 4 ) * a * b"
	showTokens(exp)

	result := DoCalculation(exp)

	fmt.Println()
	fmt.Println("expression: "+exp+" = ", result)
}

func showTokens(exp string) {
	is := antlr.NewInputStream(exp)

	lexer := parser.NewCalcLexer(is)

	// print all tokens
	fmt.Println("expression: " + exp)
	fmt.Println("tokens:")
	for {
		t := lexer.NextToken()

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("\t%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func DoCalculation(exp string) int {
	is := antlr.NewInputStream(exp)

	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create parser
	p := parser.NewCalcParser(stream)

	// parse the expression
	listener := NewCalculationListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

	return listener.pop()
}
