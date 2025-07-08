# Demo HTTP Server

This application is a web server built with the Go standard library's `net/http` package. It serves a simple, modern login form using [htmx](https://htmx.org/) for dynamic interactions.

The server demonstrates several key concepts:

-   Routing with `net/http`.
-   Serving embedded static assets.
-   HTML template rendering.
-   Configuration management with environment variables.

---

## Getting Started

### Prerequisites

-   Go 1.22+

### Configuration

The server is configured through environment variables managed in the `.env` file in the repository root. Copy the example file if you haven't already:

```bash
cp .env.example .env
```

The primary configuration for this application is `PORT`.

### Running the Server

You can run the server using the project's `Makefile` from the repository root.

#### Standard Run

This command builds and runs the application.

```bash
# Run the server
make run APP=demo-http-server
```

The server will start and listen on the port specified in your `.env` file (default is `:3000`). You can access it at [http://localhost:3000](http://localhost:3000).

#### Live Reloading

For development, you can run the server with live reloading. The server will automatically restart when you make changes to the source code.

```bash
# Run the server with live reloading
make run/live APP=demo-http-server
```