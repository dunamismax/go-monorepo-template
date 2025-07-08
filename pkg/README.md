# Public Packages

This directory is intended for Go packages that are safe to be used by external applications or libraries outside of this monorepo. These packages are designed to be reusable, stable, and well-documented, adhering to best practices for public APIs.

## Purpose

The `pkg/` directory is used for:

- **Reusable Libraries**: Code that provides general-purpose functionality and can be consumed by other Go projects.
- **Public APIs**: Packages that expose well-defined interfaces for external use.
- **Shared Components**: Components that are generic enough to be useful beyond the scope of this monorepo.

---

## Contents

- **`demopubliclibrary/`**: A demonstration of a public library, showcasing how to structure code for external consumption.
- **`demopubliclibrary2/`**: Another example of a public library.

---

## Usage

Public packages can be imported by any Go project, including those outside this monorepo, using their full module path.

### Example Import

```go
import (
    "github.com/dunamismax/go-monorepo-template/pkg/demopubliclibrary"
)
```
