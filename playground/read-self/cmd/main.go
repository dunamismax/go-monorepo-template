package main

import (
	"fmt"
	"os"
)

// main reads its own source code from the file system and prints it to the console.
func main() {
	// Get the path of the currently executing file.
	// For a simple script, we can assume we are running from the file's directory.
	filePath := "playground/read-self/main.go" // The path to this source file from the root.

	// Attempt to read the content of the source file.
	content, err := os.ReadFile(filePath)
	if err != nil {
		// If an error occurs (e.g., file not found), print an error message
		// to standard error and exit with a non-zero status code.
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	// Print a header to the console.
	fmt.Println("--- Start of Source Code ---")
	fmt.Println("")

	// Print the content of the file as a string.
	fmt.Println(string(content))

	// Print a footer to the console.
	fmt.Println("")
	fmt.Println("--- End of Source Code ---")
}
