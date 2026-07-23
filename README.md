# Base Infrastructure: Universal Bootstrap Framework

[![Build Status](https://github.com/Hardik-Sankhla/base-infrastructure/actions/workflows/ci.yml/badge.svg)](https://github.com/Hardik-Sankhla/base-infrastructure/actions/workflows/ci.yml)

## Project Vision
The Universal Bootstrap Framework is a production-grade, state-driven platform engineered for bootstrapping heterogeneous environments including Linux, Windows, macOS, Raspberry Pi, Android Termux, and more. 

It replaces monolithic setup scripts by treating environments as verifiable, idempotent target states mapped through **Capabilities** instead of static tool names. 

## Problem Statement
Traditional dotfiles and setup scripts are inherently brittle. They rely on static package names (e.g. `apt install foo`), assume specific host capabilities, and fail silently or catastrophically when moved to a new operating system or container environment.

## Goals
- Provide a robust Platform Runtime built in Go.
- Abstract specific implementations into a strict Capability Model.
- Ensure tasks are safely executed via a robust Task Engine with rollback support (Planned).
- Fully decouple the Discovery Pipeline from software installations.

## Features
### Implemented (v0.4.0)
- **Core Engine & Architecture Stabilization**: Strict dependency decoupling flowing inward: `CLI -> Bootstrap -> Core -> Platform -> Discovery -> Services`.
- **Discovery Engine**: Robust, multi-stage pipeline architecture (Hardware, OS, Filesystem, Network, Environment, Software).
- **Platform Abstraction**: Native support for Linux, Windows, macOS (Darwin), BSD, and Android.
- **Capabilities Builder**: Real-time translation of system discovery data into a generalized capability matrix.

### Planned (v0.5.0)
- **State Engine**: Configuration parsing, State Model, Drift Detection, and Desired State comparison.
- **Task & Execution Engine**: Idempotent execution of state changes and automated rollbacks.

## Architecture Overview
The platform operates through a linear execution lifecycle:
1. **Discovery Pipeline**: Interrogates the host environment without making changes.
2. **Capability Engine**: Maps raw discovery data to generic platform capabilities.
3. **Execution Engine (Planned)**: Determines drift between desired state and current capabilities, scheduling tasks.

See the [Architecture Guide](docs/architecture/overview.md) for deeper technical details.

## Quick Start
To build and run the discovery pipeline locally:
```bash
git clone https://github.com/Hardik-Sankhla/base-infrastructure.git
cd base-infrastructure
make  # Or run `go build -o platform ./cmd/platform`
./platform discover
```

## Repository Structure
```
.
+-- cmd/platform/       # Ultra-thin CLI entrypoint
+-- internal/
   +-- bootstrap/      # Dependency injection and orchestration
   +-- capabilities/   # Capability translation builder
   +-- core/           # Execution engine, DAG pipeline, and domain types
   +-- discovery/      # Flat discovery stage implementations
   +-- domain/         # Core models and contracts
   +-- platform/       # OS-specific providers (Linux, Windows, etc.)
   +-- services/       # Encapsulated infrastructure (e.g. PocketBase)
+-- pkg/sdk/            # Public SDK for plugins
+-- docs/               # Comprehensive Documentation
```

## Documentation Links
- [Getting Started](docs/getting-started/quickstart.md)
- [Architecture](docs/architecture/overview.md)
- [Platform Abstraction](docs/platform/abstraction.md)
- [Discovery Engine](docs/discovery/engine.md)
- [Plugin Development](docs/plugins/overview.md)

## Contributing
We welcome contributions! Please review our [Contributing Guidelines](CONTRIBUTING.md) and our [Code of Conduct](CODE_OF_CONDUCT.md).

## License
[MIT License](LICENSE)

