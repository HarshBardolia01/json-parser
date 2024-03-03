package main

import (
	"bufio"
	"os"

	"example.com/json-parser/lexer"
	"example.com/json-parser/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tokens := lexer.GetTokens(scanner)
	parser.ParseTokens(tokens)
}
