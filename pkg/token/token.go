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

	BANG      = "BANG"
	ASTERISK  = "ASTERISK"
	MINUS     = "MINUS"
	SLASH     = "SLASH"

	LT        = "LT"
	GT        = "GT"

	IF        = "IF"
	ELSE	  = "ELSE"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	RETURN    = "RETURN"

	EQ        = "EQ"
	NOT_EQ    = "NOT_EQ"
)

func Create(tokenType TokenType, character byte) Token {
	return CreateS(tokenType, string(character))
}
func CreateS(tokenType TokenType, s string) Token {
	return Token{Type: tokenType, Literal: s}
}

func Eof() Token {
	return Token{Type: EOF, Literal: ""}
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	} else {
		return IDENT
	}
}
