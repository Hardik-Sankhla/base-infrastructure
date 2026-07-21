# Base Infrastructure Constitution

This document defines the non-negotiable engineering rules and governance model for the Base Infrastructure repository. All human contributors and autonomous AI agents MUST adhere to these rules.

## 1. Zero Trust Implementation
Never assume a pattern exists. Every architectural claim must be verified by source code, execution logs, or test results. 

## 2. Platform Abstraction Strictness
Absolutely **NO** OS-specific code (e.g., `runtime.GOOS`, `exec.Command("uname")`) is permitted outside of the `internal/platform/` package. The core discovery pipeline must remain OS-agnostic and rely exclusively on the `Platform` interfaces.

## 3. Documentation Parity
Every feature must include corresponding updates to the `docs/` hierarchy. Documentation must be perfectly aligned with the implemented code. 
- Do not document future roadmap items as implemented features.

## 4. Quality Gates
No code may be merged if any of the following fail:
- `go mod tidy`
- `go test -race ./...`
- `golangci-lint run`
- `gofmt` and `gofumpt` checks

## 5. Architecture Decision Records (ADR)
Any changes to the pipeline execution, capability building, or platform abstraction boundaries require the creation of an ADR in `docs/adr/`.

## 6. Public APIs
All interfaces located in `internal/domain/contracts/` must be thoroughly documented with Go docstrings.

## 7. Mandatory Certification
Every AI agent onboarding onto this repository MUST complete the Repository Certification Pipeline before modifying any code. Proof of certification must be stored in the `reports/` directory.

