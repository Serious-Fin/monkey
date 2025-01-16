package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // point to current char
	readPosition int  // point to next char to read
	ch           byte // current char at position
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var newTok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			literal := string(ch) + string(lexer.ch)
			newTok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			newTok = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		newTok = newToken(token.PLUS, lexer.ch)
	case '(':
		newTok = newToken(token.LPAREN, lexer.ch)
	case ')':
		newTok = newToken(token.RPAREN, lexer.ch)
	case '{':
		newTok = newToken(token.LBRACE, lexer.ch)
	case '}':
		newTok = newToken(token.RBRACE, lexer.ch)
	case ',':
		newTok = newToken(token.COMMA, lexer.ch)
	case ';':
		newTok = newToken(token.SEMICOLON, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			literal := string(ch) + string(lexer.ch)
			newTok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			newTok = newToken(token.BANG, lexer.ch)
		}
	case '-':
		newTok = newToken(token.MINUS, lexer.ch)
	case '/':
		newTok = newToken(token.SLASH, lexer.ch)
	case '*':
		newTok = newToken(token.ASTERISK, lexer.ch)
	case '<':
		newTok = newToken(token.LT, lexer.ch)
	case '>':
		newTok = newToken(token.GT, lexer.ch)
	case 0:
		newTok.Type = token.EOF
		newTok.Literal = ""
	default:
		if isLetter(lexer.ch) {
			newTok.Literal = lexer.readIdentifier()
			newTok.Type = token.LookupIdent(newTok.Literal)
			return newTok
		} else if isDigit(lexer.ch) {
			newTok.Type = token.INT
			newTok.Literal = lexer.readNumber()
			return newTok
		} else {
			newTok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return newTok
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readIdentifier() string {
	startPos := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[startPos:lexer.position]
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) readNumber() string {
	startPos := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[startPos:lexer.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
		return
	}
	lexer.ch = lexer.input[lexer.readPosition]
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.readPosition]
}
