package main

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	INT TokenType = iota
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
	EOF
)

type Token struct {
	Type  TokenType
	Value string
}

// the preprocessing that converts input string into tokens so we can parse. (hhhhh)
func Lex(input string) []Token {
	tokens := []Token{}
	input = strings.ReplaceAll(input, " ", "")
	i := 0

	for i < len(input) {
		ch := input[i]

		if unicode.IsDigit(rune(ch)) {
			start := i
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				i++
			}
			tokens = append(tokens, Token{INT, input[start:i]})
			continue
		}

		switch ch {
		case '+':
			tokens = append(tokens, Token{PLUS, "+"})
		case '-':
			tokens = append(tokens, Token{MINUS, "-"})
		case '*':
			tokens = append(tokens, Token{MUL, "*"})
		case '/':
			tokens = append(tokens, Token{DIV, "/"})
		case '(':
			tokens = append(tokens, Token{LPAREN, "("})
		case ')':
			tokens = append(tokens, Token{RPAREN, ")"})
		default:
			panic("Unknown character: " + string(ch))
		}
		i++
	}

	tokens = append(tokens, Token{EOF, ""})
	return tokens
}
