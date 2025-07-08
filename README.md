<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png" alt="The Go programming language logo." width="100"/>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-monorepo-template">
    <img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=550&lines=Go+Monorepo+Template;A+Solid+Foundation+for+Your+Next+Project;Scalable.++Maintainable.++Efficient." alt="Typing SVG" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-monorepo-template/actions/workflows/makefile.yml"><img src="https://github.com/dunamismax/go-monorepo-template/actions/workflows/makefile.yml/badge.svg" alt="Build Status"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Language-Go-blue.svg" alt="Language: Go"></a>
  <a href="https://img.shields.io/github/repo-size/dunamismax/go-monorepo-template"><img src="https://img.shields.io/github/repo-size/dunamismax/go-monorepo-template" alt="Repo Size"></a>
  <a href="https://github.com/dunamismax/go-monorepo-template/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT"></a>
  <a href="https://github.com/dunamismax/go-monorepo-template/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
  <a href="https://github.com/dunamismax/go-monorepo-template/stargazers"><img src="https://img.shields.io/github/stars/dunamismax/go-monorepo-template" alt="GitHub Stars"></a>
</p>

---

## About This Project

This repository is a **Go Monorepo Template**, designed to provide a robust and scalable starting point for your Go projects. It includes a variety of applications and services, a unified build system using a `Makefile`, and a well-defined project structure.

The primary goal of this template is to offer a clean, maintainable, and scalable environment for building modern web applications, APIs, and command-line tools with Go.

### Target Environments

All code in this repository is designed to be developed and tested on **macOS** and deployed to **Linux (Ubuntu)**.

<details>
<summary><h3>Technology Stack (Click to Expand)</h3></summary>

This template utilizes a curated set of technologies to ensure a high-quality development experience and a performant end product.

---

This stack is for building self-contained, high-performance Go web applications. It uses the Go standard library with a few key libraries for a minimal-dependency, easy-to-deploy result.

### **Core Application & CLI**

