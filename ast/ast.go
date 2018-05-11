package ast

import (
	"BruteForceSimpleC/token"
)

/* Interfaces */

type Node interface{
	Pos() token.Pos
	End() token.Pos
}
//All Expressions implement Expr interface
type Expr interface { //Theoretically cannot have a root
	Node
	exprNode()
}

//All Statements implement Stmt interface
type Stmt interface {
	Node
	stmtNode()
}

//All Declarations implement Decl interface
type Decl interface {
	Node
	declNode()
}

/* Expressions and types */



type BadExpr struct {
	From, To token.Pos
}

type Ident struct {
	NamePos token.Pos
	Name 	string

	//Object  *Object //nil if keyword
}

type VarList struct {
	ZerothPos token.Pos
	List 	  []*Ident
}

type BasicLit struct {
	LitPos token.Pos
	Type   token.Token
	Lit    string
}

type FuncLit struct {
	Type *FunctionType
	Body *BlockStmt 
}

type ParenExpr struct {
	Lparen  token.Pos
	X 		Expr
	Rparen  token.Pos
}

type CallExpr struct {
	Fun 		string
	Lparen		token.Pos
	Args 		Expr
	Rparen 		token.Pos
}

type BinaryExpr struct {
	Lhs   	Expr
	Op    	token.Token
	OpPos 	token.Pos
	Rhs 	Expr
}

type BoolExpr struct {
	Lhs		Expr
	OpPos 	token.Pos
	Op 		token.Token
	Rhs 	Expr
}

type Expr1 struct {
	ExPos 	token.Pos
	Name 	*Ident
	Value 	Term
}


type Term struct {
	TermPos token.Pos
	Value   Expr

}

type Factor struct {
	FacPos  token.Pos
	IsNeg   bool
	Type 	Expr
	Value   Expr
}

type FunctionType struct {
	Func   token.Pos
	Params *Ident
	Type   Expr
}







func (x *BadExpr) Pos() token.Pos 		{return x.From}
func (x *Ident) Pos() token.Pos 		{return x.NamePos}
func (x *VarList) Pos() token.Pos 		{return x.ZerothPos}
func (x *BasicLit) Pos() token.Pos 		{return x.LitPos}
func (x *FuncLit) Pos() token.Pos 		{return x.Type.Pos()}
func (x *ParenExpr) Pos() token.Pos 	{return x.Lparen}
func (x *CallExpr) Pos() token.Pos 		{return x.Lparen - token.Pos(len(x.Fun))}
func (x *BinaryExpr) Pos() token.Pos 	{return x.Lhs.Pos()}
func (x *Expr1) Pos() token.Pos 		{return x.ExPos}
func (x *Term) Pos() token.Pos 			{return x.TermPos}
func (x *Factor) Pos() token.Pos		{return x.FacPos}
func (t *FunctionType) Pos() token.Pos  {return t.Func}

func (x *BadExpr) End() token.Pos 		{return x.To}
func (x *Ident) End() token.Pos 		{return token.Pos(int(x.NamePos) + len(x.Name))}
func (x *VarList) End() token.Pos 		{return token.Pos(int(x.ZerothPos) + len(x.List))}
func (x *BasicLit) End() token.Pos 		{return token.Pos(int(x.LitPos) + len(x.Lit))}
func (x *FuncLit) End() token.Pos 		{return x.Body.End()}
func (x *ParenExpr) End() token.Pos 	{return x.Rparen+1}
func (x *CallExpr) End() token.Pos 		{return x.Rparen+1}
func (x *BinaryExpr) End() token.Pos 	{return x.Rhs.End()}
func (x *Term) End() token.Pos	 		{return x.Value.End()}
func (x *Factor) End() token.Pos 		{return x.Value.End()}
func (x *Expr1) End() token.Pos 		{return x.Value.End()}

func (x *BadExpr) exprNode()		{}
func (x *Ident) exprNode()			{}
func (x *VarList) exprNode()		{}
func (x *BasicLit) exprNode()		{}
func (x *FuncLit) exprNode()		{}
func (x *ParenExpr) exprNode()		{}
func (x *CallExpr) exprNode()		{}
func (x *BinaryExpr) exprNode()		{}
func (x *Expr1) exprNode()			{}
func (x *Term) exprNode()			{}
func (x *Factor) exprNode()			{}
func (x *FunctionType) exprNode()	{}

/* Convenience functions for Ident */

func NewIdent(name string) *Ident { return &Ident{token.NoPos, name}}

/* Statements */

type BadStmt struct {
	From, To token.Pos
}

// A DeclStmt node represents a declaration in a statement list.
type DeclStmt struct {
	Decl Decl // *GenDecl with CONST, TYPE, or VAR token
}

type ExprStmt struct {
	X Expr // expression
}

type AssignStmt struct {
	Lhs    []Expr
	TokPos token.Pos   // position of Tok
	Tok    token.Token // assignment token, DEFINE
	Rhs    []Expr
}

type IfStmt struct {
	IfKey 	 token.Pos
	Cond   	 BoolExpr
	IfStmt   Stmt
	Else 	 Stmt

}

type WhileStmt struct {
	While 	 token.Pos
	Cond 	 BoolExpr
	Loop 	 Stmt
}

type ReadStmt struct {
	Read 	token.Pos
	In 		VarList
}

type WriteStmt struct {
	Write 	token.Pos
	Out 	*WriteExprList
}

type WriteExprList struct {
	ListBeg token.Pos
	Expr 	[]Expr

}

