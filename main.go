package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/json-parser/lexer"
	"example.com/json-parser/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	isValid := doParsing(scanner)
	if isValid {
		fmt.Println("Valid JSON!")
	} else {
		fmt.Println("Invalid JSON!")
	}
}

func doParsing(scanner *bufio.Scanner) bool {
	tokens, isValidLexing := lexer.GetTokens(scanner)

	if !isValidLexing {
		return false
	}
	
	isValid := parser.ParseTokens(tokens)
	return isValid
}
