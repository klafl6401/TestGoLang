package token_type

const (
	STRING = iota
	FLOAT
	IDENT

	PLUS
	MINUS
	MULT
	DIV
	UP_CARET
	HASHTAG
	COMMA
	SEMICOLON

	EQ    // =
	EQEQ  // ==
	NOTEQ // !=
	GT    // >
	LT    // <
	GTE   //>=
	LTE   // <=

	L_PAREN
	R_PAREN
	L_BRACK
	R_BRACK
	L_CBRACK
	R_CBRACK

	LOG
	LET
	FUNC

	NOT // not
	AND // and

	EOF
)

type TokenDebug struct {
	Line  int
	Start int
	End   int
}

type Token struct {
	Lexeme  string
	Literal any
	Kind    int
	TokenDebug
}

func StringToInt(c string) int {
	strToIntM := map[string]int{
		"+": PLUS,
		"-": MINUS,
		"*": MULT,
		"^": UP_CARET,
		"(": L_PAREN,
		")": R_PAREN,
		"[": L_BRACK,
		"]": R_BRACK,
		"{": L_CBRACK,
		"}": R_CBRACK,
		",": COMMA,
		";": SEMICOLON,
		"#": HASHTAG,
	}

	return strToIntM[c]
}