type ReturnStmt struct {
	Return  token.Pos // position of "return" keyword
	Results []Expr    // result expressions; or nil
}

type BlockStmt struct {
	Lbrace token.Pos // position of "{"
	List   []Stmt
	Rbrace token.Pos // position of "}"
}

func (s *BadStmt) Pos()	token.Pos 		{return s.From}
func (s *DeclStmt) Pos() token.Pos 		{return s.Decl.Pos()}
func (s *ExprStmt) Pos() token.Pos 		{return s.Pos()}
func (s *AssignStmt) Pos() token.Pos 	{return s.TokPos}
func (s *IfStmt) Pos() token.Pos 		{return s.IfKey}
func (s *WhileStmt) Pos() token.Pos		{return s.While}
func (s *ReadStmt) Pos() token.Pos 		{return s.Read}
func (s *WriteStmt) Pos() token.Pos 	{return s.Write}
func (s *WriteExprList) Pos() token.Pos {return s.ListBeg}
func (s *ReturnStmt) Pos() token.Pos 	{return s.Return}
func (s *BlockStmt) Pos() token.Pos 	{return s.Lbrace}


func (s *BadStmt) End() token.Pos	 	{return s.To}
func (s *DeclStmt) End() token.Pos	 	{return s.Decl.Pos()}
func (s *ExprStmt) End() token.Pos	 	{return s.X.Pos()}
func (s *AssignStmt) End() token.Pos 	{return s.Rhs[len(s.Rhs)-1].End()}
func (s *IfStmt) End() token.Pos	 	{
	if s.Else != nil {
	return s.Else.End()
	}
	return s.IfStmt.End()}
func (s *WhileStmt) End() token.Pos	 	{return s.Loop.End()}
func (s *ReadStmt) End() token.Pos	 	{return s.In.List[len(s.In.List)-1].End()}
func (s *WriteStmt) End() token.Pos	 	{return s.Out.End()}
func (s *WriteExprList) End() token.Pos	{return s.ListBeg}
func (s *ReturnStmt) End() token.Pos	{
	if n := len(s.Results); n > 0 {
		return s.Results[n-1].End()
	}
	return s.Return + 6 // len("return"
}
func (s *BlockStmt) End() token.Pos	 	{return s.Rbrace}


func (s *BadStmt)  stmtNode()		{}
func (s *DeclStmt) stmtNode()		{}
func (s *ExprStmt) stmtNode()		{}
func (s *AssignStmt) stmtNode()		{}
func (s *IfStmt) stmtNode()			{}
func (s *WhileStmt) stmtNode()		{}
func (s *ReadStmt) stmtNode()		{}
func (s *WriteStmt) stmtNode()		{}
func (s *WriteExprList) stmtNode()	{}
func (s *ReturnStmt) stmtNode()		{}
func (s *BlockStmt) stmtNode()		{}

/* Declarations */
//A Spec node represents a constant, type, or variable declaration
type Spec interface {
	Node
	specNode()
}

//ValueSpec node represents a constant or variable declaration
type ValueSpec struct {
	Names  []Ident
	Type   Expr
	Values []Expr
}

type BadDecl struct {
	From, To token.Pos
}

type GeneralDecl struct {
	TypePos token.Pos
	Type 	token.Token
	List 	VarList
}



type File struct {
	Programs []*Program
	Scope 	 *Scope
}


// function_call represents data reduced by production:
//
//	function_call:
//	        ID '(' expr ')'  // kind 0
type FunctionCall struct {
	Name *Ident
	Args []Expr
}

// function_decl represents data reduced by production:
//
//	function_decl:
//	        kind ID '(' kind ')' ';'  // kind 0
type FunctionDecl struct {
	DeclPos   token.Pos
	Name 	  *Ident
	FuncType  token.Token
	ArgType   token.Token
	DeclEnd	  token.Pos
}


// function_def represents data reduced by production:
//
//	function_def:
//	        kind ID '(' kind ID ')' body  // kind 0
type FunctionDef struct {
	Define token.Pos
	Name   *Ident
	Type   *Ident
	Kind   token.Token
	Body   Expr
}



type Program struct {
	Entry 		token.Pos
	Decls 		[]*Decl

}

func (d *BadDecl) Pos() 	token.Pos {return d.From}
func (d *GeneralDecl) Pos() token.Pos {return d.TypePos}
func (d *FunctionCall) Pos() token.Pos {return d.Name.NamePos}
func (d *FunctionDecl) Pos() token.Pos {return d.DeclPos}
func (d *FunctionDef) Pos() token.Pos {return d.Define}
func (d *Program) Pos() token.Pos {return d.Entry}
//func (d *) Pos() token.Pos {return }


func (d *BadDecl) End() token.Pos	  {return d.To}
func (d *GeneralDecl) End() token.Pos {return d.List.End()}
func (d *FunctionCall) End() token.Pos {return d.Args[len(d.Args)-1].End()}
func (d *FunctionDecl) End() token.Pos {return d.DeclEnd}
func (d *FunctionDef) End() token.Pos {return d.Body.End()}
//func (d *Program) End() token.Pos {return d.*Decls[-1]}
//func (d *) End() token.Pos {return }

func (d *BadDecl) declNode() {}
func (d *GeneralDecl) declNode() {}
func (d *FunctionCall) declNode() {}
func (d *FunctionDecl) declNode() {}
func (d *FunctionDef) declNode() {}
func (d *Program) declNode() {}
//func (d *) declNode() {}



