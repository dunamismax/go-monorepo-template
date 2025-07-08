# Internal Packages

This directory is reserved for private application and service-specific Go packages. Code within `internal/` cannot be imported by projects outside of this monorepo, ensuring strict encapsulation and preventing unintended external dependencies.

## Purpose

The `internal/` directory is used for:

- **Shared Logic**: Code that is shared between multiple applications or services within this monorepo but should not be exposed as a public library.
- **Domain-Specific Code**: Business logic, data access layers (DALs), or other components that are integral to the monorepo's applications but not intended for general reuse outside of it.
- **Avoiding Circular Dependencies**: Helps in structuring code to prevent circular imports between different parts of the monorepo.

---

## Contents

- **`demo-api/`**: Contains internal API definitions and data access logic for a demonstration API.
- **`sharedinternal/`**: Houses common utilities or types that are used across various internal components.

---

## Usage

Internal packages are used just like any other Go package, but their import scope is limited to this monorepo. This design pattern helps maintain a clear separation of concerns and prevents accidental exposure of internal implementation details.

### Example Import

```go
import (
    "github.com/dunamismax/go-monorepo-template/internal/demo-api/api"
    "github.com/dunamismax/go-monorepo-template/internal/sharedinternal"
)
```
