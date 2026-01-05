package main

import "testing"

func TestEvalExpressions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple precedence", "3 + 4 * (2 - 1)", 7},
		{"no parens", "2+3*4", 14},
		{"with parens", "(1+2)*3", 9},
		{"division and subtract", "10/2-3", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := Lex(tt.input)
			parser := Parser{tokens: tokens}
			ast := parser.parseE()
			got := Eval(ast)
			if got != tt.want {
				t.Fatalf("Eval(%q) = %d; want %d; AST=%s", tt.input, got, tt.want, ast.String())
			}
		})
	}
}
