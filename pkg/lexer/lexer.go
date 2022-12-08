package lexer
import "github.com/arialdomartini/monkey/pkg/token"

type Lexer struct {
	input string
	position int
	readPosition int
	character byte
}

func New(input string) *Lexer {
	l := &Lexer{input:input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.character {
	case '=':
		tok = token.New(token.ASSIGN, l.character)
	case ';':
		tok = token.New(token.SEMICOLON, l.character)
	case '(':
		tok = token.New(token.LPAREN, l.character)
	case ')':
		tok = token.New(token.RPAREN, l.character)
	case ',':
		tok = token.New(token.COMMA, l.character)
	case '+':
		tok = token.New(token.PLUS, l.character)
	case '{':
		tok = token.New(token.LBRACE, l.character)
	case '}':
		tok = token.New(token.RBRACE, l.character)
	case 0:
		tok = token.Eof()
	}

	l.readChar()
	return tok
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
