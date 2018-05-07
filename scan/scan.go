package scan

import (
	"BruteForceSimpleC/token"
	"unicode"
)

type Scanner struct {
	ch  	rune
	offset  int
	roffset int
	src 	string
	file 	*token.File
}

func (s *Scanner) Init(file *token.File, src string) {
	s.file = file
	s.offset, s.roffset = 0, 0
	s.src = src
	s.file.AddLine(s.offset)

	s.next()
}

func (s *Scanner) Scan() (lit string, tok token.Token, pos token.Pos) {
	s.skipWhitespace()

	if unicode.IsLetter(s.ch) || s.ch == '_'{
		return s.scanIdent()
	}

	if unicode.IsDigit(s.ch) {
		return s.scanNumber()
	}

	if s.ch == '"' {
		return s.scanString()
	}


	ch := s.ch //Save Current
	lit, pos = string(s.ch), s.file.Pos(s.offset)
	s.next()//Look ahead
	switch ch {
	case '+':
		tok = token.OP_PLUS
	case '-':
		tok = token.OP_MINUS
	case '*':
		tok = token.OP_MULTIPLY
	case '/':
		tok = token.OP_DIVIDE
	case '=':
		tok = s.selectToken('=', token.OP_EQ, token.OP_ASSIGN)
	case '<':
		tok = s.selectToken('=', token.OP_LE, token.OP_LT)
	case '>':
		tok = s.selectToken('=', token.OP_GE, token.OP_GT)
	case ',':
		tok = token.COMMA
	case ';':
		tok = token.SEMI
	case '(':
		tok = token.LPAR
	case ')':
		tok = token.RPAR
	case '{':
		tok = token.LBRACE
	case '}':
		tok = token.RBRACE


	default:
		if s.offset >= len(s.src)-1 {
			tok = token.EOF
		} else {
			tok = token.ILLEGAL
		}
	}

	return
}

func (s *Scanner) next() {
	s.ch = rune(0)
	if s.roffset < len(s.src) {
		s.offset = s.roffset
		s.ch = rune(s.src[s.offset])
		if s.ch == '\n' {
			s.file.AddLine(s.offset)
		}
		s.roffset++
	}
}

func (s *Scanner) scanIdent() (string, token.Token, token.Pos) {
	start := s.offset

	for unicode.IsLetter(s.ch) || unicode.IsDigit(s.ch) || s.ch == '_' {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	lit := s.src[start:offset]
	return lit, token.Lookup(lit), s.file.Pos(start)
}

func (s *Scanner) scanNumber() (string, token.Token, token.Pos) {
	var numType token.Token
	start := s.offset

	for unicode.IsDigit(s.ch) {
		s.next()
	}
	if s.ch == '.'{
		numType = token.FLOAT
		s.next()
		for unicode.IsDigit(s.ch) {
			s.next()
		}
	} else {
		numType = token.INTEGER
	}
	offset := s.offset
	lit := s.src[start:offset]
	return lit, numType, s.file.Pos(start)
}

func (s *Scanner) scanString() (string, token.Token, token.Pos) {
	start := s.offset
	s.next() //offset double quote
	for s.ch != '"' {
		s.next()
	}
	s.next()
	offset := s.offset
	lit := s.src[start:offset]
	return lit, token.STRING, s.file.Pos(start)
}

//select Token determines multi-character operators
func (s *Scanner) selectToken(r rune, a, b token.Token) token.Token  {
	if s.ch == r {
		s.next()
		return a
	}
	return b
}

func (s *Scanner) skipWhitespace() {
	for unicode.IsSpace(s.ch) {
		s.next()
	}
}