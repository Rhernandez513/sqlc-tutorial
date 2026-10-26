# AI Agent Instructions

## Project Overview
This repository is a Go application demonstrating type-safe PostgreSQL database operations using `sqlc`. 

## Tech Stack
- **Language**: Go
- **Database**: PostgreSQL 15
- **Driver**: `github.com/jackc/pgx/v5`
- **Code Generator**: `sqlc`
- **Container Runtime**: Rancher Desktop (`nerdctl`)

## Local Environment
The application connects to a local PostgreSQL container. Do not assume a standard Docker daemon; use `nerdctl` for container management. 
- **Database**: `tutorial`
- **Port**: `5432`

If the database container needs to be restarted, use:
```bash
nerdctl start <container_id> # or postgres name

```

## Development Workflow & Rules

### 1. Code Generation (CRITICAL)

**Never manually edit the Go files inside the generated `tutorial/` package.**
All database interactions are defined in SQL.

* To update tables: Edit `schema.sql`
* To add/modify queries: Edit `query.sql`

Whenever `schema.sql` or `query.sql` are modified, you MUST run:

```bash
sqlc generate

```

*Note: If verifying against the sqlc cloud project, use `sqlc verify --against tutorial`.*

### 2. Dependencies

Ensure Go modules are tidy after generating new code that might import new packages:

```bash
go mod tidy

```

### 3. Execution

To execute the main application:

```bash
go run tutorial.go

```
