# Awesome Go Projects

A curated list of legendary, exemplary, and perfect code GitHub projects written in Go. This list is designed to guide developers from beginner to pro, providing the best possible examples to learn from for building CLIs, services, and large-scale systems.

## Beginner Level: Grasping the Fundamentals

At this stage, the goal is to understand Go's syntax, core concepts, the standard library, and how to build simple, complete applications.

- **Gophercises by Jon Calhoun:** A fantastic, free course with exercises designed to take you from a beginner to a competent Go developer. Each exercise is a small, self-contained project.

  - **GitHub Link:** [https://github.com/gophercises/gophercises](https://github.com/gophercises/gophercises)
  - **What you'll learn:** Go syntax, standard library packages (`net/http`, `html/template`, `encoding/csv`), building a quiz game, a URL shortener, a file renamer, and even a simple photo transformation service. It's a hands-on-keyboard approach to learning.

- **A Simple Command-Line To-Do Manager:** Building a CLI tool is a rite of passage for Go developers. This project is a great example of a simple yet complete application.

  - **GitHub Link:** [https://github.com/spf13/cobra-cli](https://github.com/spf13/cobra-cli)
  - **What you'll learn:** While this is a tool to generate CLI apps, its own code and the projects it generates are a masterclass in structuring a command-line application. You'll learn how to define commands and sub-commands, manage flags, and create a professional-feeling CLI experience using the `cobra` library.

- **Building a Basic Web Server with the Standard Library:** Before jumping into frameworks, every Go developer should build a web application using only the `net/http` package. This tutorial provides the code and guidance to do so.

  - **Go.dev Tutorial Link:** [https://go.dev/doc/tutorial/web-service-gin](https://go.dev/doc/tutorial/web-service-gin) (While it uses Gin, the core concepts can be easily adapted to the standard library by following the `net/http` docs). For a pure `net/http` example, see the [Let's Go book by Alex Edwards](https://lets-go.alexedwards.net/).
  - **What you'll learn:** The fundamentals of the `net/http` package, how to create request handlers, how to serve JSON data, and the basic structure of a Go web service without external dependencies.

## Intermediate Level: Building Full-Fledged Applications

Intermediate projects introduce more complex topics like robust project structure, database integration, concurrency patterns, and in-depth testing.

- **Ardan Labs Service: A Production-Ready Service Template:** This project from Ardan Labs is the gold standard for how to structure a production-grade Go service. It's a highly opinionated but incredibly well-documented starting point.

  - **GitHub Link:** [https://github.com/ardanlabs/service](https://github.com/ardanlabs/service)
  - **What you'll learn:** Professional project layout, configuration management, structured logging, database integration and testing, authentication with JWTs, and how to build a robust, observable, and maintainable service.

- **GitHub CLI (`cli/cli`):** The official command-line tool for GitHub. It's a masterclass in building a complex, real-world CLI application that interacts with a third-party API.

  - **GitHub Link:** [https://github.com/cli/cli](https://github.com/cli/cli)
  - **What you'll learn:** Advanced CLI design, managing API clients and authentication, testing complex commands, user experience design for a CLI, and structuring a large Go application.

- **Go RealWorld Example App (Standard Library):** This project implements the "RealWorld" specification (a Medium.com clone) using only the Go standard library. It's a fantastic, real-world benchmark for understanding how a complete API is built from the ground up.

  - **GitHub Link:** [https://github.com/Clivern/Go-RealWorld-Example-App](https://github.com/Clivern/Go-RealWorld-Example-App)
  - **What you'll learn:** Building a full REST API without frameworks, JWT-based authentication, separating application logic into layers (handlers, models, etc.), and integrating with a database (like PostgreSQL) using standard library-compatible drivers.

## Advanced & Professional Level: Mastering Complexity and Scale

These projects showcase advanced patterns for building high-performance, scalable, and resilient systems suitable for production environments at the highest level.

- **Caddy Web Server:** A powerful, enterprise-ready, open-source web server with automatic HTTPS, written in Go. Its architecture is modern, modular, and incredibly well-designed.

  - **GitHub Link:** [https://github.com/caddyserver/caddy](https://github.com/caddyserver/caddy)
  - **What you'll learn:** Building a highly extensible plugin system, advanced concurrency patterns, managing complex configurations, handling low-level networking, and the architecture of a long-running, high-performance server.

- **Prometheus:** The legendary open-source monitoring and alerting toolkit. It's a cornerstone of the cloud-native ecosystem and an exemplary case study of a large-scale, distributed Go application.

  - **GitHub Link:** [https://github.com/prometheus/prometheus](https://github.com/prometheus/prometheus)
  - **What you'll learn:** Designing a time-series database (TSDB), building efficient data scraping mechanisms, advanced API design, creating a powerful query language (`PromQL`), and the architecture of a highly available, distributed system.

- **CockroachDB:** A cloud-native, distributed SQL database built in Go. It's a masterpiece of engineering that solves some of the hardest problems in distributed systems.

  - **GitHub Link:** [https://github.com/cockroachdb/cockroach](https://github.com/cockroachdb/cockroach)
  - **What you'll learn:** Implementing distributed consensus algorithms (like Raft), designing a distributed key-value store, building a SQL layer from scratch, advanced testing strategies for distributed systems, and a deep dive into high-performance Go for data-intensive applications.
