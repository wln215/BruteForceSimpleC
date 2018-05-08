package ast

type Scope struct {
	Parent *Scope
	Table  map[string]*Object
}


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

//Objects

type Object struct {
	Name 	string
	Kind 	ObKind
	Decl 	interface{}
	Data 	interface{}
	Type 	interface{}
}

type ObKind int

const (
	Bad ObKind = iota
	Var
	Con
	Function
)

var objKindStrings = [...]string{
	Bad: "bad",
	Var: "var",
	Con: "const",
	Function: "func",
}

func (kind ObKind) String() string {return objKindStrings[kind] }


