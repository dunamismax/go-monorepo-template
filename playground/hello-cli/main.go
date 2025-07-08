package main

import (
	"fmt"
	"os"
)

// main is the entry point for the hello-cli application.
// It prints a greeting message. If a command-line argument
// is provided, it includes it in the greeting.
func main() {
	// Default greeting message
	message := "Hello from the playground CLI!"

	// Check if a command-line argument is provided
	if len(os.Args) > 1 {
		// Use the provided argument in the message
		name := os.Args[1]
		message = fmt.Sprintf("Hello, %s! Welcome to the playground.", name)
	}

	// Print the final message to the console
	fmt.Println(message)
}
