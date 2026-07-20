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

## License
MIT License
