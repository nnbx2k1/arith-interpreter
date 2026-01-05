package main

import "fmt"

func main() {
	input := "3 + 4 / (2 - 1)"
	tokens := Lex(input)
	parser := Parser{tokens: tokens}
	ast := parser.parseE()

	fmt.Println("AST:", ast.String())
	result := Eval(ast)
	fmt.Println("Result:", result)
}
