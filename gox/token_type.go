package gox

// Object interface
type Object interface{}

// Token stores information about token
type Token struct {
	tokenType string
	lexeme    string
	line      int
	object    Object
}

// ToString gives string representation for token
func (t *Token) ToString() string {
	return t.tokenType + " " + t.lexeme + " " + string(t.line)
}
