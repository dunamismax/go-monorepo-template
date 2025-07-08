# Applications

This directory contains the main applications for this monorepo. Each subdirectory is a distinct, runnable Go application.

## Overview

Each application in this directory is a self-contained executable. This promotes a clear separation of concerns and allows for independent development, testing, and deployment.

---

## Applications

- **[demo-http-server](./demo-http-server)**: A web server built with the Go standard library that serves a simple login form using **htmx**. It demonstrates routing, serving embedded static assets and HTML template rendering.

- **[demo-cli-tool](./demo-cli-tool)**: A command-line interface (CLI) application built with the Go standard library's `flag` package. It demonstrates how to create a modern, user-friendly CLI with subcommands.

---

## Getting Started

To run any of these applications, use the `Makefile` from the root of the monorepo. Set the `APP` variable to the application's directory name.

### Examples

```bash
# Run the demo-http-server
make run APP=demo-http-server

# Run the demo-cli-tool
make run APP=demo-cli-tool -- hello
```

For more details on running projects, refer to the main `README.md`.
