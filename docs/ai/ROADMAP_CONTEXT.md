# ROADMAP_CONTEXT.md
# Roadmap — AI Context Edition

> This is the machine-readable roadmap for AI agents.
> Unlike `docs/roadmap.md` (human-facing), this file includes source evidence
> and technical implementation hints for each item.
> **Last verified**: 2026-07-21

---

## Legend

- ✅ **Implemented** — verified in source code
- 🔄 **In Progress** — contracts/models exist, no full implementation
- 📋 **Planned** — no source evidence yet, design only
- 🔲 **Stub** — code skeleton exists but does nothing

---

## v0.1.0 — Current (Shipped)

| Feature | Status | Evidence |
|---|---|---|
| Discovery Engine orchestrator | ✅ | `internal/discovery/engine.go` |
| Pipeline (priority-sort + DFS validation) | ✅ | `internal/discovery/pipeline.go`, `validator.go` |
| OS stage | ✅ | `internal/discovery/os/stage.go` |
| Hardware stage | ✅ | `internal/discovery/hardware/stage.go` |
| Network stage | ✅ | `internal/discovery/network/stage.go` |
| Filesystem stage | ✅ | `internal/discovery/filesystem/stage.go` |
| Environment stage | ✅ | `internal/discovery/environment/stage.go` |
| Software stage | ✅ | `internal/discovery/software/stage.go` |
| Capability Builder (network + software only) | ✅ | `internal/capabilities/builder.go:22-28` |
| Platform abstraction (Linux, Windows, Darwin, Android, BSD) | ✅ | `internal/platform/*/` |
| Mock platform for testing | ✅ | `internal/platform/mock/` |
| Plugin manifest loader | ✅ | `internal/runtime/plugin/manifest.go` |
| SQLite state store | ✅ | `internal/state/db.go` |
| Event bus | ✅ | `internal/runtime/events/bus.go` |
| Task engine | ✅ | `internal/runtime/tasks/engine.go` |
| CLI: `bootstrap` | ✅ | `cmd/platform/cmd/bootstrap.go` |
| CLI: `sdk validate` (stub) | 🔲 | `cmd/platform/cmd/sdk.go:23-30` — prints only |
| CLI: `sdk create-plugin` (stub) | 🔲 | `cmd/platform/cmd/sdk.go:14-22` — prints only |

---

## v0.2.0 — Next Milestone

### High Priority (Unblock Feature Development)

| Feature | What to Build | Key Files to Touch |
|---|---|---|
| Wire up `sdk validate` | Call `plugin.LoadManifest()` in validate command | `cmd/platform/cmd/sdk.go:28` |
| Fix bundled plugin manifests | Add `schema_version: "1.0"` to all 4 plugins | `plugins/*/manifest.yaml` |
| Expand Capability Builder | Evaluate `os` and `hardware` artifacts | `internal/capabilities/builder.go` |
| PlannerEngine implementation | Implement `contracts.PlannerEngine` | `internal/planner/` |
| Plugin execution (subprocess) | STDIN/STDOUT JSON-RPC runner | `internal/runtime/plugin/runtime.go` |

### Test Coverage Improvements

| Target | Current | Goal |
|---|---|---|
| `cmd/platform` | 0% | >50% via integration tests |
| `internal/capabilities` | Unknown | >80% |
| `internal/discovery` (pipeline) | ~13% | >60% |

---

## v0.3.0 — Planned

| Feature | Contract Defined? | Implementation |
|---|---|---|
| ExecutorEngine | ✅ `engines.go:20-22` | Not started |
| ValidatorEngine | ✅ `engines.go:25-27` | Not started |
| State drift detection | Partial (`models/plan.go`) | Not started |
| Documentation CI (dead links, API drift) | No | Not started |

---

## v1.0.0+ — Future Vision

| Feature | Status |
|---|---|
| Multi-node orchestration (SSH/WinRM) | 📋 |
| Agent mode (persistent daemon) | 📋 |
| Visual dashboard | 📋 |
| Knowledge Engine (auto-generated docs) | 📋 |
