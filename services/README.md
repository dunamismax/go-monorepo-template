# Services

This directory contains independent microservices that are part of the larger monorepo ecosystem. Each subdirectory represents a distinct service with its own responsibilities, APIs, and potentially its own set of dependencies.

## Purpose

The `services/` directory is used for:

- **Microservices**: Self-contained, deployable units that perform specific business functions.
- **API Endpoints**: Services that expose APIs (e.g., REST, gRPC) for communication with other services or frontend applications.
- **Scalability**: Designed to be independently scalable and deployable.

---

## Services

- **[demo-product-service](./demo-product-service/)**: A demonstration microservice responsible for managing product-related data and operations.
- **[demo-user-service](./demo-user-service/)**: A microservice dedicated to user management, including authentication and user profile handling.

---

## Getting Started

Services are typically run as separate processes. They communicate with each other via their defined APIs. Refer to each service's individual `README.md` for specific instructions on how to run, configure, and interact with it.

### Example `Makefile` usage

```bash
# Run a specific service
make run APP=demo-user-service

# Build a specific service for production
make production/deploy APP=demo-product-service
```
