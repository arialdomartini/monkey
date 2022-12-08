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
	character := l.character
	token := token.Parse(character)
	l.readChar()
	return token
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
