package scanner

import (
	"log"
	"strconv"

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

func (s *Scanner) peek() string {
	if s.AtEnd() {
		return "\000"
	}

	return string(s.Source[s.Pos])
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

func (s *Scanner) AddToken(lexeme string, literal any, kind int) {
	token := token_type.Token{
		Lexeme:  lexeme,
		Literal: literal,
		Kind:    kind,
		TokenDebug: token_type.TokenDebug{
			Line:  s.Line,
			Start: s.Start,
			End:   s.Pos,
		},
	}

	s.Tokens = append(s.Tokens, token)
}

func isAlpha(char string) bool {
	return (char >= "a" && char <= "z") || (char >= "A" && char <= "Z") || char == "_"
}

func isDigit(char string) bool {
	return (char >= "0" && char <= "9")
}

func isAlphaNumeric(char string) bool {
	return isAlpha(char) || isDigit(char)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == "." && !s.AtEnd() {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
		if s.peek() == "." {
			log.Fatalf("Unterminated float %s", s.Source[s.Start:s.Pos])
		}
	}
	lexeme := s.Source[s.Start:s.Pos]
	literal, err := strconv.ParseFloat(lexeme, 32)
	if err != nil {
		log.Fatalf("Problem tokenizing float %s %s", lexeme, err.Error())
	}

	s.AddToken(lexeme, literal, token_type.FLOAT)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) && !s.AtEnd() {
		s.advance()
	}
	lexeme := s.Source[s.Start:s.Pos]
	s.AddToken(lexeme, lexeme, token_type.IDENT)
}

func (s *Scanner) ScanToken() {
	c := s.advance()
	switch c {
	case "+", "-", "*", "^", "(", ")", "[", "]", "{", "}", ",", ";", "#":
		// Single character simplistic tokens
		s.AddTokenL(c, token_type.StringToInt(c))
	case " ", "\r", "\t":
	case "\n":
		s.Line++
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		}

	}
}

func (s *Scanner) Scan() {
	s.Line = 1
	for !s.AtEnd() {
		s.Start = s.Pos
		s.ScanToken()
	}
}
