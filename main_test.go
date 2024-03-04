package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestDoParsing(t *testing.T) {
	// fileNameFromArgs := flag.Arg(0)
	// fmt.Println(fileNameFromArgs)

	files, err := os.Open("./test")

	if err != nil {
		t.Errorf("Error opening Directory: %s", err)
	}

	defer files.Close()
	fileInfo, err := files.ReadDir(-1)

	if err != nil {
		t.Errorf("Error reading Directory: %s", err)
	}

	var failedTests []string

	for _, file := range fileInfo {
		fileName := file.Name()
		curFile, err := os.Open(fmt.Sprintf("./test/%s", fileName))

		if err != nil {
			t.Errorf("Error opening file: %s", err)
		}
		scanner := bufio.NewScanner(curFile)
		isValid := doParsing(scanner)
		shouldBeValid := (fileName[0] == 'p')

		if isValid != shouldBeValid {
			failedTests = append(failedTests, fileName)
		}
	}

	if len(failedTests) > 0 {
		fmt.Println("Following test cases were Failed")
		for _, fileName := range failedTests {
			fmt.Println(fileName)
		}
	}
}
