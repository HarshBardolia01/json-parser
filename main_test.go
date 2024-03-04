package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"testing"
	"unicode"
)

func TestDoParsing(t *testing.T) {
	fileName := flag.Arg(0)
	curFile, err := os.Open(fileName)

	if err != nil {
		t.Errorf("Error opening file: %s", err)
	}

	fmt.Printf("\nTesting %s: ", fileName)

	scanner := bufio.NewScanner(curFile)
	isValid := doParsing(scanner)

	size := len(fileName)
	temp := fileName[0 : size-5]

	for len(temp) > 0 && unicode.IsDigit(rune(temp[len(temp)-1])) {
		temp = temp[0 : len(temp)-1]
	}

	verdict := temp[len(temp)-4:]
	shouldBeValid := (verdict == "pass")

	if isValid != shouldBeValid {
		t.Errorf("Test Failed: %s", fileName)
	}
}
