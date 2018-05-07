package ast

import (
	"BruteForceSimpleC/token"
	"reflect"
)

type Node interface{
	Pos() token.Pos
	End() token.Pos
}

type SubTree interface { //Theoretically cannot have a root
	Node
	subTreeNode()
}
type Ident struct {
	NamePos token.Pos
	Name 	string
	Object  *Object //nil if keyword
}

type Object struct {
	NamePos token.Pos
	Name 	string
	Kind 	ObKind
	Offset  int
	Type 	*Ident //var type
}

type ObKind int

const (
	Decl ObKind = iota
	Int
	Float
)

type Param struct {
	Name 	*Ident
	Type 	*Ident
}

type Scope struct {
	Parent *Scope
	Table  map[string]*Object
}

type File struct {
	Programs []*Program
	Scope
}

type BasicLiteral struct {
	LitPos token.Pos
	Kind   token.Token
	Lit    string
}


// add_op represents data reduced by productions:
//
//	add_op:
//	        '+'  // kind 0
//	|       '-'  // kind 1
type AddOp struct {
	Op 		token.Token //Position of operator
	OpPos 	token.Pos
	List 	[]SubTree
}

// body represents data reduced by production:
//
//	body:
//	        '{' decl stmts '}'  // kind 0


// bool_expr represents data reduced by production:
//
//	bool_expr:
//	        expr bool_op expr  // kind 0


// bool_op represents data reduced by productions:
//
//	bool_op:
//	        '<'    // kind 0
//	|       '>'    // kind 1
//	|       OP_EQ  // kind 2
//	|       OP_LE  // kind 3
//	|       OP_GE  // kind 4


// decl represents data reduced by production:
//
//	decl:
//	        kind var_list ';'  // kind 0


// expr represents data reduced by productions:
//
//	expr:
//	        ID '=' expr  // kind 0
//	|       expr1        // kind 1


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


// factor represents data reduced by productions:
//
//	factor:
//	        ID             // kind 0
//	|       INT_LIT        // kind 1
//	|       FLOAT_LIT      // kind 2
//	|       function_call  // kind 3
//	|       '(' expr ')'   // kind 4
type Factor struct {
	Pos    token.Pos
	Kind   Kind
	IsNeg  bool
}

// function_call represents data reduced by production:
//
//	function_call:
//	        ID '(' expr ')'  // kind 0
type FunctionCall struct {
	Name *Ident
	Args []SubTree
}

// function_decl represents data reduced by production:
//
//	function_decl:
//	        kind ID '(' kind ')' ';'  // kind 0
type FunctionDecl struct {
	Define token.Pos
	Name   *Ident
	Type   *Ident
	Kind   Kind
}

// function_def represents data reduced by production:
//
//	function_def:
//	        kind ID '(' kind ID ')' body  // kind 0
type FunctionDef struct {
	Define token.Pos
	Name   *Ident
	Type   *Ident
	Kind   Kind
	Body   SubTree
}

// kind represents data reduced by productions:
//
//	kind:
//	        KW_INT    // kind 0
//	|       KW_FLOAT  // kind 1


// mul_op represents data reduced by productions:
//
//	mul_op:
//	        '*'  // kind 0
//	|       '/'  // kind 1


// program represents data reduced by productions:
//
//	program:
//	        /* empty */            // kind 0
//	|       program function_def   // kind 1
//	|       program decl           // kind 2
//	|       program function_decl  // kind 3
type Programs struct {
	Programs []*Program
}

type Program struct {
	Define token.Pos
	Name   *Ident
	Type   *Ident
}

// start represents data reduced by production:
//
//	start:
//	        program  // kind 0


// stmt represents data reduced by productions:
//
//	stmt:
//	        expr ';'                                   // kind 0
//	|       KW_IF '(' bool_expr ')' stmt               // kind 1
//	|       KW_IF '(' bool_expr ')' stmt KW_ELSE stmt  // kind 2
//	|       KW_WHILE '(' bool_expr ')' stmt            // kind 3
//	|       KW_READ var_list ';'                       // kind 4
//	|       KW_WRITE write_expr_list ';'               // kind 5
//	|       KW_RETURN expr ';'                         // kind 6
//	|       '{' stmts '}'                              // kind 7


// stmts represents data reduced by productions:
//
//	stmts:
//	        /* empty */  // kind 0
//	|       stmts stmt   // kind 1


// term represents data reduced by productions:
//
//	term:
//	        factor                  // kind 0
//	|       '-' factor              // kind 1
//	|       term mul_op factor      // kind 2
//	|       term mul_op '-' factor  // kind 3
type Term struct {

}


// var_list represents data reduced by productions:
//
//	var_list:
//	        ID               // kind 0
//	|       var_list ',' ID  // kind 1
type  VarList struct {
	p.listok = false
}

// write_expr_list represents data reduced by production:
//
//	write_expr_list:
//	        expr_or_str  // kind 0


func NewScope(parent *Scope) *Scope {
	return &Scope{Parent:parent, Table: make(map[string]*Object)}
}

func (s *Scope) Insert(ob *Object) *Object {
	if old, ok := s.Table[ob.Name]; ok {
		return old
	}
	s.Table[ob.Name] = ob
	return nil
}

func (s *Scope) Lookup(ident string) *Object {
	ob, ok := s.Table[ident]
	if ok || s.Parent == nil {
		return ob
	}
	return s.Parent.Lookup(ident)
}