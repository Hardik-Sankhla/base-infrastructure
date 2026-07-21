# Roadmap

This document outlines the strategic roadmap for the Base Infrastructure Platform. Features are categorized strictly by their current implementation status.

## ✅ Implemented (v0.1.0)

- **Platform Abstraction Layer**: Built-in support for Linux, Windows, macOS, Android (Termux), and BSD.
- **Provider Interfaces**: OS, Hardware, Network, Software, Environment, Filesystem providers.
- **Discovery Pipeline Engine**: Priority-sorted, DFS-validated discovery pipeline.
- **Core Discovery Stages**: Environment, Hardware, Network, OS, Filesystem, Software.
- **Capability Builder**: Translates `DiscoveryManifest` into `[]Capability`.
- **Plugin Infrastructure (Alpha)**: Manifest schema and loader (`LoadManifest`).
- **CLI Commands**: `bootstrap`, `sdk validate`, `sdk create-plugin`, `sdk test`.
- **AI Context Layer**: `docs/ai/` directory with machine-readable context for AI agents.
- **Documentation Website**: VitePress site deployed to GitHub Pages.

## 🔄 In Progress (v0.2.0)

- **Task & Execution Engine**: Core state execution model (idempotent task runs).
- **Configuration Parser**: Loading desired states via YAML configurations.
- **Plugin Execution**: STDIN/STDOUT JSON-RPC mechanism for external plugin scripts.
- **Capability Builder Expansion**: Evaluate `os`, `hardware`, `filesystem`, `environment` artifacts.

## 📋 Planned (v0.3.0)

- **State Drift Detection**: Diff desired config against current Capabilities.
- **Rollback Engine**: Safely reverse failed state changes.
- **Secure Secret Management**: Local/remote keystore integration.
- **Documentation CI**: Dead-link checking, API drift detection, architecture drift alerts.

## 🔮 Future (v1.0.0+)

- **Multi-node Orchestration**: Bootstrap remote nodes via SSH/WinRM.
- **Agent Mode**: Persistent background daemon for continuous state reconciliation.
- **Visual Dashboard**: Web-based UI for bootstrap progress and capability status.
- **Knowledge Engine**: Automated documentation generation from source analysis.
