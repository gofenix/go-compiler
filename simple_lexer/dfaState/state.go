package dfaState

type DfaState string

const (
	Initial    = "initial"
	If         = "if"
	Id_if1     = "id_if1"
	Id_if2     = "id_if2"
	Else       = "else"
	Id_else1   = "id_else1"
	Id_else2   = "id_else2"
	Id_else3   = "id_else3"
	Id_else4   = "id_else4"
	Int        = "int"
	Id_int1    = "id_int1"
	Id_int2    = "id_int2"
	Id_int3    = "id_int3"
	Id         = "id"
	GT         = "gt"
	GE         = "ge"
	Assignment = "assignment"

	Plus  = "plus"
	Minus = "minus"
	Star  = "star"
	Slash = "slash"

	SemiColon  = "semi_colon"
	LeftParen  = "left_paren"
	RightParen = "right_paren"

	IntLiteral = "int_literal"
)
