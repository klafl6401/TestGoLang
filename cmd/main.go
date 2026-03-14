package main

import (
	"fmt"

	"github.com/klafl6401/TestGoLang/internal/scanner"
)

func main() {
	newS := scanner.Scanner{
		Source: `georgie /* testing
		 */`,
	}

	newS.Scan()
	fmt.Println(newS.Tokens)
}
