package lexer

import (
	"testing"

	"github.com/matthewrosse/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
  input := `=+(){},;`

  tests := []struct {
    expectedType  token.TokenType
    expectedLiteral string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
  }

  lexer = New(input)

  for i, testToken := range tests {
    token = lexer.NextToken()

    if token.Type != testToken.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q\n", i, testToken.expectedType, token.Type)
    }

    if token.Literal != testToken.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q\n", i, testToken.expectedLiteral, token.Literal)
    }
  }
}
