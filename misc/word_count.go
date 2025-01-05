package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Define command-line flags
	wordPtr := flag.String("word", "", "The word to search for")
	caseInsensitivePtr := flag.Bool("case-insensitive", false, "Perform case-insensitive search")
	flag.Parse()

	// Check if a word is provided
	if *wordPtr == "" {
		fmt.Println("Usage: go run main.go -word <word> [-case-insensitive] <filename>")
		return
	}

	// Get the filename from the remaining arguments
	if flag.NArg() == 0 {
		fmt.Println("Missing filename")
		return
	}
	filename := flag.Arg(0)

	// Read the file content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Perform the search
	word := *wordPtr
	if *caseInsensitivePtr {
		word = strings.ToLower(word)
		content = []byte(strings.ToLower(string(content)))
	}

	count := strings.Count(string(content), word)

	// Print the result
	fmt.Printf("The word '%s' appears %d times in the file.\n", *wordPtr, count)
}
