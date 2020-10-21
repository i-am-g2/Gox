package gox

// Scanner c
type Scanner struct {
	source               string
	start, current, line int
	tokens               []Token
	keywords             map[string]string
}

// NewScanner aa
func NewScanner(src string) *Scanner {
	scanner := Scanner{}
	scanner.source = src
	scanner.start = 0
	scanner.current = 0
	scanner.line = 1

	if scanner.keywords == nil {
		scanner.keywords = make(map[string]string)
	}

	scanner.keywords["and"] = "AND"
	scanner.keywords["class"] = "CLASS"
	scanner.keywords["else"] = "ELSE"
	scanner.keywords["false"] = "FALSE"
	scanner.keywords["for"] = "FOR"
	scanner.keywords["fun"] = "FUN"
	scanner.keywords["if"] = "IF"
	scanner.keywords["nil"] = "NIL"
	scanner.keywords["or"] = "OR"
	scanner.keywords["print"] = "PRINT"
	scanner.keywords["return"] = "RETURN"
	scanner.keywords["super"] = "SUPER"
	scanner.keywords["this"] = "THIS"
	scanner.keywords["true"] = "TRUE"
	scanner.keywords["var"] = "VAR"
	scanner.keywords["while"] = "WHILE"

	return &scanner
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// ScanTokens ss
func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{"EOF", "", s.line, nil})

	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addNilToken("LEFT_PAREN")
	case ')':
		s.addNilToken("RIGHT_PAREN")
	case '{':
		s.addNilToken("LEFT_BRACE")
	case '}':
		s.addNilToken("RIGHT_BRACE")
	case ',':
		s.addNilToken("COMMA")
	case '.':
		s.addNilToken("DOT")
	case '-':
		s.addNilToken("MINUS")
	case '+':
		s.addNilToken("PLUS")
	case ';':
		s.addNilToken("SEMICOLON")
	case '*':
		s.addNilToken("STAR")
	case '!':
		if s.match('=') {
			s.addNilToken("BANG_EQUAL")
		} else {
			s.addNilToken("BANG")
		}
	case '=':
		if s.match('=') {
			s.addNilToken("EQUAL_EQUAL")
		} else {
			s.addNilToken("EQUAL")
		}
	case '<':
		if s.match('=') {
			s.addNilToken("LESS_EQUAL")
		} else {
			s.addNilToken("LESS")
		}
	case '>':
		if s.match('=') {
			s.addNilToken("GREATER_EQUAL")
		} else {
			s.addNilToken("GREATER")
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		}
	case ' ', '\r', '\t':
		break
	case '\n':
		s.line++
	case '"':
		s.scanString()
	default:
		if s.isDigit(c) {
			s.scanNum()
		} else if s.isAlpha(c) {
			s.scanIdent()
		} else {
			error(s.line, "Unexpected character")
		}
	}
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) addNilToken(tokenType string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType, text, s.line, nil})
}
func (s *Scanner) addToken(tokenType string, value string) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, Token{tokenType, text, s.line, value})

}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) scanString() {

	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		error(s.line, "Unterminated String")
		return
	}
	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addToken("STRING", value)
}

func (s *Scanner) isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) isAlphaNum(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) scanNum() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	num := s.source[s.start:s.current]

	s.addToken("INT", num)

}

func (s *Scanner) scanIdent() {
	for s.isAlphaNum(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType, exists := s.keywords[text]
	if !exists {
		tokenType = "IDENTIFIER"
	}
	s.addNilToken(tokenType)
}
