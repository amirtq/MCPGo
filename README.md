# MCPGo: The Go-Powered MCP Gateway

MCPGo is a lightweight, high-performance, and future-ready gateway that connects AI Agents with multiple Model Context Protocol (MCP) servers. It is built with a clean, domain-driven architecture and CQRS to ensure scalability and maintainability from day one.

## 🎯 Core Principles

- **Language:** Go (latest stable version)
- **Architecture:** Feature-Sliced Design
- **Dependency Rule:** Code is organized by feature, not by technical layer.
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

The server will start on `https://localhost:443`.

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

The project follows a **Feature-Sliced Design**. Instead of being organized by technical layers, the code is grouped by feature. This enhances modularity and makes the application easier to develop and maintain.

The main components are:
- **`backend/apps`**: Contains the core business logic for each feature.
- **`backend/api`**: Exposes the features via an HTTP API.
- **`backend/services`**: Holds shared, cross-cutting concerns like SSL management.

For a more detailed explanation, see [docs/architecture.md](docs/architecture.md).
