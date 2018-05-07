package parse

import (
	"BruteForceSimpleC/token"
	"BruteForceSimpleC/scan"
	"BruteForceSimpleC/ast"
	"golang.org/x/tools/present"
)

type parser struct {
	file 	token.File
	errors	token.ErrorList
	scanner scan.Scanner
	listok  bool

	curScope 	*ast.Scope
	topScope	*ast.Scope

	pos token.Pos
	tok token.Token
	lit string
}

/* Utilities */
func (p *parser) addError(args...interface{}) {

}

func (p *parser) expect(tok token.Token) token.Pos {
	if p.tok != tok {
		p.addError("Expected '" + tok.String() + "' got '" + p.lit + "'")
		return p.pos
	}
	defer p.next()//happens last
	return p.pos
}

func (p *parser) init(file token.File, fname, src string, s *ast.Scope) {
	if s == nil {
		s = ast.NewScope(nil)
	}
	p.file = file
	p.scanner.Init(p.file, src)
	p.listok = false
	p.curScope = s
	p.topScope = p.curScope
	p.next()
}

func (p *parser) next() {
	p.lit, p.tok, p.pos = p.scanner.Scan()
}

func (p *parser) parseExprList(open token.Pos) ast.SubTree {
	p.listok = false
	var []*ast.expr
}

//File serves as the top level
func (p *parser) parseFile() *ast.File {
	for p.tok != token.EOF {
		p.parsePrograms()
	}

	return &ast.File{Scope: p.topScope}
}

func (p *parser) parsePrograms() *ast.Programs {

}
// program represents data reduced by productions:
//
//	program:
//	        /* empty */            // kind 0
//	|       program function_def   // kind 1
//	|       program decl           // kind 2
//	|       program function_decl  // kind 3
func (p *parser) parseProgram() *ast.SubTree {
	var expr *ast.SubTree

	if p.tok != token.KW_FLOAT || p.tok != token.KW_INT {
		p.addError("no function or declaration")
	}
	p.next()


}

// var_list represents data reduced by productions:
//
//	var_list:
//	        ID               // kind 0
//	|       var_list ',' ID  // kind 1

func (p *parser) parseVarList() []*ast.Ident {
	list:= make([]*ast.Ident, 0)
	for (p.tok == token.IDENT) || (p.tok == token.COMMA) {
		if p.tok == token.IDENT{
			list = append(list, p.parseIdent())
		} else {
			p.next() //skips comma
		}
	}
	return list
}

func (p *parser) parseFactor() *ast.Factor {
	switch p.tok {
	case token.IDENT:
		obj := p.curScope.Lookup(p.tok.String())
		if obj.Kind == ast. {
			node := parseCall()
			return &ast.Factor{Pos: node.}
		}
		node := p.parseIdent()
		return &ast.Factor{Pos:node.NamePos, Kind:node.Object, IsNeg:false}
	case token.INTEGER, token.FLOAT:
		node := p.parseBasicLit()
		return &ast.Factor{Pos:node.LitPos, Kind:node.Kind, IsNeg:false}
	case token.RPAR:
		p.next()
		node := p.parseExpr()
		return &ast.Factor{Pos:node.ExprPos, Kind:node.Kind, IsNeg:false}
	default:
		p.addError("No valid factors")
	}
	return nil
}

func (p *parser) parseNeg() *ast.Factor {
	p.next()
	//For handling multiple negations
	negate := true
	for p.tok == '-' {
		negate = !(negate)
	}
	node := p.parseFactor()
	if !negate{
		return node //Double negative
	}
	return &ast.Factor{Pos: node.Pos, Kind: node.Kind, IsNeg: !node.IsNeg}
}

func (p *parser) parseTerm() *ast.Term {
	var a, b *ast.Factor
	switch p.tok {
	case '-':
		a = p.parseNeg()
	}
}


func (p *parser) parseCall() *ast.SubTree{

}

func (p *parser) parseBasicLit() *ast.BasicLiteral {
	pos, tok, lit := p.pos, p.tok, p.lit
	p.next()
	return &ast.BasicLiteral{LitPos: pos, Kind: tok, Lit: lit}
}
func (p *parser) parseIdent() *ast.Ident {
	name := p.lit
	return &ast.Ident{NamePos: p.expect(token.IDENT)}
}