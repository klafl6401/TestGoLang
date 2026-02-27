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

func (s *Scanner) match(expected string) bool {
	if s.AtEnd() || string(s.Source[s.Pos]) != expected {
		return false
	}
	s.Pos++

	return true
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

	switch lexeme {
	case "let":
		s.AddToken(lexeme, lexeme, token_type.LET)
	case "log":
		s.AddToken(lexeme, lexeme, token_type.LOG)
	case "func":
		s.AddToken(lexeme, lexeme, token_type.FUNC)
	case "not":
		s.AddToken(lexeme, lexeme, token_type.NOT)
	case "and":
		s.AddToken(lexeme, lexeme, token_type.AND)
	default:
		s.AddToken(lexeme, lexeme, token_type.IDENT)
	}
}

func (s *Scanner) string() {
	for s.peek() != "\"" && !s.AtEnd() {
		s.advance()
	}
	if s.peek() != "\"" {
		log.Fatalf("Unterminated string %s", s.Source[s.Start:s.Pos])
	} else if !s.AtEnd() {
		s.advance()
	}
	lexeme := s.Source[s.Start:s.Pos]
	literal := s.Source[s.Start+1 : s.Pos-1]
	s.AddToken(lexeme, literal, token_type.STRING)
}

func (s *Scanner) ScanToken() {
	c := s.advance()
	switch c {
	case "+", "-", "*", "^", "(", ")", "[", "]", "{", "}", ",", ";", "#":
		// Single character simplistic tokens
		s.AddTokenL(c, token_type.StringToInt(c))
	case "=":
		if s.match("=") {
			s.AddToken("==", "==", token_type.EQEQ)
		} else {
			s.AddTokenL("=", token_type.EQ)
		}
	case ">":
		if s.match("=") {
			s.AddToken(">=", ">=", token_type.GTE)
		} else {
			s.AddTokenL(">", token_type.GT)
		}
	case "<":
		if s.match("=") {
			s.AddToken("<=", "<=", token_type.LTE)
		} else {
			s.AddTokenL("<", token_type.LT)
		}
	case "!":
		if s.match("=") {
			s.AddToken("!=", "!=", token_type.NOTEQ)
		} else {
			s.AddTokenL("!", token_type.NOT)
		}
	case " ", "\r", "\t":
	case "\n":
		s.Line++
	case "\"":
		s.string()
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
	s.AddToken("\000", "\000", token_type.EOF)
}
