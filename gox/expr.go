package gox

// Expr type
type Expr interface {
	dummyMethod() Expr
}

type exprBinary struct {
	left     Expr
	operator Token
	right    Expr
}

type exprUnary struct {
	operator Token
	right    Expr
}

type exprLiteral struct {
	value interface{}
}

type exprGrouping struct {
	grpingExpression Expr
}

func (p exprBinary) dummyMethod() Expr {
	return p
}

func (p exprUnary) dummyMethod() Expr {
	return p
}

func (p exprLiteral) dummyMethod() Expr {
	return p
}

func (p exprGrouping) dummyMethod() Expr {
	return p
}
