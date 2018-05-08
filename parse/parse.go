package parse

import (
	"BruteForceSimpleC/token"
	"BruteForceSimpleC/scan"
	"BruteForceSimpleC/ast"

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

	exprLev int
	inRhs 	bool
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


//File serves as the top level
func (p *parser) parseFile() *ast.File {
	topLevels := make([]*ast.Program, 0)
	for p.tok != token.EOF {
		dataType := p.tok
		p.next()
		nom := p.parseIdent()


		}
	}
	return &ast.File{Scope: p.topScope}
}


func (p *parser) parseProgram() *ast.Expr {
	var expr *ast.Expr

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
	var list []*ast.Ident
	list = append(list, p.parseIdent())
	for p.tok == ',' {
		p.next()
		list = append(list, p.parseIdent())
	}
	return list
}

func (p *parser) parseStmt() (s ast.Stmt) {
	switch p.tok {
	case token.IDENT, '-', token.INTEGER, token.FLOAT, '(':
		//Tokens that may start an expression
		s = p.parseSimpleStmt
	case token.KW_IF:
		p.parseIfStmt()
	case token.KW_WHILE:
		p.parseWhileStmt()
	case token.KW_READ:
		p.parseReadStmt()
	case token.KW_WRITE:
		p.parseWriteExprList()
	case token.KW_RETURN:
		p.parseReturnExpr()
	case '{':
		p.parseBlock()
	default:
		//no statement found
		pos := p.pos
		p.errorExpected(pos, "statement")
		s = ast.BadStmt{From: pos, To: p.pos}
		}

		return
}

func (p *parser) parseIfStmt() *ast.IfStmt {
	ifKey := p.expect(token.KW_IF)
	p.openScope()
	defer p.closeScope()

	p.expect('(')
	cond := p.parseBoolExpr()
	p.expect(')')
	ifStmt := p.parseStmt()
	if p.tok == token.KW_ELSE { //Else will always apply to innermost if
		elseStmt := p.parseElseStmt(ifKey)
		return &ast.IfStmt {Cond: cond, IfStmt:ifStmt, ElseStmt:elseStmt }
	}
	return &ast.IfStmt {Cond: cond, IfStmt:ifStmt, ElseStmt:nil }
}

func (p *parser) parseElseStmt(openIf token.Pos) *ast.Expr {
	elseKey := p.pos
	p.next()
	elseStmt := p.parseStmt()
	return &ast.ElseStmt{}

}
func (p *parser) parseWhileStmt() *ast.WhileStmt {
	whileKey := p.pos
	p.next()
	cond := p.parseBoolExpr()
	loop := p.parseStmt()
	return &ast.WhileStmt{Entry:whileKey, Cond:cond, Loop:loop}
}

func (p *parser) parseReadStmt() *ast.ReadStmt {

}

func (p *parser) parseWriteExprList() *ast.WriteStmt {

}
func (p *parser) parseReturnExpr() *ast.ReturnStmt {

}





func (p *parser) parseFactor() *ast.Factor {
	switch p.tok {
	case token.IDENT:
		obj := p.curScope.Lookup(p.tok.String())
		if obj.Kind == ast. {
			node := parseCall()
			return &ast.Factor{FacPos: node.Entry}
		}
		node := p.parseIdent()
		return &ast.Factor{FacPos:node.NamePos, Kind:node.Type, IsNeg:false}
	case token.INTEGER, token.FLOAT:
		node := p.parseBasicLit()
		return &ast.Factor{FacPos:node.LitPos, Kind:node.Kind, IsNeg:false}
	case token.RPAR:
		p.next()
		node := p.parseExpr()
		return &ast.Factor{FacPos:node.ExprPos, Kind:node.Kind, IsNeg:false}
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
	return &ast.Factor{FacPos: node.FacPos, Kind: node.Kind, IsNeg: !node.IsNeg}
}

func (p *parser) parseTerm() *ast.Term {
	var lhs *ast.Factor
	if p.tok == '-' {
		lhs = p.parseNeg()
	}
	var tail *ast.MulOp
	p.next()
	for p.tok == '*' || p.tok == '/'{
		tail = p.parseMulOp(lhs)
		p.next()
	}
	return &ast.Term{TermPos:lhs.FacPos, Value:tail}
}

//TODO chained operator support
func (p *parser) parseMulOp(lhs *ast.Factor) *ast.MulOp{
	pos := p.pos
	operator := p.tok

	oprands := make([]*ast.Factor, 0)
	oprands = append(oprands, lhs)

	p.next()
	var rhs *ast.Factor
	if p.tok == '-' {
		rhs = p.parseNeg()
	} else {
		rhs = p.parseFactor()
	}
	oprands = append(oprands, rhs)



	return &ast.MulOp{OpPos:pos, Op:operator, List:oprands}
}

func (p *parser) parseBoolExpr() ast.BoolExpr {
	lhs := p.parseExpr1()
	p.next()
	opType := p.tok
	opPos  := p.pos
	switch opType {
	case token.OP_EQ, token.OP_GT, token.OP_LT, token.OP_GE, token.OP_LE:
		p.next()
		rhs := p.parseExpr1()
		p.expect(')')
		return ast.BoolExpr{Lhs:lhs, Op:opType, OpPos:opPos, Rhs:rhs}
	case ')':
		p.addError("no boolean operator found")
		return ast.BoolExpr{Lhs:lhs, Op:nil, OpPos:nil}
	default:
		p.addError("no valid boolean expression")
		return ast.BoolExpr{Lhs:lhs, Op:nil, OpPos:nil}

	}

}

func (p *parser) parseExpr1() *ast.Expr {
	lhs := p.parseTerm()
	p.next()
	for p.tok == '+' || p.tok == '-' {
		tail = p.parseAddOp(lhs)
		p.next()
	}
	return &ast.Term{TermPos:lhs.TermPos, Value:tail}


}

func (p *parser) parseAddOp(lhs *ast.Node) *ast.Expr1 {
	pos := p.pos
	operator := p.tok

	oprands := make([]*ast.Term, 0)
	oprands = append(oprands, lhs)

	p.next()

	rhs := p.parseTerm()
	oprands = append(oprands, rhs)

	return &ast.Expr1{ExPos:pos, Op:operator, List:oprands}
}


func (p *parser) parseExpr() *ast.Expr {
	if p.tok != token.IDENT {
		p.parseExpr1()
	} else {
		lhs := p.parseIdent()
		lht := ast.Term{}
		p.next()
		switch p.tok {
		case '=':
			p.parseAssign(lhs)
		case '+', '-':
			p.parseAddOp(lhs)
		case '*', '/':
			p.parseMulOp(lhs)
		}


	}
}


func (p *parser) parseCall() *ast.Expr{

}

func (p *parser) parseBasicLit() *ast.BasicLiteral {
	pos, tok, lit := p.pos, p.tok, p.lit
	p.next()
	return &ast.BasicLiteral{LitPos: pos, Kind: tok, Lit: lit}
}

func (p *parser) parseType() *ast.VarList {
	p.parseVarList()
}

func (p *parser) parseIdent() *ast.Ident {
	name := p.lit
	return &ast.Ident{NamePos: p.expect(token.IDENT)}
}

