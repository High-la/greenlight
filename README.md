# Greenlight API

A production-ready REST API built with Go for managing movie data, users, and authentication — designed with a strong focus on robustness, security, and real-world backend patterns.

This project is based on concepts from _Let's Go Further book_, extended and implemented as a complete backend system.

# Overview

Greenlight is a backend service that provides:

- Movie management (CRUD)
- User registration and activation
- Token-based authentication
- Permission-based authorization
- Filtering, sorting, and pagination
- Rate limiting and security controls
- Metrics and observability

The API is designed to reflect real production systems, not just basic CRUD operations.

# Architecture

The project follows a modular and layered design, separating responsibilities clearly:

```
cmd/api        → Application entry point (HTTP server, routing, middleware)
internal/data  → Business models + database logic
internal/validator → Input validation
internal/mailer → Email system
internal/jsonlog → Structured logging
migrations     → Database schema management
```

### Key Design Principles

- **Dependency Injection via** application struct
- **Separation of concerns** (handlers vs data layer)
- **Reusable helpers** (JSON responses, validation, errors)
- **Context-aware operations** (timeouts, cancellation)
- **Minimal external dependencies** (focus on stdlib + clarity)

# Tech Stack

- **Language:** Go
- **HTTP Router:** httprouter
- **Database:** PostgreSQL
- **Auth:** Token-based authentication
- **Email:** SMTP / background workers
- **Logging:** Structured logging (slog)
- **Metrics:** expvar

# Core Features

### Movies API

- Create, read, update, delete movies
- Optimistic concurrency control (versioning)
- Full-text search support
- Filtering, sorting, pagination

### User Management

- User registration
- Email-based activation
- Password updates
- Secure password hashing

### Authentication

- Token-based authentication
- Secure token generation & validation
- Stateless API design

### Authorization (RBAC)

- Permission-based access control
- Middleware for protected routes
- Fine-grained authorization checks

### Rate Limiting

- Global rate limiting
- IP-based rate limiting
- Protects API from abuse and overload

### Observability & Metrics

- `/debug/vars` endpoint using `expvar`
- Custom runtime metrics
- Request-level insights

### Email System

- Background email sending
- User activation workflow
- Token-based email verification

### Concurrency & Background Tasks

- Safe background job execution
- Graceful shutdown handling
- Proper goroutine lifecycle management

### Error Handling

- Consistent JSON error responses
- Centralized error helpers
- Validation error reporting

### API Endpoints

- Health
  `GET /v1/healthcheck`
- Movies

```
GET /v1/movies
POST /v1/movies
GET /v1/movies/:id
PATCH /v1/movies/:id
DELETE /v1/movies/:id
```

- Users

```
POST /v1/users
PUT /v1/users/activated
PUT /v1/users/password
```

- Tokens

```
POST /v1/tokens/authentication
POST /v1/tokens/password-reset
```

- Metrics

```
GET /debug/vars
```

# Key Engineering Concepts Demonstrated

This project is particularly strong because it implements:

**API Design**

- RESTful routing with proper HTTP semantics
- Versioned endpoints (/v1)

**Data Integrity**

- Optimistic locking using version fields
- SQL constraints and validation

**Performance**

- Efficient query design
- Pagination to limit load
- Connection pooling/

**Security**

- Token-based auth (no sessions)
- Input validation and sanitization
- Rate limiting

**Reliability**

- Graceful shutdown
- Context timeouts for DB queries
- Background job safety