- **Language:** [**Go**](https://go.dev/doc/) (v1.22+)
  - A statically typed, compiled language known for performance, concurrency, and single-binary deployments.
- **Web Router:** [**`net/http`**](https://pkg.go.dev/net/http/)
  - Go's standard library HTTP server and router for handling web requests.
- **CLI Framework:** [**`flag`**](https://pkg.go.dev/flag/)
  - Go's standard library package for parsing command-line flags.
- **Database ORM:** [**GORM**](https://gorm.io/docs/)
  - A developer-friendly ORM for database interactions like CRUD, queries, and schema management.
- **Database Access:** [**`database/sql`**](https://pkg.go.dev/database/sql/)
  - The standard library's generic SQL interface, providing a stable base for database drivers.
- **Database Driver (PostgreSQL):** [**`lib/pq`**](https://pkg.go.dev/github.com/lib/pq)
  - A popular and stable PostgreSQL driver for Go that works with `database/sql`.
- **Database Migrations:** [**`golang-migrate/migrate`**](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)
  - A dedicated tool that manages database schema changes using versioned SQL files, runnable as a CLI or a library for robust version control.

### **Developer Experience & Tooling**

- **Package & Environment Management:** [**Go Modules & Toolchain**](https://go.dev/doc/tool/)
  - Native Go tooling for managing dependencies, builds, and tests.
- **Linter & Formatter:** [**`go fmt`**](https://pkg.go.dev/cmd/gofmt/) & [**`go vet`**](https://pkg.go.dev/cmd/vet/)
  - `go fmt` ensures consistent code style; `go vet` finds potential bugs.
- **Configuration:** [**Viper**](https://pkg.go.dev/github.com/spf13/viper)
  - A complete configuration solution handling various formats (JSON, TOML, YAML), environment variables, and remote config systems.
- **Live Reloading:** [**Air**](https://github.com/air-verse/air)
  - A command-line tool that automatically rebuilds and restarts the app on file changes.

### **Frontend & User Experience**

- **Client-Side Interactivity:** [**htmx**](https://htmx.org/docs/) (v2.0.0)
  - A small JavaScript library enabling modern AJAX and partial page updates directly in HTML.
- **Templating:** [**`html/template`**](https://pkg.go.dev/html/template/)
  - Go's standard library for secure, server-side HTML rendering with automatic XSS protection.
- **Go/htmx Integration:** **Standard Handlers**
  - Standard Go HTTP handlers render full pages or partial HTML fragments for htmx.
- **Forms & Validation:** **Manual Struct Population & Methods**
  - Form data is manually parsed and validated using methods on Go structs.
- **Client-Side Validation:** [**HTML5 Validation**](https://developer.mozilla.org/en-US/docs/Learn/Forms/Form_validation#using_built-in_form_validation)
  - Uses built-in browser validation for instant feedback on user input.

### **Authentication**

- **Core Authentication:** [**`golang.org/x/crypto/bcrypt`**](https://pkg.go.dev/golang.org/x/crypto/bcrypt) & [**`crypto`**](https://pkg.go.dev/crypto/) Packages
  - Uses `bcrypt` for secure password hashing and standard crypto packages for session management (e.g., JWTs).

### **Deployment & Production**

- **Web Server / Reverse Proxy:** [**Caddy**](https://caddyserver.com/docs/) (v2)
  - A modern web server and reverse proxy with automatic HTTPS.
- **Asset Management:** [**`embed`**](https://pkg.go.dev/embed/)
  - Go's standard library package to bundle static assets (CSS, JS, images) into the application binary.

</details>

<details>
<summary><h3>Repository Structure (Click to Expand)</h3></summary>

```sh
/
├── api/
├── cmd/
│   ├── demo-cli-tool/
│   └── demo-http-server/
├── docs/
├── playground/
├── services/
│   ├── demo-product-service/
│   └── demo-user-service/
├── shared/
├── tools/
│   └── username-changer/
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└��─ README.md
```

</details>

---

## Project Overview

This monorepo is organized into several directories, each with a specific purpose:

- **[api](./api)**: Contains API definitions and data access logic.
- **[cmd](./cmd)**: Contains the main applications within this monorepo.
- **[docs](./docs)**: Contains supplementary documentation.
- **[playground](./playground)**: A directory for experimental and single-file Go applications.
- **[services](./services)**: Contains independent microservices.
- **[shared](./shared)**: Houses common utilities or types that are used across various components.
- **[tools](./tools)**: Stores Go tools and utilities used for development.

---

<details>
<summary><h3>Getting Started (Click to Expand)</h3></summary>

#### 1. Prerequisites

- Go 1.22+
- Docker (for running services with Docker Compose)
- A running PostgreSQL instance

#### 2. Clone the Repository

First, fork this repository. Then, clone your forked repository to your local machine.

```bash
git clone https://github.com/dunamismax/go-monorepo-template.git
cd go-monorepo-template
```

#### 3. Customize for Your Use

This template is configured with the username `dunamismax`. To make it your own, you need to replace this username with your own GitHub username. You can do this in two ways:

**Option 1: Use the Automated Script**

This repository includes a tool to automate the process. To run it, use the following command:

```bash
make change-username
```

The script will prompt you for your GitHub username and replace all instances of `dunamismax` with your input.

**Option 2: Manual Search and Replace**

If you prefer to make the changes manually, you can use the search and replace feature in your code editor (like VS Code).

1. Open the project in your editor.
2. Search for `dunamismax` across all files.
3. Replace all instances with your GitHub username.

#### 4. Set Up Environment

```bash
# Copy the example environment file
cp .env.example .env
```

Update your `.env` file with the correct connection string for your PostgreSQL database.

#### 5. Running a Project

To run a project, you can use the provided `Makefile` or `docker-compose`.

**Using Make**

The `APP` variable specifies which application to run.

```bash
# Run the demo-http-server (default)
make run

# Run the demo-user-service
make run APP=demo-user-service

# Run the demo-cli-tool with the "hello" command
make run APP=demo-cli-tool ARGS="hello --name World"
```

For live reloading during development:

```bash
# Live-reload the demo-http-server
make run/live APP=demo-http-server
```

**Using Docker Compose**

You can also run the `demo-http-server` using Docker Compose:

```bash
docker-compose up --build
```

To see a full list of available commands and applications, run:

```bash
make help
```

</details>

---

## Contributing

Contributions to this template are welcome! Please feel free to fork the repository, create a feature branch, and open a pull request.

---

## License

This repository is licensed under the **MIT License**. See the `LICENSE` file for more details.
