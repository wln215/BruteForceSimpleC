package parse

import (
	"BruteForceSimpleC/token"
	"BruteForceSimpleC/scan"
	"BruteForceSimpleC/ast"

)

type parser struct {
	file 	*token.File
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

func (p *parser) init(file *token.File, fname, src string, s *ast.Scope) {
	if s == nil {
		s = ast.NewScope(nil)
	}
	p.file = file
	p.scanner.Init(file, src)
	p.listok = false
	p.curScope = s
	p.topScope = p.curScope
	p.next()
}

func (p *parser) next() {
	p.lit, p.tok, p.pos = p.scanner.Scan()
}

// Scoping Support

func (p *parser) openScope() {
	p.topScope = ast.NewScope(p.topScope)
}

func (p *parser) closeScope() {
	p.topScope = p.topScope.Parent
}


//File serves as the top level
func (p *parser) parseFile() *ast.File {
	topLevels := make([]*ast.Program, 0)
	for p.tok != token.EOF {
		p.parseProgram()
	}
	return &ast.File{Programs:topLevels, Scope:nil}
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


func (p *parser) parseVarList() ast.VarList {
	var list []*ast.Ident
	startPos := p.pos

	list = append(list, p.parseIdent())
	for p.tok == ',' {
		p.next()
		list = append(list, p.parseIdent())
	}
	return ast.VarList{ZerothPos:startPos, List:list}
}

func (p *parser) parseStmt() (s ast.Stmt) {
	switch p.tok {
	case token.IDENT, '-', token.INTEGER, token.FLOAT, '(':
		//Tokens that may start an expression
		s = p.parseStmt()
		p.expect(';')
	case token.KW_IF:
		p.parseIfStmt()
	case token.KW_WHILE:
		p.parseWhileStmt()
	case token.KW_READ:
		p.parseReadStmt()
		p.expect(';')
	case token.KW_WRITE:
		p.parseWriteExprList()
		p.expect(';')
	case token.KW_RETURN:
		p.parseReturnExpr()
		p.expect(';')
	case '{':
		p.parseBlock()
	default:
		//no statement found

		p.addError("InvalidStatement")
		s = parseBadStatement()
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
	if p.tok == token.KW_ELSE {
		elseStmt := p.parseElseStmt(ifKey)
		return &ast.IfStmt {Cond: cond, IfStmt:ifStmt, ElseStmt:elseStmt }
	}
	return &ast.IfStmt {ifKey, cond, ifStmt, nil }
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
	return &ast.WhileStmt{While:whileKey, Cond:cond, Loop:loop}
}

func (p *parser) parseReadStmt() *ast.ReadStmt {
	readPos := p.pos
	list := p.parseVarList()

	return &ast.ReadStmt{readPos, list}
}

func (p *parser) parseWriteExprList() *ast.WriteStmt {

}
func (p *parser) parseReturnExpr() *ast.ReturnStmt {

}

func (p *parser) parseBlock() ast.BlockStmt {
	open := p.pos
	var block []ast.Stmt

	for p.tok != '}' {
		block = append(block, p.parseStmt())
	}


	return ast.BlockStmt{open, block, p.pos}
}



func (p *parser) parseFactor() *ast.Factor {
	switch p.tok {
	case token.IDENT:
		obj := p.curScope.Lookup(p.tok.String())
		if obj.Kind == ast.Function {
			node := parseCall()
			return &ast.Factor{FacPos: node.Entry}
		}
		node := p.parseIdent()
		return &ast.Factor{FacPos:node.NamePos, Type:node.Object.Type, IsNeg:false}
	case token.INTEGER, token.FLOAT:
		node := p.parseBasicLit()
		return &ast.Factor{FacPos:node.LitPos, Type:node.Type, IsNeg:false}
	case token.RPAR:
		node := p.parseParenExpr()
		return &ast.Factor{FacPos:node.Lparen, Type:node.X.Type, IsNeg:false}
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
	return &ast.Factor{FacPos: node.FacPos, Type: node.Type, IsNeg: !node.IsNeg}
}

func (p *parser) parseTerm() *ast.Term {
	var lhs *ast.Factor
	if p.tok == '-' {
		lhs = p.parseNeg()
	}
	var tail *ast.BinaryExpr
	p.next()
	for p.tok == '*' || p.tok == '/'{
		tail = p.parseMulOp(lhs)
		p.next()
	}

	return &ast.Term{TermPos:lhs.FacPos, Value:tail}
}

//TODO chained operator support
func (p *parser) parseMulOp(lhs *ast.Factor) *ast.BinaryExpr{
	pos := p.pos
	operator := p.tok

	p.next()
	var rhs *ast.Factor
	if p.tok == '-' {
		rhs = p.parseNeg()
	} else {
		rhs = p.parseFactor()
	}

	return &ast.BinaryExpr{Lhs:lhs, OpPos:pos, Op:operator, Rhs:rhs}
}


func (p *parser) parseExpr1() *ast.Term {
	lhs := p.parseTerm()
	p.next()
	var tail *ast.BinaryExpr
	for p.tok == '+' || p.tok == '-' {
		tail = p.parseAddOp(lhs)
		p.next()
	}
	return &ast.Term{TermPos:lhs.TermPos, Value:tail}


}

func (p *parser) parseAddOp(lhs *ast.Term) *ast.BinaryExpr {
	pos := p.pos
	operator := p.tok
	p.next()

	rhs := p.parseTerm()


	return &ast.BinaryExpr{Lhs:lhs, Op:operator, OpPos:pos, Rhs:rhs}
}


func (p *parser) parseExpr() ast.Expr {
	if p.tok != token.IDENT {
		return p.parseExpr1()
	} else {
		lhs := p.parseIdent()
		p.next()

		var x *ast.BinaryExpr
		switch p.tok {
		case '=':
			x = p.parseAssign(lhs)
		case '+', '-':
			x = p.parseAddOp(&ast.Term{TermPos: lhs.NamePos, Value:lhs})
		case '*', '/':
			x = p.parseMulOp(&ast.Factor{FacPos: lhs.NamePos, Value:lhs})
		}

		return x
	}
}

func (p *parser) parseAssign(lhs *ast.Ident) *ast.BinaryExpr {
	op := p.tok
	opPos := p.pos
	rhs := p.parseExpr()
	return &ast.BinaryExpr{Lhs:lhs, OpPos:opPos, Op:op, Rhs:rhs}
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


func (p *parser) parseCall() *ast.CallExpr{

}

func (p *parser) parseBasicLit() *ast.BasicLit {
	pos, tok, lit := p.pos, p.tok, p.lit
	p.next()
	return &ast.BasicLit{LitPos: pos, Type: tok, Lit: lit}
}

func (p *parser) parseType() ast.VarList {
	return p.parseVarList()
}

func (p *parser) parseIdent() *ast.Ident {
	name := p.lit
	return &ast.Ident{NamePos: p.expect(token.IDENT), Name:name}
}

