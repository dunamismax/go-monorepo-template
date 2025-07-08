package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GitHub Username Changer")
	fmt.Println("This script will replace all instances of 'dunamismax' with your GitHub username.")
	fmt.Print("Enter your GitHub username: ")

	newUsername, _ := reader.ReadString('\n')
	newUsername = strings.TrimSpace(newUsername)

	if newUsername == "" {
		fmt.Println("Username cannot be empty.")
		return
	}

	fmt.Printf("Are you sure you want to replace 'dunamismax' with '%s'? (y/n): ", newUsername)

	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(strings.ToLower(confirmation))

	if confirmation != "y" {
		fmt.Println("Operation cancelled.")
		return
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip .git directory and the tool's own directory
		if info.IsDir() && (info.Name() == ".git" || info.Name() == "username-changer") {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			newContent := strings.Replace(string(content), "dunamismax", newUsername, -1)

			if string(content) != newContent {
				err = os.WriteFile(path, []byte(newContent), info.Mode())
				if err != nil {
					return err
				}
				fmt.Printf("Updated: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Println("Successfully replaced all instances of 'dunamismax'.")
	}
}
