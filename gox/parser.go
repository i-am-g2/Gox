package gox

// Parser class
type Parser struct {
	tokens  []Token
	current int
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparision()
	for p.match("BANG_EQUAL", "EQUAL_EQUAL") {
		operator := p.previous()
		right := p.comparision()
		expr = exprBinary{expr, operator, right}
	}
	return expr
}

func (p *Parser) comparision() Expr {
	expr := p.term()
	for p.match("GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL") {
		operator := p.previous()
		right := p.term()
		expr = exprBinary{expr, operator, right}
	}
	return expr

}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match("MINUS", "PLUS") {
		operator := p.previous()
		right := p.factor()
		expr = exprBinary{expr, operator, right}
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match("SLASH", "STAR") {
		operator := p.previous()
		right := p.unary()
		expr = exprBinary{expr, operator, right}
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match("BANG", "MINUS") {
		operator := p.previous()
		right := p.unary()
		return exprUnary{operator, right}
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match("FALSE") {
		return exprLiteral{false}
	}
	if p.match("TRUE") {
		return exprLiteral{true}
	}
	if p.match("NIL") {
		return exprLiteral{nil}
	}
	if p.match("INT", "STRING") {
		return exprLiteral{p.previous().object}
	}

	if p.match("LEFT_PAREN") {
		expr := p.expression()
		p.consume("RIGHT_PAREN", "Expect ')' after expression.")
		return exprGrouping{expr}
	}

}

func (p *Parser) consume(tokenType, message string) {
	// TODO : Correct this:

}

// NewParser Constructor
func NewParser(tokens []Token) *Parser {
	temp := Parser{tokens, 0}
	return &temp
}

func (p *Parser) isAtEnd() bool {
	return p.peek().tokenType == "EOF"
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) check(tokenType string) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().tokenType == tokenType
}

func (p *Parser) match(tokenType ...string) bool {
	for i := range tokenType {
		if p.check(tokenType[i]) {
			p.advance()
			return true
		}
	}
	return false
}
