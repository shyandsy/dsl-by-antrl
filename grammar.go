package main

import (
	"fmt"
	"grammar/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	showTokens("1 + 2 * 3")

	doCalculation("1 + 2 * 3")
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

func doCalculation(exp string) {
	is := antlr.NewInputStream(exp)

	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create parser
	p := parser.NewCalcParser(stream)

	// parse the expression
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	fmt.Println()
	fmt.Println("expression: "+exp+" = ", listener.pop())

}
