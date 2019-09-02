package simple_lexer

import (
	"bytes"
	"compilers/simple_lexer/dfaState"
	"compilers/simple_lexer/tokenType"
	"fmt"
)

type SimpleLexer struct {
	TokenText string
	Tokens    []Token
	Token     SimpleToken
}

func NewSimpleLexer() *SimpleLexer {
	return &SimpleLexer{}
}

func (s *SimpleLexer) Tokenize(code string) *SimpleTokenReader {
	var state dfaState.DfaState = dfaState.DfaState(dfaState.Initial)
	var err error

	reader := bytes.NewReader([]byte(code))
	var ch byte
	for {
		ch, err = reader.ReadByte()
		if err != nil {
			break
		}

		switch state {
		case dfaState.Initial:
			state = s.initToken(ch)
		case dfaState.Id:
			if isAlpha(ch) || isDigit(ch) {
				s.TokenText += string(ch)
			} else {
				state = s.initToken(ch)
			}
		case dfaState.GT:
			if ch == '=' {
				s.Token.Type = tokenType.GE
				state = dfaState.GE
				s.TokenText += string(ch)
			} else {
				state = s.initToken(ch)
			}
		case dfaState.GE, dfaState.Assignment, dfaState.Plus, dfaState.Minus, dfaState.Star, dfaState.Slash, dfaState.SemiColon, dfaState.LeftParen, dfaState.RightParen:
			state = s.initToken(ch)
		case dfaState.IntLiteral:
			if isDigit(ch) {
				s.TokenText += string(ch)
			} else {
				state = s.initToken(ch)
			}
		case dfaState.Id_int1:
			if ch == 'n' {
				state = dfaState.Id_int2
				s.TokenText += string(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				state = dfaState.Id
				s.TokenText += string(ch)
			} else {
				state = s.initToken(ch)
			}

		case dfaState.Id_int2:
			if ch == 't' {
				state = dfaState.Id_int3
				s.TokenText += string(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				state = dfaState.Id
				s.TokenText += string(ch)
			} else {
				state = s.initToken(ch)
			}

		case dfaState.Id_int3:
			if isBlank(ch) {
				s.Token.Type = tokenType.Int
				state = s.initToken(ch)
			} else {
				state = dfaState.Id
				s.TokenText += string(ch)
			}
		}
	}

	if len(s.TokenText) > 0 {
		s.initToken(ch)
	}

	return &SimpleTokenReader{Tokens: s.Tokens}
}

func (s *SimpleLexer) initToken(ch byte) dfaState.DfaState {
	if len(s.TokenText) > 0 {
		s.Token.Text = string(s.TokenText)
		s.Tokens = append(s.Tokens, s.Token)
		s.TokenText = ""
		s.Token = *NewSimpleToken()
	}

	newState := dfaState.Initial

	if isAlpha(ch) {
		if ch == 'i' {
			newState = dfaState.Id_int1
		} else {
			newState = dfaState.Id
		}
		s.Token.Type = tokenType.Identifier
		s.TokenText += string(ch)
	} else if isDigit(ch) {
		newState = dfaState.IntLiteral
		s.Token.Type = tokenType.IntLiteral
		s.TokenText += string(ch)
	} else if ch == '>' {
		newState = dfaState.GT
		s.Token.Type = tokenType.GT
		s.TokenText += string(ch)
	} else if ch == '+' {
		newState = dfaState.Plus
		s.Token.Type = tokenType.Plus
		s.TokenText += string(ch)
	} else if ch == '-' {
		newState = dfaState.Minus
		s.Token.Type = tokenType.Minus
		s.TokenText += string(ch)
	} else if ch == '*' {
		newState = dfaState.Star
		s.Token.Type = tokenType.Star
		s.TokenText += string(ch)
	} else if ch == '/' {
		newState = dfaState.Slash
		s.Token.Type = tokenType.Slash
		s.TokenText += string(ch)
	} else if ch == ';' {
		newState = dfaState.SemiColon
		s.Token.Type = tokenType.SemiColon
		s.TokenText += string(ch)
	} else if ch == '(' {
		newState = dfaState.LeftParen
		s.Token.Type = tokenType.LeftParen
		s.TokenText += string(ch)
	} else if ch == ')' {
		newState = dfaState.RightParen
		s.Token.Type = tokenType.RightParen
		s.TokenText += string(ch)
	} else if ch == '=' {
		newState = dfaState.Assignment
		s.Token.Type = tokenType.Assignment
		s.TokenText += string(ch)
	} else {
		newState = dfaState.Initial
	}
	return dfaState.DfaState(newState)
}

func Dump(tokenReader *SimpleTokenReader) {
	fmt.Println("text\t\ttype")
	for {
		token := tokenReader.Read()
		if token == nil {
			break
		}
		fmt.Println(token.GetText(), "\t\t", token.GetType())
	}
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isBlank(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
