package ast

import (
	"BruteForceSimpleC/token"
)

/* Interfaces */

type Node interface{
	Pos() token.Pos
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

type  VarList struct {
	List []*Ident
	Type Type
}

type BadExpr struct {
	From, To token.Pos
}

type Ident struct {
	NamePos token.Pos
	Name 	string
	Object  *Object //nil if keyword
}

type BasicLiteral struct {
	LitPos token.Pos
	Kind   token.Token
	Lit    string
}

type FuncLit struct {
	Type *FunctionType
	Body *BlockStmt 
}

type ParenExpr struct {
	Lparen  token.Pos
	X 		Expr
}

type CallExpr struct {
	Fun 		Expr
	Lparen		token.Pos
	Args 		[]Expr
	Rparen 		token.Pos
}

type BinaryExpr struct {
	Lhs   	*Expr
	Op    	token.Token
	OpPos 	token.Pos
	Rhs 	*Expr
}

type Expr1 struct {
	ExPos 	token.Pos
	Name 	*Ident
	Value 	Expr
}

/* Convenience functions for Ident */

func NewIdent(name string) *Ident { return &Ident{token.NoPos, name, nil}}

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
	Cond   	 BinaryExpr
	IfStmt   *Expr
	ElseStmt *Expr

}

type WhileStmt struct {
	Entry 	 token.Pos
	Cond 	 BinaryExpr
	Loop 	 *Expr
}

type ReadStmt struct {
	Read 	token.Pos
	In 		VarList
}

type WriteStmt struct {
	Write 	token.Pos
	Out 	WriteExprList
}

type ReturnStmt struct {
	Return  token.Pos // position of "return" keyword
	Results []Expr    // result expressions; or nil
}

type Term struct {
	TermPos token.Pos
	Kind    Kind
	Value   *Expr

}


/* Declarations */
//A Spec node represents a constant, type, or variable declaration
type Spec interface {
	Node
	specNode()
}

//ValueSpec node represents a constant or variable declaration
type ValueSpec struct {
	Names *[]Ident
	Type   Expr
	Values []Expr


}

type BadDecl struct {
	From, To token.Pos
}

type GeneralDecl struct {
	TokPos  token.Pos
	Tok 	token.Token
	Specs   []Spec
}

type FunctionDecl struct {
	Name *Ident
	Type *FuncType

}

type FunctionType struct {
	Func   token.Pos
	Params *[]VarList
	Type   Expr
}







type File struct {
	Programs []*Program
	Scope
}

type AddOp struct {
	Op 		token.Token //Position of operator
	OpPos 	token.Pos
	List 	[]Expr
}

// body represents data reduced by production:
//
//	body:
//	        '{' decl stmts '}'  // kind 0



// expr represents data reduced by productions:
//
//	expr:
//	        ID '=' expr  // kind 0
//	|       expr1        // kind 1

type AssignExpr struct {
	LhsPos token.Pos
	ID 	   *Ident
}


// expr1 represents data reduced by productions:
//
//	expr1:
//	        term               // kind 0
//	|       expr1 add_op term  // kind 1




// expr_or_str represents data reduced by productions:
//
//	expr_or_str:
//	        expr                        // kind 0
//	|       STRING_LIT                  // kind 1
//	|       expr_or_str ',' expr        // kind 2
//	|       expr_or_str ',' STRING_LIT  // kind 3
type ExprOrString struct {

}


type Factor struct {
	FacPos token.Pos
	Kind   Kind
	IsNeg  bool
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


// function_def represents data reduced by production:
//
//	function_def:
//	        kind ID '(' kind ID ')' body  // kind 0
type FunctionDef struct {
	Define token.Pos
	Name   *Ident
	Type   *Ident
	Kind   Kind
	Body   Expr
}

// kind represents data reduced by productions:
//
//	kind:
//	        KW_INT    // kind 0
//	|       KW_FLOAT  // kind 1

type Kind token.Token //Can only be int or float


type MulOp struct {
	OpPos token.Pos
	Op 	   token.Token
	List   []*Factor
}

// program represents data reduced by productions:
//
//	program:
//	        /* empty */            // kind 0
//	|       program function_def   // kind 1
//	|       program decl           // kind 2
//	|       program function_decl  // kind 3

type Program struct {
	Define token.Pos
	Name   *Ident
	Type   *Ident
}






// write_expr_list represents data reduced by production:
//
//	write_expr_list:
//	        expr_or_str  // kind 0

func (a *AddOp) Pos() token.Pos 	{return a.OpPos}
func (m *MulOp) Pos() token.Pos 	{return m.OpPos}
func (f *Factor) Pos() token.Pos	{return f.FacPos}
func (t *Term) Pos() token.Pos 		{return t.TermPos}
func (i *Ident) Pos() token.Pos     {return i.NamePos}
func (i *IfStmt) Pos() token.Pos	{return i.IfKey}


func (a *AddOp) exprNode()  {}
func (m *MulOp) exprNode()  {}
func (f *Factor) exprNode() {}
func (t *Term) exprNode()   {}
func (i *Ident) exprNode()  {}

