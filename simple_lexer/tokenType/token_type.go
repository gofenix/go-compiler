package tokenType

type TokenType string

const (
	Plus  = "plus"
	Minus = "minus"
	Star  = "star"
	Slash = "slash"

	GE = "ge"
	GT = "gt"
	EQ = "eq"
	LE = "le"
	LT = "lt"

	SemiColon  = "semi_colon"
	LeftParen  = "left_paren"
	RightParen = "right_paren"

	Assignment = "assignment"

	If   = "if"
	Else = "else"

	Int = "int"

	Identifier = "identifier"

	IntLiteral    = "int_literal"
	StringLiteral = "string_literal"
)
