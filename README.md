# MCPGo: The Go-Powered MCP Gateway

MCPGo is a lightweight, high-performance, and future-ready gateway that connects AI Agents with multiple Model Context Protocol (MCP) servers. It is built with a clean, domain-driven architecture and CQRS to ensure scalability and maintainability from day one.

## 🎯 Core Principles

- **Language:** Go (latest stable version)
- **Architecture:** Domain-driven, CQRS, Hexagonal (Ports & Adapters)
- **Dependency Rule:** `interfaces` → `application` → `domain`. `infrastructure` only implements ports defined in the `application` layer.
- **Simplicity:** No heavy code generation or magic. Just clean, idiomatic Go.

## ⚖️ License

This project is licensed under the **MIT License with Commons Clause restriction**. You may use this software freely for non-commercial purposes. For commercial use, please open an issue. See the [LICENSE](LICENSE) file for full details.

## 🚀 Getting Started

### Prerequisites

- Go (version 1.21 or later)
- Make

### Build

To build the application, run:

```sh
make build
```

This will create a binary at `bin/mcpgo`.

### Run

To run the gateway, use:

```sh
make run
```

The server will start on the address specified in the configuration (default: `:8080`).

### Test

To run the test suite:

```sh
make test
```

### Lint

To format and lint the code:

```sh
make lint
```

## 🏛️ Architecture

The project follows a clean architecture pattern, separating concerns into four main layers:

- **Domain:** Contains the core business logic, entities, and value objects. It has no dependencies on any other layer.
- **Application:** Orchestrates the business logic using CQRS. It defines ports (interfaces) for external dependencies like databases or event buses.
- **Interfaces:** Exposes the application to the outside world via APIs (e.g., HTTP, WebSockets).
- **Infrastructure:** Provides concrete implementations for the ports defined in the application layer (e.g., database repositories, event bus clients).

For a more detailed explanation, see [docs/architecture.md](docs/architecture.md).
