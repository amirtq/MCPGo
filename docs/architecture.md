# MCPGo Architecture

This document provides a high-level overview of the architecture used in the MCPGo project.

## 1. Core Principles

The architecture is based on a **Feature-Sliced Design**. This means that instead of being organized by technical layers (e.g., "controllers," "services," "repositories"), the code is grouped by feature. This approach enhances modularity, making it easier to develop, maintain, and scale individual parts of the application.

Each feature slice is self-contained and exposes its functionality through a well-defined public API.

## 2. Core Components

The project is primarily organized into three main directories within the `backend/` folder:

```
backend/
  ├── api/
  │   └── [feature]/
  │       ├── router.go
  │       └── handlers.go
  ├── apps/
  │   └── [feature]/
  │       └── app.go
  └── services/
      └── [service]/
          └── service.go
```

### 2.1. `apps`

- **Location:** `backend/apps/`
- **Responsibility:** This is the heart of the application. Each subdirectory in `apps` represents a distinct feature and contains its core business logic. It defines the primary application object for that feature.
- **Example:** `backend/apps/health/app.go` contains the `App` struct and the `CheckHealth` method, which implements the logic for the health check feature.

### 2.2. `api`

- **Location:** `backend/api/`
- **Responsibility:** This layer exposes the application's features to the outside world via an HTTP API. It adapts incoming requests into calls to the corresponding `app` in the `apps` directory.
- **Key Components:**
    - **Router:** Defines the HTTP routes for a feature and maps them to the appropriate handlers.
    - **Handlers:** Contain the logic for handling incoming HTTP requests, calling the `app` to execute business logic, and formatting the HTTP response.
- **Example:** `backend/api/health/router.go` creates a new router for the health feature, defines the `/health` endpoint, and uses the `health.App` to get the application's status.

### 2.3. `services`

- **Location:** `backend/services/`
- **Responsibility:** This directory contains shared, cross-cutting concerns that can be used by any feature. These are the "infrastructure" level components of the application.
- **Examples:**
    - **SSL:** A service for managing TLS certificates.
    - **Database:** A service for connecting to and querying a database (if one were present).
    - **Connectors:** Clients for communicating with external MCP servers.

## 3. Request Flow Example

Here’s how a typical HTTP request flows through the system:

1.  An HTTP request (e.g., `GET /health`) is received by the main router in `backend/main.go`.
2.  The main router delegates the request to the router for the `health` feature, which is defined in `backend/api/health/router.go`.
3.  The `health` router invokes the `healthHandler`.
4.  The `healthHandler` calls the `CheckHealth` method on the `health.App` instance (from `backend/apps/health/app.go`) to execute the business logic.
5.  The `health.App` returns the result to the handler.
6.  The handler formats the result as a JSON response and sends it back to the client.

This structure ensures that each feature is well-encapsulated, making the codebase easier to understand and maintain.
