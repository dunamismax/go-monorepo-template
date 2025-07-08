package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	helloCmd := flag.NewFlagSet("hello", flag.ExitOnError)
	name := helloCmd.String("name", "World", "The name to say hello to")

	if len(os.Args) < 2 {
		fmt.Println("expected 'hello' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hello":
		helloCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *name)
	default:
		fmt.Println("expected 'hello' subcommand")
		os.Exit(1)
	}
}
