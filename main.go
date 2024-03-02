package main

import (
	"fmt"
	"os"

	filereader "example.com/json-parser/fileReader"
	"example.com/json-parser/lexer"
)

func main() {
	fileName := "test.json"
	inputJSON, err := filereader.ReadContent(fileName)

	if err != nil {
		fmt.Printf("Error reading the file\n%s\n", err)
		os.Exit(1)
	}

	lexer.GetTokens(inputJSON)
}
