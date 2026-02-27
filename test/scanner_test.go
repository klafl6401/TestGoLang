package scanner

import (
	"fmt"
	"testing"

	"github.com/klafl6401/TestGoLang/internal/scanner"
)

func TestFloat(t *testing.T) {
	newS := scanner.Scanner{
		Source: "99 99.1 100 123123123",
	}

	newS.Scan()

	tokens := newS.Tokens
	if tokens[0].Literal != float64(99) {
		t.Fatalf("Token 1 is %v of %T, wanted %v\n", tokens[0].Lexeme, tokens[0].Literal, 99)
	} else {
		t.Logf("Token 1 is lowk a float64 with value 99\n")
	}
}

func TestIdentifiers(t *testing.T) {
	newS := scanner.Scanner{
		Source: "testing one two three",
	}

	newS.Scan()

	tokens := newS.Tokens
	cToken := tokens[3]
	if cToken.Literal != "three" {
		t.Fatalf("Last Token is %v of %T, length: %d; wanting %v\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme), "three")
	}
	fmt.Printf("'%s' of %T length: %d\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme))
}

func TestString(t *testing.T) {
	newS := scanner.Scanner{
		Source: `"This is a string"`,
	}

	newS.Scan()

	tokens := newS.Tokens
	cToken := tokens[0]
	if cToken.Literal != "This is a string" {
		t.Fatalf("Last Token is %v of %T, length: %d; wanting %v\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme), "This is a string")
	}
	fmt.Printf("'%s' of %T length: %d\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme))
}

func TestKeywords(t *testing.T) {
	newS := scanner.Scanner{
		Source: `let a`,
	}

	newS.Scan()

	tokens := newS.Tokens
	cToken := tokens[0]
	if cToken.Literal != "let" {
		t.Fatalf("Last Token is %v of %T, length: %d; wanting %v\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme), "let")
	}
	fmt.Printf("'%s' of %T length: %d of kind %d\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme), cToken.Kind)
}

func TestOperators(t *testing.T) {
	newS := scanner.Scanner{
		Source: `=`,
	}

	newS.Scan()

	tokens := newS.Tokens
	cToken := tokens[0]
	if cToken.Literal != "=" {
		t.Fatalf("Last Token is %v of %T, length: %d; wanting %v\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme), "=")
	}
	fmt.Printf("'%s' of %T length: %d of kind %d\n", cToken.Literal, cToken.Literal, len(cToken.Lexeme), cToken.Kind)
}
