# Roadmap

This document outlines the strategic roadmap for the Universal Bootstrap Framework. Features are categorized strictly by their current implementation status.

## ✅ Implemented (v0.4.0)
The foundation of the project has been established with a heavy focus on the Discovery Engine, Platform abstraction, and Core Architecture Stabilization.

- **Platform Abstraction Layer**: Built-in support and interfaces for Linux, Windows, macOS, Android (Termux), and BSD.
- **Provider Interfaces**: Standardization of OS, Hardware, Network, Software, Environment, Filesystem, Process, Security, User, and Service providers.
- **Discovery Pipeline Engine**: A highly concurrent, dependency-aware DAG execution engine for discovery stages.
- **Core Discovery Stages**: Implemented and tested stages for Environment, Hardware, Network, OS, Filesystem, and Software.
- **Capability Builder**: Translates raw `DiscoveryManifest` outputs into a structured list of platform `Capabilities`.
- **Plugin Infrastructure (Alpha)**: Foundational JSON-RPC sub-process execution model for external scripts.
- **Core Architecture Freeze**: Strict one-way dependency decoupling flowing inward (`CLI -> Bootstrap -> Core -> Platform -> Discovery -> Services`). PocketBase is perfectly isolated as a service capability.

## 🚀 In Progress (v0.5.0)
Currently under active development by contributors.

- **v0.4 Architecture Stabilization**: ✅ Deep dependency decoupling and capability builder stabilization.
- **Engineering OS v2**: ✅ Rebuilding the agent governance framework (`.ai/`) for scalable long-term development.
- **Engineering OS Stabilization**: ⏳ Synchronizing human documentation, website navigation, and CI gates to match the new OS.
- **State Engine**: Configuration parsing and State Model building (Paused during OS Stabilization).
- **Drift Detection**: Comparing desired configuration against current Capabilities to determine required execution graphs.

## 🗓️ Planned (v0.6.0)
Features planned for the medium term.

- **Task & Execution Engine**: Development of the core state execution model (Idempotent task runs).
- **Dependency Resolution**: Task Graph creation based on execution plan.
- **Rollback Engine**: Safely reversing failed state changes.
- **Secure Secret Management**: Integrating secure local/remote keystores for sensitive credentials during bootstrap.
- **Remote Telemetry**: Optional telemetry and logging for remote dashboard visibility.

## ?? Future (v1.0.0+)
Long-term vision for a mature framework.

- **Multi-node Orchestration**: Bootstrapping remote nodes via SSH/WinRM from a master controller.
- **Agent Mode**: Running as a persistent background daemon for continuous state reconciliation.
- **Visual Dashboard**: Web-based UI to monitor bootstrap progress and capability status.

