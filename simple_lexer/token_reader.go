package simple_lexer

type TokenReader interface {
	Read() Token
	Peek() Token
	Unread()
	GetPosition() int
	SetPosition(position int)
}

type SimpleTokenReader struct {
	Tokens []Token
	Pos    int
}

func NewSimpleTokenReader(tokens []Token) *SimpleTokenReader {
	return &SimpleTokenReader{
		Tokens: tokens,
	}
}

func (s *SimpleTokenReader) Read() Token {
	if s.Pos < len(s.Tokens) {
		token := s.Tokens[s.Pos]
		s.Pos++
		return token
	}
	return nil
}

func (s *SimpleTokenReader) Peek() Token {
	if s.Pos < len(s.Tokens) {
		return s.Tokens[s.Pos]
	}
	return nil
}

func (s *SimpleTokenReader) Unread() {
	if s.Pos > 0 {
		s.Pos--
	}
}

func (s *SimpleTokenReader) GetPosition() int {
	return s.Pos
}

func (s *SimpleTokenReader) SetPosition(position int) {
	if position > 0 && position < len(s.Tokens) {
		s.Pos = position
	}
}
