# Base Infrastructure: Universal Bootstrap Framework

![Build Status](https://github.com/Hardik-Sankhla/base-infrastructure/actions/workflows/ci.yml/badge.svg)

## Overview
The Universal Bootstrap Framework is a production-grade, state-driven platform engineered for bootstrapping heterogeneous environments including Linux, Windows, macOS, Raspberry Pi, Android Termux, and more. 

It is designed to replace monolithic setup scripts by treating environments as verifiable, idempotent target states mapped through Capabilities instead of static tool names.

## Goals
- Provide a robust **Platform Runtime** built in Go.
- Abstract specific implementations into a **Capability Model**.
- Ensure tasks are safely executed via a robust **Task Engine** with rollback support.
- Fully decouple the **Discovery Pipeline** from software installations.

## Architecture
See [docs/README.md](docs/README.md) for architectural guides, ADRs, and SDK details.

## Installation
*(Coming soon)*

### Termux Setup Guide

To run the bootstrap framework on Android via Termux, you will need to prepare a Go environment:

1. Update and upgrade Termux packages:
   ```bash
   pkg update && pkg upgrade -y
   ```
2. Install Go and essential build tools:
   ```bash
   pkg install golang git make clang -y
   ```
3. Clone the repository:
   ```bash
   git clone https://github.com/Hardik-Sankhla/base-infrastructure.git
   cd base-infrastructure
   ```
4. Build the platform CLI:
   ```bash
   go build -o platform ./cmd/platform
   ```
5. Run the platform executable:
   ```bash
   ./platform --help
   ```

## License
MIT License
