# Go Playground

This directory is a dedicated space for experimental and single-file Go applications. It's a sandbox for testing new ideas, learning Go features, and creating simple scripts without the overhead of the main application structure.

---

## Applications

The playground includes the following examples:

- **hello-cli/**: A simple command-line application that prints a greeting. It demonstrates basic command-line argument handling.
- **simple-server/**: A minimal web server that listens on port `8081` and responds to requests. It showcases the `net/http` package.
- **read-self/**: A program that reads and prints its own source code to the console, demonstrating file I/O.

---

## Getting Started

You can run any application in the playground using either the project's `Makefile` from the repository root or standard `go run` commands.

### Using `make`

The `Makefile` provides a convenient way to run the applications. The `APP` variable specifies which application to run.

```bash
# Run the hello-cli application
make run APP=hello-cli

# Run hello-cli with an argument
make run APP=hello-cli ARGS="World"

# Run the simple web server
make run APP=simple-server

# Run the self-reading program
make run APP=read-self
```

### Using `go run`

You can also run the applications directly using the `go run` command from within the repository root.

```bash
# Run the hello-cli application
go run ./playground/hello-cli/cmd/main.go

# Run hello-cli with an argument
go run ./playground/hello-cli/cmd/main.go World

# Run the simple web server
go run ./playground/simple-server/cmd/main.go
# Server will be available at http://localhost:8081

# Run the self-reading program
go run ./playground/read-self/cmd/main.go
```
