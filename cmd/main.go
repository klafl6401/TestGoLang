package main

import "github.com/klafl6401/TestGoLang/internal/scanner"

func main() {
	newS := scanner.Scanner{
		Source: "Asdf 123.12",
	}

	newS.Scan()
}
