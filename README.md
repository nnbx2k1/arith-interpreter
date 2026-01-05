# Arith Interpreter

During my computer science studies at ESI Algiers in the Theory of Languages (THP) course, we learned about automata, grammars, and formal languages. To put that theory into practice I implemented a small arithmetic expression interpreter that demonstrates parsing and evaluation of expressions.
A demonstration PDF illustrating the parser and examples is provided in the repository; see `thp-demonstration.pdf`

A tiny arithmetic expression interpreter in Go. It tokenizes, parses, and evaluates simple integer arithmetic with +, -, *, / and parentheses. This repository is a small learning/demo project.

Features
- Integer literals
- Binary operations: +, -, *, /
- Parentheses for grouping
- Simple AST printing via `Expr.String()`

Prerequisites
- Go 1.18+ installed (module-aware workflow).

Files
- `main.go` - example runner that lexes, parses and evaluates a sample expression.
- `lexer.go` - tokenizer that converts an input string into tokens.
- `parser.go` - recursive-descent parser that builds an AST and provides `Eval`.
- `evaluator_test.go` - table-driven unit tests for expression evaluation.

Quick start

Run the example program:

```bash
cd /path/to/arith-interpreter
go run .
```

Or run the specific files directly:

```bash
go run main.go lexer.go parser.go
```

Run tests

```bash
go test
```

Example

Input: `3 + 4 * (2 - 1)`

Output (example program):

```
AST: (3 + (4 * (2 - 1)))
Result: 7
```

Notes and next steps
- The interpreter currently panics on invalid input; adding proper error types would improve robustness.
- Only integer arithmetic is supported; consider adding floats, unary operators, and operator precedence tests.

License

This project is provided as is for learning and demonstration.
