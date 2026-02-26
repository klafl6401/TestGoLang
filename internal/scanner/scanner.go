package scanner

import (
	"github.com/klafl6401/TestGoLang/internal/token_type"
)

type Scanner struct {
	Pos    int
	Tokens []token_type.Token
	Line   int
	Source string
	Start  int
}

func (s *Scanner) AtEnd() bool {
	return s.Pos >= len(s.Source)
}

func (s *Scanner) advance() (ret string) {
	ret = string(s.Source[s.Pos])
	s.Pos++

	return
}

func (s *Scanner) AddTokenL(c string, kind int) {
	token := token_type.Token{
		Lexeme:  c,
		Literal: c,
		Kind:    kind,
		TokenDebug: token_type.TokenDebug{
			Line:  s.Line,
			Start: s.Start,
			End:   s.Pos,
		},
	}

	s.Tokens = append(s.Tokens, token)
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func (s *Scanner) ScanToken() {
	c := s.advance()
	switch c {
	case "+", "-", "*", "^", "(", ")", "[", "]", "{", "}", ",", ";", "#":
		// Single character simplistic tokens
		s.AddTokenL(c, token_type.StringToInt(c))
	}
}

func (s *Scanner) Scan() {
	for !s.AtEnd() {
		s.Start = s.Pos
		s.ScanToken()
	}
}
