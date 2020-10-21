package main

// Object ad
type Object interface{}

// Token determines the type o
type Token struct {
	tokenType string
	lexeme    string
	line      int
	object    Object
}

// New - Creates new
func (t *Token) New(tokenType, lexeme string, line int, obj Object) {
	t.tokenType = tokenType
	t.lexeme = lexeme
	t.line = line
	t.object = obj
}

// ToString conv
func (t *Token) ToString() string {
	return t.tokenType + " " + t.lexeme + " " + string(t.line)
}
