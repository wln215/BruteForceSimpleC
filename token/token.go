package token

type Token int

const (
	tok_start Token = iota

	EOF
	ILLEGAL
	COMMENT

	keyword_start
	KW_INT
	KW_FLOAT
	KW_IF
	KW_ELSE
	KW_WHILE
	KW_RETURN
	KW_READ
	KW_WRITE
	keyword_end

	op_start
	OP_PLUS
	OP_MINUS
	OP_MULTIPLY
	OP_DIVIDE
	OP_ASSIGN

	OP_EQ
	OP_LT
	OP_LE
	OP_GT
	OP_GE
	op_end

	group_start
	LPAR
	RPAR
	LBRACE
	RBRACE
	SEMI
	COMMA
	group_end

	lit_start
	IDENT
	INTEGER
	FLOAT
	STRING
	lit_end

	tok_end
)

var tokStrings = map[Token]string {
	EOF 		:		 "EOF",
	ILLEGAL		:		 "Illegal",
	COMMENT 	: 		 "Comment",
	KW_INT		:        "int",
	KW_FLOAT	:        "float",
	KW_IF		:        "if",
	KW_ELSE    	:        "else",
	KW_WHILE   	:        "while",
	KW_RETURN  	:        "return",
	KW_READ    	:        "read",
	KW_WRITE   	:        "write",

	OP_PLUS	  	:		 "+",
	OP_MINUS	:  		 "-",
	OP_MULTIPLY	:		 "*",
	OP_DIVIDE	:		 "/",
	OP_ASSIGN 	:		 "=",
	OP_EQ		:		"==",
	OP_LT		:		"<",
	OP_LE		:		"<=",
	OP_GT		:		">",
	OP_GE		:		">=",

	LPAR    	:        "(",
	RPAR    	:        ")",
	LBRACE  	:        "{",
	RBRACE  	:        "}",
	SEMI    	:        ";",
	COMMA   	:        ",",

	IDENT 		: 		 "Identifier",
	INTEGER 	: 		 "Integer",
	FLOAT 		: 		 "Float",
	STRING		: 		 "String",

}

func (t Token) IsKeyword() bool {
	return t > keyword_start && t < keyword_end
}

func (t Token) IsOperator() bool {
	return  t > op_start && t < op_end
}

func (t Token) IsGrouping() bool {
	return t > group_start && t < group_end
}

func (t Token) IsLiteral() bool {
	return t > lit_start && t < lit_end
}

func Lookup(str string) Token {
	for t, s := range tokStrings {
		if s == str {
			return t
		}
	}
	return IDENT
}

func (t Token) String() string {
	return tokStrings[t]
}

func (t Token) Valid() bool {
	return t > tok_start && t < tok_end
}





