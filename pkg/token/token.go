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


func Create(tokenType TokenType, character byte) Token {
	return Token{Type: tokenType, Literal: string(character)}
}

func Eof() Token {
	return Token{Type: EOF, Literal: ""}
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	} else {
		return IDENT
	}
}
