package lexer
import "github.com/arialdomartini/monkey/pkg/token"

type Lexer struct {
	input string
	position int
	readPosition int
	character byte
}

func New(input string) *Lexer {
	return &Lexer{input:input}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.character {
	case '=':
		tok = newToken(token.ASSIGN, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
