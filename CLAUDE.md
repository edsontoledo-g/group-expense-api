# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the server (defaults to :8080)
go run ./cmd/api/main.go

# Build the binary
go build -o bin/api ./cmd/api/main.go

# Run all tests
go test ./...

# Run tests for a specific package
go test ./internal/handlers/...

# Run a single test
go test ./internal/handlers/... -run TestFunctionName

# Tidy dependencies
go mod tidy
```

## Architecture

This is a Go REST API built with [Gin](https://github.com/gin-gonic/gin) and MongoDB (`go.mongodb.org/mongo-driver/v2`). The entry point is `cmd/api/main.go`.

All API routes are prefixed `/api/v1`. The `/api/v1` group is created in `main.go` and passed to each feature's `RegisterRoutes` function.

**Feature-based structure** — each domain lives in its own package under `internal/` with four files:

| File | Purpose |
|------|---------|
| `model.go` | Domain struct |
| `dto.go` | JSON request/response types |
| `handler.go` | HTTP handler struct + methods |
| `routes.go` | Exported `RegisterRoutes(*gin.RouterGroup)` |

Features: `internal/users/`, `internal/expenses/`, `internal/groups/`, `internal/invites/`

**Cross-feature dependency:** `groups/dto.go` imports `users.UserResponse` and `expenses.ExpenseResponse` because `GroupResponse` embeds them. The dependency graph is acyclic: `groups → users`, `groups → expenses`; neither `users` nor `expenses` imports `groups`.

**Current API surface:**

| Method | Path | Handler |
|--------|------|---------|
| GET | `/api/v1/me` | `users.UserHandler.GetInformation` |
| PATCH | `/api/v1/me/groups/:id/archive` | `users.UserHandler.ArchiveGroup` |
| GET | `/api/v1/groups` | `groups.GroupsHandler.GetGroups` |
| GET | `/api/v1/groups/:id` | `groups.GroupsHandler.GetGroup` |
| POST | `/api/v1/groups` | `groups.GroupsHandler.CreateGroup` |
| POST | `/api/v1/groups/:id/invite` | `groups.GroupsHandler.CreateInvite` |
| GET | `/api/v1/expenses` | `expenses.ExpensesHandler.GetExpensesHandler` |
| POST | `/api/v1/expenses` | `expenses.ExpensesHandler.CreateExpenseHandler` |
| DELETE | `/api/v1/expenses/:id` | `expenses.ExpensesHandler.DeleteExpenseHandler` |
| POST | `/api/v1/invites/:token` | `invites.InvitesHandler.AcceptInvite` |

**Status:** Most handlers are stubs (`// TODO: Implement`). The expenses handler has placeholder implementations using hardcoded data. No database layer or service layer exists yet — the next natural step is injecting a MongoDB client into the handler structs and adding a service/repository layer within each feature package.
