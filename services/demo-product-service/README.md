# Demo Product Service

This application is a microservice responsible for managing product-related data and operations within the monorepo. It is built with the Go standard library and GORM, adhering to "The Pragmatic Go Stack" principles.

## Features

- **Product Management**: Handles creation, retrieval, updating, and deletion of product information.
- **Database Integration**: Connects to a PostgreSQL database using GORM.

---

## Getting Started

### Prerequisites

- Go 1.22+
- A running PostgreSQL instance.

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

GORM will automatically migrate the database schema on application startup based on the service's Go structs (models).

### Running the Service

You can run the service using the project's `Makefile` from the repository root.

#### Standard Run

This command builds and runs the application.

```bash
# Run the service
make run APP=demo-product-service
```

The service will start and listen on a default port (e.g., `:3002` or as configured internally).

#### Live Reloading

For development, you can run the service with live reloading.

```bash
# Run the service with live reloading
make run/live APP=demo-product-service
```

---

## API Endpoints

(Note: Specific API endpoints would be detailed here once implemented within the service's code.)
