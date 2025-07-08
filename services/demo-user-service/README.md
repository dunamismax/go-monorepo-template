# Demo User Service

This application is a microservice for managing users, built with the Go standard library, GORM, and bcrypt. It provides a JSON API for creating users and handling JWT-based authentication.

## Features

-   **User Creation**: An endpoint to create new users.
-   **JWT Authentication**: An endpoint to log in and receive a JSON Web Token (JWT).
-   **Database Integration**: Connects to a PostgreSQL database using GORM.

---

## Getting Started

### Prerequisites

-   Go 1.22+
-   A running PostgreSQL instance.

### Set Up Environment

The service requires a database connection string. Copy the example `.env` file in the repository root and update it with your PostgreSQL credentials.

```bash
# Copy the environment file
cp .env.example .env
```

Edit the `.env` file:

```
DATABASE_URL="postgres://user:password@localhost:5432/database_name"
```

### Database Migrations

GORM will automatically migrate the database schema on application startup based on the `User` struct.

### Running the Service

You can run the service using the project's `Makefile` from the repository root.

#### Standard Run

This command builds and runs the application.

```bash
# Run the service
make run APP=demo-user-service
```

The service will start and listen on the port specified in your `.env` file (default is `:3001`).

#### Live Reloading

For development, you can run the service with live reloading.

```bash
# Run the service with live reloading
make run/live APP=demo-user-service
```

---

## API Endpoints

-   `POST /users`: Create a new user.
-   `POST /login`: Authenticate and receive a JWT.
