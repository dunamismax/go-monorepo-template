<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png" alt="The Go programming language logo." width="100"/>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-monorepo-template">
    <img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=Go+Monorepo+Template.;A+Solid+Foundation+for+Your+Next+Project.;Scalable.++Maintainable.++Efficient." alt="Typing SVG" />
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
<summary><h3>The Pragmatic Go Stack (Click to Expand)</h3></summary>

This stack is designed for building self-contained, high-performance, and concurrent web applications. The architecture is centered around the Go standard library, supplemented by a minimal set of highly-regarded libraries to enhance productivity and security. This approach yields a robust, minimal-dependency application that is simple to deploy and maintain.

---

#### **Core Application & CLI**

- **Language:** [**Go**](https://go.dev/doc/) (v1.22+)
  - A statically typed, compiled language that serves as the application's foundation, known for its performance, native concurrency, and single-binary deployments.
- **Web Router:** [**`net/http`**](https://pkg.go.dev/net/http/)
  - The standard library's production-grade HTTP server and multiplexer (`http.ServeMux`), used to route incoming requests to the appropriate handler functions.
- **CLI Framework:** [**`flag`**](https://pkg.go.dev/flag/)
  - The standard library package for parsing command-line flags, used to configure the application's behavior at startup.
- **Database ORM:** [**GORM**](https://gorm.io/docs/)
  - A full-featured Object-Relational Mapper for Golang that provides a developer-friendly API for database interactions, simplifying common CRUD operations, queries, and schema management.
- **Database Access:** [**`database/sql`**](https://pkg.go.dev/database/sql/)
  - The standard library's generic SQL interface. It provides the underlying foundation upon which GORM and the database driver operate, ensuring stability and standardization.
- **Database Driver (PostgreSQL):** [**`lib/pq`**](https://pkg.go.dev/github.com/lib/pq)
  - A widely-used and stable PostgreSQL driver for Go. It implements the standard `database/sql` interface, enabling the application to communicate with a PostgreSQL database.
- **Database Migrations:** [**`golang-migrate/migrate`**](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)
  - A dedicated tool that manages database schema changes using versioned SQL files, runnable as a CLI or a library for robust version control.

#### **Developer Experience & Tooling**

- **Package & Environment Management:** [**Go Modules & Toolchain**](https://go.dev/doc/tool/)
  - The native Go toolchain manages dependencies, builds, testing, and other development tasks, providing a unified and consistent experience.
- **Linter & Formatter:** [**`go fmt`**](https://pkg.go.dev/cmd/gofmt/) & [**`go vet`**](https://pkg.go.dev/cmd/vet/)
  - `go fmt` automatically formats code to the canonical Go style, and `go vet` is a static analyzer that reports suspicious code constructs to help find bugs.
- **Configuration:** [**Viper**](https://pkg.go.dev/github.com/spf13/viper)
  - A complete configuration solution handling various formats (JSON, TOML, YAML), environment variables, and remote config systems.
- **Live Reloading:** [**Air**](https://github.com/air-verse/air)
  - A live-reloading command-line utility for Go applications. Air monitors file changes in the project directory and automatically recompiles and restarts the application, streamlining the development feedback loop.

#### **Frontend & User Experience**

- **Client-Side Interactivity:** [**htmx**](https://htmx.org/docs/) (v2.0.0)
  - A compact JavaScript library that enables modern user experiences like AJAX requests and partial page updates directly within HTML attributes, eliminating the need for custom client-side JavaScript. The library is served as a static asset.
- **Templating:** [**`html/template`**](https://pkg.go.dev/html/template/)
  - The standard library's server-side HTML rendering engine. It provides fast, secure templating with context-aware escaping to automatically prevent Cross-Site Scripting (XSS) vulnerabilities.
- **Go/htmx Integration:** **Standard Handlers**
  - Integration is achieved using standard `http.HandlerFunc` implementations. These handlers process requests and write back either full HTML documents or partial template fragments to the `http.ResponseWriter`, seamlessly responding to htmx-driven interactions.
- **Forms & Validation:** **Manual Struct Population & Methods**
  - Form data is parsed from incoming requests using `r.ParseForm()`, and the values are used to manually populate data structs. Validation logic is implemented as explicit methods on these structs for clear and precise control.
- **Client-Side Validation:** [**HTML5 Validation**](https://developer.mozilla.org/en-US/docs/Learn/Forms/Form_validation#using_built-in_form_validation)
  - Built-in browser features provide instant client-side validation for a responsive user experience, acting as the first line of defense for data integrity.

#### **Authentication**

- **Core Authentication:** [**`golang.org/x/crypto/bcrypt`**](https://pkg.go.dev/golang.org/x/crypto/bcrypt) & [**`crypto`**](https://pkg.go.dev/crypto/) Packages
  - Password security is handled using the industry-standard `bcrypt` hashing algorithm, provided by the official Go crypto repository. For session management, JSON Web Tokens (JWTs) are constructed and verified using the standard library's `crypto/hmac` and `encoding/base64` packages.

#### **Deployment & Production**

- **Web Server / Reverse Proxy:** [**Caddy**](https://caddyserver.com/docs/) (v2)
  - A production-grade, open-source web server with automatic HTTPS. It serves as a reverse proxy, securely routing traffic to the compiled Go application binary.
- **Asset Management:** [**`embed`**](https://pkg.go.dev/embed/)
  - The standard library's `embed` package bundles static assets—including CSS, images, and the `htmx.js` library—directly into the Go binary at compile time. This creates a single, self-contained executable that is incredibly easy to deploy.

</details>

---

## Project Structure

This monorepo is organized into several directories, each with a specific purpose:

- **[`cmd`](./cmd)**: Contains the main applications (entry points) within this monorepo. Each subdirectory is a runnable application.
- **[`db`](./db)**: Contains database-related files, including migrations.
- **[`docs`](./docs)**: Houses supplementary documentation for the project.
- **[`playground`](./playground)**: A directory for experimental code, proofs-of-concept, and single-file Go applications.
- **[`services`](./services)**: Contains independent microservices, each with its own business logic.
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
