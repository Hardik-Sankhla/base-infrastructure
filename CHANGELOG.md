# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Tasks and State execution logic structure (pending).

## [0.4.0] - 2026-07-23
### Added
- **Core Architecture Stabilization**: Deep dependency decoupling flowing inward (`CLI -> Bootstrap -> Core -> Platform -> Discovery -> Services`).
- **Bootstrap Layer**: Centralized dependency injection via `internal/bootstrap`.
- **Core Execution Extraction**: The execution engine and pipelines are now extracted to `internal/core`.
- **Discovery Flattening**: Discovery capability modules are now flat files in `internal/discovery` (`stage_*.go`).
- **PocketBase Isolation**: Fully encapsulated infrastructure in `internal/services/pocketbase`.

## [0.1.0] - 2026-07-21
### Added
- **Discovery Engine**: Added DAG-based execution pipeline for system discovery.
- **Platform Interfaces**: Abstracted Hardware, OS, Filesystem, Network, Software, Process, Environment, User, Service, and Security providers.
- **Capabilities Builder**: Translates Discovery Manifests into functional environment Capabilities.
- **Built-in Stages**: Complete implementations for Network, OS, Hardware, Environment, Filesystem, and Software discovery.
- **CLI Framework**: Implemented Cobra-based `platform` CLI with `bootstrap` and `sdk` commands.
- **Mock Implementations**: Complete mock provider suite for unit testing.

