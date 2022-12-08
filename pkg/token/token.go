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


func New(tokenType TokenType, character byte) Token {
	return Token{Type: tokenType, Literal: string(character)}
}

func Eof() Token {
	return Token{Type: EOF, Literal: ""}
}
