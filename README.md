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
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8.svg" alt="Go Version"></a>
  <img src="https://img.shields.io/github/repo-size/dunamismax/go-monorepo-template" alt="Repo Size">
  <a href="https://github.com/dunamismax/go-monorepo-template/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT"></a>
  <a href="https://github.com/dunamismax/go-monorepo-template/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome"></a>
  <a href="https://github.com/dunamismax/go-monorepo-template/stargazers"><img src="https://img.shields.io/github/stars/dunamismax/go-monorepo-template" alt="GitHub Stars"></a>
</p>

---

## About This Project

This repository is a **Go Monorepo Template** designed to provide a robust and scalable starting point for your projects. It includes a well-defined project structure, a unified build system using a `Makefile`, and a curated tech stack perfect for modern web applications, APIs, and command-line tools.

The primary goal is to offer a clean, maintainable, and efficient environment that lets you focus on building features, not boilerplate.

### Target Environments

All code is designed to be developed on **macOS** and deployed to **Linux (Ubuntu)**.

<details>
<summary><h3>Technology Stack (Click to Expand)</h3></summary>

The technology stack for this template is carefully curated to build high-performance, self-contained Go web applications. The philosophy is to lean heavily on Go's robust standard library, supplementing it only where necessary with a minimal set of highly-regarded libraries. This approach ensures a high-quality development experience and results in a performant, minimal-dependency application that is simple to deploy and maintain.

---

#### **Core Application & CLI**

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

#### **Developer Experience & Tooling**

- **Package & Environment Management:** [**Go Modules & Toolchain**](https://go.dev/doc/tool/)
  - Native Go tooling for managing dependencies, builds, and tests.
- **Linter & Formatter:** [**`go fmt`**](https://pkg.go.dev/cmd/gofmt/) & [**`go vet`**](https://pkg.go.dev/cmd/vet/)
  - `go fmt` ensures consistent code style; `go vet` finds potential bugs.
- **Configuration:** [**Viper**](https://pkg.go.dev/github.com/spf13/viper)
  - A complete configuration solution handling various formats (JSON, TOML, YAML), environment variables, and remote config systems.
- **Live Reloading:** [**Air**](https://github.com/air-verse/air)
  - A command-line tool that automatically rebuilds and restarts the app on file changes.

#### **Frontend & User Experience**

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

#### **Authentication**

- **Core Authentication:** [**`golang.org/x/crypto/bcrypt`**](https://pkg.go.dev/golang.org/x/crypto/bcrypt) & [**`crypto`**](https://pkg.go.dev/crypto/) Packages
  - Uses `bcrypt` for secure password hashing and standard crypto packages for session management (e.g., JWTs).

#### **Deployment & Production**

- **Web Server / Reverse Proxy:** [**Caddy**](https://caddyserver.com/docs/) (v2)
  - A modern web server and reverse proxy with automatic HTTPS.
- **Asset Management:** [**`embed`**](https://pkg.go.dev/embed/)
  - Go's standard library package to bundle static assets (CSS, JS, images) into the application binary.

</details>

---

## Project Structure

This monorepo is organized into several directories, each with a specific purpose:

- **[`api`](./api)**: Contains API definitions (e.g., OpenAPI/Swagger specs) and generated client/server code.
- **[`cmd`](./cmd)**: Contains the main applications (entry points) within this monorepo. Each subdirectory is a runnable application.
- **[`docs`](./docs)**: Houses supplementary documentation for the project.
- **[`playground`](./playground)**: A directory for experimental code, proofs-of-concept, and single-file Go applications.
- **[`services`](./services)**: Contains independent microservices, each with its own business logic.
- **[`shared`](./shared)**: A place for common utilities, types, or libraries that are used across multiple applications or services.
- **[`tools`](./tools)**: Stores Go-based scripts and utilities used for development and CI/CD tasks.

<details>

<summary><h3>Getting Started (Click to Expand)</h3></summary>

#### 1. Prerequisites

- **Go 1.22+**
- **Docker** & **Docker Compose**
- **A running PostgreSQL instance** (can be easily started with Docker)

#### 2. Clone the Repository

First, **fork** this repository to your own GitHub account. Then, clone your fork to your local machine, replacing `<Your-GitHub-Username>` with your actual username.

```bash
git clone https://github.com/<Your-GitHub-Username>/go-monorepo-template.git
cd go-monorepo-template
```

#### 3. Customize for Your Use

This template is configured with the module path `github.com/dunamismax/go-monorepo-template`. To make it your own, you must replace this path with your own.

**Option 1: Use the Automated Script (Recommended)**

This repository includes a tool to automate the process. It will prompt you for your GitHub username and replace all instances of `dunamismax` accordingly.

```bash
make change-username
```

**Option 2: Manual Search and Replace**

Use the search and replace feature in your code editor (like VS Code) to find all instances of `dunamismax` and replace them with your GitHub username.

#### 4. Set Up Environment

Copy the example environment file and update it with your database connection string.

```bash
# Copy the example environment file
cp .env.example .env
```

Now, open `.env` and configure your `DATABASE_URL`.

#### 5. Running a Project

You can run applications using the provided `Makefile` or `Docker Compose`. For a full list of commands, run `make help`.

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

You can also run the `demo-http-server` and a PostgreSQL database using Docker Compose:

```bash
docker-compose up --build
```

</details>

---

## Contributing

Contributions are welcome! Please feel free to fork the repository, create a feature branch, and open a pull request.

## License

This project is licensed under the **MIT License**. See the `LICENSE` file for more details.
