# Demo CLI Tool

This application is a command-line interface (CLI) tool built with the Go standard library's `flag` package. It demonstrates how to structure a simple CLI with subcommands and flags.

## Features

- A `hello` subcommand that prints a greeting.
- A `--name` flag to customize the greeting.

---

## Getting Started

### Prerequisites

- Go 1.22+

### Running the CLI

You can run the CLI tool using the project's `Makefile` from the repository root. The `APP` variable specifies that you want to run the `demo-cli-tool`. Any arguments after `--` are passed directly to the application.

```bash
# Run the 'hello' subcommand with the default name
make run APP=demo-cli-tool -- hello

# Run the 'hello' subcommand with a custom name
make run APP=demo-cli-tool -- hello --name "World"

# Display help information
make run APP=demo-cli-tool -- --help
```

Alternatively, you can run the application directly with `go run`:

```bash
# Run the 'hello' subcommand
go run ./cmd/demo-cli-tool/main.go hello --name "Universe"
```
