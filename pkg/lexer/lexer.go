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
	l.skipWhiteSpaces()
	character := l.character
	token := l.parse(character)
	return token
}

func (l *Lexer) skipWhiteSpaces() {
	for l.character == ' ' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) tryEvaluateSecondCharacter(character byte, firstType token.TokenType, secondCharacter byte, secondType token.TokenType) token.Token{
	if(l.peekChar() == secondCharacter) {
		first := string(character)
		l.readChar()
		second := string(l.character)
		return token.CreateS(secondType, first + second)
	} else {
		return token.Create(firstType, character)
	}
}


func (l *Lexer) parse(character byte) token.Token {
	var t token.Token

	switch character {
	case '=':
		t = l.tryEvaluateSecondCharacter(character, token.ASSIGN, '=', token.EQ)
	case ';':
		t = token.Create(token.SEMICOLON, character)
	case '(':
		t = token.Create(token.LPAREN, character)
	case ')':
		t = token.Create(token.RPAREN, character)
	case ',':
		t = token.Create(token.COMMA, character)
	case '+':
		t = token.Create(token.PLUS, character)
	case '{':
		t = token.Create(token.LBRACE, character)
	case '}':
		t = token.Create(token.RBRACE, character)
	case '!':
		t = l.tryEvaluateSecondCharacter(character, token.BANG, '=', token.NOT_EQ)
	case '-':
		t = token.Create(token.MINUS, character)
	case '/':
		t = token.Create(token.SLASH, character)
	case '*':
		t = token.Create(token.ASTERISK, character)
	case '<':
		t = token.Create(token.LT, character)
	case '>':
		t = token.Create(token.GT, character)
	case 0:
		t = token.Eof()
	default:
		if isLetter(character) {
			identifier := l.readIdentifier()
			t.Literal = identifier
			t.Type = token.LookupIdentifier(identifier)
			return t
		} else if isDigit(character) {
			number := l.readNumber()
			t.Literal = number
			t.Type = token.INT
			return t
		} else {
			t = token.Create(token.ILLEGAL, character)
		}
	}
	l.readChar()
	return t
}

func (l *Lexer) readIdentifier() string {
	return l.readWhile(isLetter)
}


func (l *Lexer) readNumber() string {
	return l.readWhile(isDigit)
}

func (l *Lexer) readWhile(predicate func(byte) bool) string {
	position := l.position
	for predicate(l.character) {
		l.readChar()
	}
	return l.input[position : l.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && 'z' >= character || 'A' <= character && 'Z' >= character || '_' == character
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
