package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	oldUsername = "dunamismax"
	projectRoot = "."
)

var ignoredDirs = map[string]bool{
	".git":  true,
	".idea": true,
	"tmp":   true,
}

func main() {
	fmt.Print("Enter your GitHub username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newUsername := scanner.Text()

	if newUsername == "" {
		fmt.Println("Username cannot be empty. Aborting.")
		os.Exit(1)
	}

	fmt.Printf("Replacing all instances of '%s' with '%s'...\n", oldUsername, newUsername)

	err := filepath.WalkDir(projectRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if ignoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			// Ignore read errors for binary files, etc.
			return nil
		}

		// Perform replacement
		if strings.Contains(string(content), oldUsername) {
			fmt.Printf("Updating %s\n", path)
			newContent := strings.ReplaceAll(string(content), oldUsername, newUsername)
			err = os.WriteFile(path, []byte(newContent), d.Type().Perm())
			if err != nil {
				fmt.Printf("Error writing to %s: %v\n", path, err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Username replacement complete.")
}
