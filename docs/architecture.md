# MCPGo Architecture

This document provides a high-level overview of the architecture used in the MCPGo project.

## 1. Core Principles

The architecture is based on a combination of well-established patterns:

- **Clean Architecture:** Enforces a strict separation of concerns, ensuring the core business logic (Domain) is independent of external frameworks and technologies.
- **CQRS (Command Query Responsibility Segregation):** Separates read and write operations, allowing for optimized and scalable data models for each.
- **Hexagonal Architecture (Ports and Adapters):** The application communicates with the outside world through "ports" (interfaces), which are implemented by "adapters" (infrastructure).

## 2. The Layers

The project is divided into four primary layers, with a strict dependency rule: outer layers can depend on inner layers, but not the other way around.

```
+-----------------------------------------------------------------+
|                           Interfaces                            |
|              (HTTP Handlers, WebSocket Controllers)             |
+-----------------------------------------------------------------+
      |
      v
+-----------------------------------------------------------------+
|                          Application                            |
| (Commands, Queries, Handlers, DTOs, Ports/Interfaces)           |
+-----------------------------------------------------------------+
      |
      v
+-----------------------------------------------------------------+
|                             Domain                              |
|              (Entities, Value Objects, Domain Events)           |
+-----------------------------------------------------------------+

+-----------------------------------------------------------------+
|                         Infrastructure                          |
| (DB Repos, Event Bus Impl, Config Loader, Logger, Connectors)   |
+-----------------------------------------------------------------+
```

### 2.1. Domain Layer

- **Location:** `internal/domain/`
- **Responsibility:** This is the heart of the application. It contains the core business logic, rules, and data structures (Entities, Value Objects).
- **Rules:**
    - It is pure and has **zero dependencies** on any other layer or external library.
    - It knows nothing about databases, APIs, or frameworks.

### 2.2. Application Layer

- **Location:** `internal/application/`
- **Responsibility:** This layer orchestrates the use cases of the application. It defines the commands and queries that represent user intent.
- **Key Components:**
    - **Commands/Queries:** Simple data structures that represent a request to change state or retrieve data.
    - **Handlers:** Execute the logic for a specific command or query. They use domain entities to perform business operations.
    - **Ports:** Interfaces that define the contracts for external dependencies (e.g., `ServerRepository`, `EventBus`). This is the key to the "Ports and Adapters" pattern.
    - **DTOs (Data Transfer Objects):** Used to transfer data between layers without exposing domain entities.

### 2.3. Interfaces Layer

- **Location:** `internal/interfaces/`
- **Responsibility:** This layer exposes the application's functionality to the outside world. It adapts incoming requests into application-layer commands and queries.
- **Examples:**
    - HTTP API handlers (`http/`)
    - WebSocket controllers (`ws/`)
    - gRPC services

### 2.4. Infrastructure Layer

- **Location:** `internal/infrastructure/`
- **Responsibility:** This layer provides the concrete implementations (adapters) for the ports defined in the application layer.
- **Examples:**
    - **Persistence:** A PostgreSQL or in-memory implementation of a `ServerRepository` interface.
    - **Connectors:** Clients for communicating with external MCP servers.
    - **EventBus:** An implementation using RabbitMQ or an in-memory bus.
    - **Observability:** Logging, metrics, and tracing implementations.

## 3. The Dependency Rule in Action

1. An HTTP request hits a handler in the **Interfaces** layer.
2. The handler creates a `Command` or `Query` object and dispatches it to the **Application** layer.
3. The corresponding handler in the **Application** layer is invoked.
4. The handler interacts with **Domain** entities to perform business logic.
5. If the handler needs to persist data, it calls a method on a **Port** (e.g., `serverRepo.Save(...)`).
6. The dependency injection container provides a concrete implementation of that port from the **Infrastructure** layer (e.g., `PostgresServerRepo`), which performs the actual database operation.

This structure ensures that the core logic is decoupled and testable in isolation, while allowing the technical implementation details to be swapped out easily.
