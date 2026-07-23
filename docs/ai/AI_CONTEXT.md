# AI_CONTEXT.md
# Base Infrastructure Platform — Master AI Context

> **Purpose**: This file is the entry point for any AI coding agent onboarding to this repository.
> Read this file first. It will direct you to all other relevant context files.
> **Last verified**: 2026-07-21 via full source scan.

---

## What This Project Is

A **cross-platform capability-aware bootstrap framework** written in Go.

It answers the question: *"What can this machine do?"* — and in future versions — *"What does it need to become?"*

Unlike Ansible (requires Python + SSH) or Terraform (manages remote infra), this project compiles to a **single static binary** that runs natively on the target host, interrogates the local environment, and produces a structured capability report.

---

## Repository Identity

| Field | Value |
|---|---|
| Go Module | `github.com/base-infrastructure/platform` |
| Go Version | `1.22` (source: `go.mod:3`) |
| Primary Language | Go |
| Build System | `make` + `Taskfile.yml` |
| CI | `.github/workflows/ci.yml` — GitHub Actions, ubuntu-latest |
| License | See `LICENSE` |

---

## Current Implementation Status (v0.4.0)

| Subsystem | Status |
|---|---|
| Discovery Engine | ✅ Implemented |
| Platform Abstraction (Linux/Windows/Darwin/Android/BSD) | ✅ Implemented |
| Capability Builder | ✅ Implemented (network + software only) |
| Plugin Manifest Loader | ✅ Implemented |
| Planner Engine | 🔲 Interface defined, NOT implemented |
| Executor Engine | 🔲 Interface defined, NOT implemented |
| Validator Engine | 🔲 Interface defined, NOT implemented |
| Plugin Execution (STDIN/STDOUT) | 🔲 Planned |

For full implementation status with source evidence, see [`REPOSITORY_STATE.md`](./REPOSITORY_STATE.md).

---

## Critical Architecture Rules

1. **No OS-specific code outside `internal/platform/`** — enforced by `CONSTITUTION.md`
2. Discovery stages must implement `discovery.Stage` interface (`internal/discovery/stage.go`)
3. The Capability Builder reads `DiscoveryManifest.Artifacts` via type assertions — only recognized types produce capabilities
4. `PlatformContext` is the single dependency injection root — all engines receive it
5. Cycle detection uses DFS (`internal/discovery/validator.go:43`) — NOT Kahn's algorithm

For all engineering rules, see [`ENGINEERING_RULES.md`](./ENGINEERING_RULES.md).

---

## Key AI Context Files

| File | Purpose |
|---|---|
| `AI_CONTEXT.md` | This file — master entry point |
| `CURRENT_STATUS.md` | What works right now + what doesn't |
| `REPOSITORY_STATE.md` | Source-verified implementation status |
| `PACKAGE_INDEX.md` | Every package, file, dependency |
| `ARCHITECTURE_INDEX.md` | Interface → implementation → test with line numbers |
| `ENGINEERING_RULES.md` | 10 non-negotiable invariants |
| `CODING_RULES.md` | Formatting, error handling, naming, imports |
| `TESTING_RULES.md` | Coverage gaps, mock requirements, race detector |
| `CI_RULES.md` | All 13 CI steps documented |
| `PLUGIN_GUIDE.md` | Manifest schema + plugin development |
| `AGENT_GUIDE.md` | Certification pipeline for AI agents |
| `ROADMAP_CONTEXT.md` | Machine-readable roadmap with source evidence |
| `DECISION_LOG.md` | 7 architectural decisions with context |
| `KNOWN_ISSUES.md` | Verified bugs registry |

---

## Before You Write Any Code

You MUST read:
1. [`ENGINEERING_RULES.md`](./ENGINEERING_RULES.md)
2. [`ARCHITECTURE_INDEX.md`](./ARCHITECTURE_INDEX.md)
3. [`KNOWN_ISSUES.md`](./KNOWN_ISSUES.md)
4. [`AGENT_GUIDE.md`](./AGENT_GUIDE.md) — the certification pipeline

**Never modify code without completing the Repository Certification Pipeline.**
