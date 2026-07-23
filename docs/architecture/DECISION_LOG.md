# DECISION_LOG.md
# Architecture Decision Log

> This document records key decisions made during the project's evolution,
> with context and rationale. For formal ADRs, see `docs/adr/`.
> **Last updated**: 2026-07-21

---

## DL-001: Go as the Implementation Language
**Date**: Project inception  
**ADR**: `docs/adr/0001-go.md`  
**Decision**: Go chosen for static compilation, cross-compilation, native concurrency.  
**Key implication**: Plugin execution requires a subprocess boundary since Go is compiled — cannot dynamically load shell scripts as Go packages.

---

## DL-002: Priority-Sort Instead of Kahn's Algorithm
**Date**: 2026-07-21 (discovered via source audit)  
**Decision**: Pipeline uses Priority-integer insertion sort (`pipeline.go:253`), not topological sort.  
**Context**: Documentation incorrectly claimed Kahn's algorithm was used. Source verification proved it false.  
**Key implication**: Stages do not execute in strict dependency tiers concurrently — they execute sequentially by priority. Parallelism is NOT currently implemented.

---

## DL-003: DFS for Cycle Detection, Not Kahn's BFS
**Date**: 2026-07-21 (discovered via source audit)  
**Decision**: `validator.go` uses recursive DFS `hasCycle` closure.  
**Key implication**: The validator detects cycles and validates dependencies but does NOT produce an execution tier ordering — that is purely Priority-based.

---

## DL-004: PlatformContext as Unified DI Root
**Date**: v0.1.0 design  
**ADR**: Inline in `ENGINEERING_RULES.md#Rule-6`  
**Decision**: All engines receive `*context.PlatformContext`. No engine may take individual dependencies.  
**Key implication**: Adding any new runtime capability (e.g., secrets manager) requires adding a field to `PlatformContext`.

---

## DL-005: SQLite for State Persistence
**Date**: v0.1.0 design  
**Decision**: `modernc.org/sqlite` is the embedded state store (see `internal/state/db.go`).  
**Key implication**: The `DB *sql.DB` field in `PlatformContext` is nullable — passing `nil` is valid for runs that don't require state persistence.

---

## DL-006: VitePress for Documentation Website
**Date**: 2026-07-21  
**Decision**: VitePress chosen over Docusaurus, MkDocs, or Hugo.  
**Rationale**: Markdown-native, zero-backend, TypeScript config, excellent search, used by major open-source Go-adjacent projects. No MDX/JSX required.  
**Key implication**: `srcDir: '../docs'` means the `docs/` folder is the single Markdown source — the website just renders it.

---

## DL-007: AI Context as First-Class Subsystem
**Date**: 2026-07-21  
**Status**: Migrated to Engineering OS v2.0 (`.ai/`).
**Rationale**: AI Agents need machine-readable context (rules, constraints, states) that shouldn't clutter human documentation.
**Decision**: `.ai/` directory established as the machine-readable Engineering OS.  

**Key implication**: Every AI agent MUST read `.ai/START_HERE.md` before modifying code. This is enforceable via `.ai/ENGINEERING_LIFECYCLE.md`.
