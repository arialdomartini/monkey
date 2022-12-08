package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"

	IDENT     = "IDENT"
	INT       = "INT"

	ASSIGN    = "="
	PLUS      = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN    = "("
	RPAREN    = ")"

	LBRACE    = "{"
	RBRACE    = "}"

	FUNCTION  = "FUNCTION"
	LET       = "LET"
)


func create(tokenType TokenType, character byte) Token {
	return Token{Type: tokenType, Literal: string(character)}
}

func eof() Token {
	return Token{Type: EOF, Literal: ""}
}

func Parse(character byte) Token {
	var tok Token

	switch character {
	case '=':
		tok = create(ASSIGN, character)
	case ';':
		tok = create(SEMICOLON, character)
	case '(':
		tok = create(LPAREN, character)
	case ')':
		tok = create(RPAREN, character)
	case ',':
		tok = create(COMMA, character)
	case '+':
		tok = create(PLUS, character)
	case '{':
		tok = create(LBRACE, character)
	case '}':
		tok = create(RBRACE, character)
	case 0:
		tok = eof()
	}
	return tok
}
