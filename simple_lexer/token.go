package simple_lexer

import "compilers/simple_lexer/tokenType"

type Token interface {
	GetType() tokenType.TokenType
	GetText() string
}

type SimpleToken struct {
	Type tokenType.TokenType
	Text string
}

func (s SimpleToken) GetType() tokenType.TokenType {
	return s.Type
}

func (s SimpleToken) GetText() string {
	return s.Text
}

func NewSimpleToken() *SimpleToken {
	return &SimpleToken{}
}
