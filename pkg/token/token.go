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
