package lexer

import (
	"testing"

	"monkey/token"
)

func TestOperatorAndSeparatorLexing(t *testing.T) {
	input := `=+(){},;`
	expectedTokens := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	lexer := NewLexer(input)
	ValidateLexerTokens(lexer, expectedTokens, t)
}

func TestNextTokenOnRealisticCode(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
}

let result = add(five, ten);
`
	expectedTokens := []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "result"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	lexer := NewLexer(input)
	ValidateLexerTokens(lexer, expectedTokens, t)
}

func ValidateLexerTokens(lexer *Lexer, expectedTokens []token.Token, t *testing.T) {
	for index, expectedToken := range expectedTokens {
		tok := lexer.NextToken()
		if expectedToken.Type != tok.Type {
			t.Fatalf("token[%d] - tokentype wrong. expected \"%s\", got \"%s\"", index, expectedToken.Type, tok.Type)
		}

		if expectedToken.Literal != tok.Literal {
			t.Fatalf("token[%d] - literal wrong. expected \"%s\", got \"%s\"", index, expectedToken.Literal, tok.Literal)
		}
	}
}
