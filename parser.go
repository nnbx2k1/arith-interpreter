package main
import (
	"fmt"
	"strconv"
)

type Expr interface {
	String() string
}

type Number struct {
	Value int
}

type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

func (b BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Op, b.Right.String())
}

type Parser struct {
	tokens []Token
	pos    int
}

func (p *Parser) current() Token {
	return p.tokens[p.pos]
}

func (p *Parser) parseE() Expr {
	left := p.parseT()
	for p.current().Type == PLUS || p.current().Type == MINUS {
		op := "+"
		if p.current().Type == MINUS {
			op = "-"
		}
		p.pos++
		right := p.parseT()
		left = BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (p *Parser) parseT() Expr {
	left := p.parseF()
	for p.current().Type == MUL || p.current().Type == DIV {
		op := "*"
		if p.current().Type == DIV {
			op = "/"
		}
		p.pos++
		right := p.parseF()
		left = BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (p *Parser) parseF() Expr {
	token := p.current()
	if token.Type == INT {
		p.pos++
		val, _ := strconv.Atoi(token.Value)
		return Number{Value: val}
	} else if token.Type == LPAREN {
		p.pos++
		node := p.parseE()
		if p.current().Type != RPAREN {
			panic("Expected ')'")
		}
		p.pos++
		return node
	} else {
		panic("Unexpected token")
	}
}

func Eval(expr Expr) int {
	switch node := expr.(type) {
	case Number:
		return node.Value
	case BinaryExpr:
		l := Eval(node.Left)
		r := Eval(node.Right)
		switch node.Op {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			return l / r
		}
	}
	panic("Unknown expression")
}
