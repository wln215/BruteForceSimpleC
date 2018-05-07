package token_test

import (
	"testing"

	"BruteForceSimpleC/token"
)

var test_expr = "(+ 2 3)\n(- 5 4)"

func TestFilePosition(t *testing.T) {
	//var tests = []struct {
	//	col, row int
	//	pos 	 token.Pos
	//}{
	//	{1, 1, token.Pos(1)},
	//	{1, 2, token.Pos(8)},
	//	{7, 2, token.Pos(14)},
	//}
	f := token.NewFile("", "")
	f.AddLine(int(token.Pos(1)))
	p := f.Position(token.Pos(1))
	if p.String() != "1:1" {
		t.Fatal("Nameless file, expected 1:1, got ", p.String())
	}

	f = token.NewFile("test.sic", "")
	f.AddLine(1)
	p = f.Position(token.Pos(1))
	if p.String() != "test.sic:1:1" {
		t.Fatal("Nameless file: Expected: test.calc:1:1, Got:", p.String())
	}
/*
	f = token.NewFile("test", test_expr)
	f.AddLine(6)
	f.AddLine(15)
	for _, v := range tests {
		p := f.Position(v.pos)
		if p.Col != v.col || p.Row != v.row {
			t.Fatal("For:", v.pos, "Expected:", v.col, "and", v.row, "Got:",
				p.Col, "and", p.Row)
		}
	}
*/
}

func TestLookup(t *testing.T) {
	var tests = []struct {
		str string
		tok token.Token
	}{
		{"+", token.OP_PLUS},
		{"%", token.OP_DIVIDE},
		{"EOF", token.EOF},
		{"Integer", token.INTEGER},
		{"Comment", token.COMMENT},
		{"", token.ILLEGAL},
	}

	for i, v := range tests {
		if res := token.Lookup(v.str); res != v.tok {
			t.Fatal(i, "- Expected:", v.tok, "Got:", res)
		}
	}
}

func TestIsLiteral(t *testing.T) {
	var tests = []struct {
		tok token.Token
		exp bool
	}{
		{token.OP_PLUS, false},
		{token.OP_DIVIDE, false},
		{token.EOF, false},
		{token.INTEGER, true},
		{token.COMMENT, false},
	}

	for _, v := range tests {
		if res := v.tok.IsLiteral(); res != v.exp {
			t.Fatal(v.tok, "- Expected:", v.exp, "Got:", res)
		}
	}
}

func TestIsOperator(t *testing.T) {
	var tests = []struct {
		tok token.Token
		exp bool
	}{
		{token.OP_PLUS, false},
		{token.OP_DIVIDE, false},
		{token.EOF, false},
		{token.INTEGER, true},
		{token.COMMENT, false},
	}

	for _, v := range tests {
		if res := v.tok.IsOperator(); res != v.exp {
			t.Fatal(v.tok, "- Expected:", v.exp, "Got:", res)
		}
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		str string
		tok token.Token
	}{
		{"+", token.OP_PLUS},
		{"/", token.OP_DIVIDE},
		{"EOF", token.EOF},
		{"Integer", token.INTEGER},
		{"Comment", token.COMMENT},
	}

	for i, v := range tests {
		if res := v.tok.String(); res != v.str {
			t.Fatal(i, "- Expected:", v.str, "Got:", res)
		}
	}
}

func TestValid(t *testing.T) {
	var tests = []struct {
		tok token.Token
		exp bool
	}{
		{token.OP_PLUS, true},
		{token.OP_DIVIDE, true},
		{token.EOF, true},
		{token.INTEGER, true},
		{token.COMMENT, true},
		{token.Token(-1), false},
		{token.Token(999999), false},
	}

	for _, v := range tests {
		if res := v.tok.Valid(); res != v.exp {
			t.Fatal(v.tok, "- Expected:", v.exp, "Got:", res)
		}
	}
}

