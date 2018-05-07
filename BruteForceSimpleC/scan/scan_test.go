package scan_test

import (
	"testing"

	"BruteForceSimpleC/scan"
	"BruteForceSimpleC/token"

)

func test_handler(t *testing.T, src string, expected []token.Token) {
	var s scan.Scanner
	s.Init(token.NewFile("", src), src)
	lit, tok, pos := s.Scan()
	for i := 0; tok != token.EOF; i++ {
		if tok != expected[i] {
			t.Fatal(pos, "Expected:", expected[i], "Got:", tok, lit)
		}
		lit, tok, pos = s.Scan()
	}
}

func TestIdentifier(t *testing.T) {
	src := "a A z23 Zasdf"
	expected := []token.Token{
		token.IDENT,
		token.IDENT,
		token.IDENT,
		token.IDENT,
		token.EOF,
	}

	test_handler(t, src, expected)
}

func TestNumber(t *testing.T) {
	src := "9"
	expected := []token.Token{
		token.INTEGER,
		token.EOF,
	}

	test_handler(t, src, expected)
}

func TestScan(t *testing.T) {
	src := "int a (int b){ c = b + a * 4;}"
	expected := []token.Token{
		token.KW_INT,
		token.IDENT,
		token.LPAR,
		token.KW_INT,
		token.IDENT,
		token.RPAR,
		token.LBRACE,
		token.IDENT,
		token.OP_ASSIGN,
		token.IDENT,
		token.OP_PLUS,
		token.IDENT,
		token.OP_MULTIPLY,
		token.INTEGER,
		token.SEMI,
		token.RBRACE,
		token.EOF,
	}
	test_handler(t, src, expected)
}

func TestScanAllTokens(t *testing.T) {
	src := "()+-*/ 1 12\t 12345 1234.56789 | a as ! < <=! = == > >= &" +
		"| ; \\ \r :"
	expected := []token.Token{
		token.LPAR,
		token.RPAR,
		token.OP_PLUS,
		token.OP_MINUS,
		token.OP_MULTIPLY,
		token.OP_DIVIDE,
		token.INTEGER,
		token.INTEGER,
		token.INTEGER,
		token.FLOAT,
		token.ILLEGAL,
		token.IDENT,
		token.IDENT,
		token.ILLEGAL,
		token.OP_LT,
		token.OP_LE,
		token.ILLEGAL,
		token.OP_ASSIGN,
		token.OP_EQ,
		token.OP_GT,
		token.OP_GE,
		token.ILLEGAL,
		token.ILLEGAL,
		token.SEMI,
		token.ILLEGAL,
		token.ILLEGAL,
		token.EOF,
	}
	test_handler(t, src, expected)
}