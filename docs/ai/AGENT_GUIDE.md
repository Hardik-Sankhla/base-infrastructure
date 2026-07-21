# AGENT_GUIDE.md
# AI Agent Onboarding Guide

> This guide tells an AI coding agent exactly how to operate on this repository safely.
> Follow this guide completely before writing a single line of code.

---

## Step 1: Read Context Files (Mandatory)

In this order:
1. `docs/ai/AI_CONTEXT.md` — Master overview
2. `docs/ai/REPOSITORY_STATE.md` — What is and is not implemented
3. `docs/ai/ENGINEERING_RULES.md` — What you must never do
4. `docs/ai/ARCHITECTURE_INDEX.md` — Where everything lives
5. `docs/ai/KNOWN_ISSUES.md` — What is broken and why

---

## Step 2: Repository Certification

Before modifying any code, you MUST verify:

### 2a. Build Certification
Confirm that `go build ./...` passes without errors.
If you cannot run it directly, use the last verified CI run from `.github/workflows/ci.yml`.

### 2b. Test Certification
Confirm that `go test -race ./...` passes without failures.

### 2c. Documentation Certification
Identify any drift between `docs/` and source code.
Mark drifted claims as `UNVERIFIED` until corrected.

---

## Step 3: Understand the Change Scope

For every task, answer these questions before starting:

| Question | Why |
|---|---|
| Which package(s) does this touch? | Understand blast radius |
| Does this cross a package boundary? | May require new interface |
| Does this add OS-specific logic? | Must go inside `internal/platform/` |
| Does this change a public interface? | Must update `ARCHITECTURE_INDEX.md` |
| Does this change CLI behavior? | Must update `docs/cli/reference.md` |
| Does this change architecture? | Must write an ADR |

---

## Step 4: Execute Atomically

Follow these rules during implementation:
- One logical change per commit
- Each commit must leave the build and tests green
- Never commit a half-refactored state

---

## Step 5: Update Documentation

After every code change, update:
- `docs/ai/REPOSITORY_STATE.md` if any engine/provider status changed
- `docs/ai/ARCHITECTURE_INDEX.md` if any interface or implementation changed
- The relevant `docs/` subsystem file if behavior changed
- `docs/ai/KNOWN_ISSUES.md` if a bug was introduced or fixed

---

## Step 6: Produce a Summary

End every task with a summary covering:
- What was changed and why
- What files were modified
- What documentation was updated
- What was explicitly NOT changed (and why)

---

## Anti-Patterns to Avoid

| Anti-Pattern | Why It's Dangerous |
|---|---|
| Adding `runtime.GOOS` outside `internal/platform/` | Violates the core architectural invariant |
| Documenting planned features as implemented | Creates hallucinated source of truth |
| Assuming a test exists without checking the file | Coverage gaps cause silent regressions |
| Making multiple large commits | Hard to review and revert |
| Modifying code without running tests | May break working subsystems |
| Calling `LoadManifest()` on bundled plugins | Will fail — schema mismatch bug |

---

## How to Add a New Discovery Stage

1. Create `internal/discovery/<stagename>/stage.go`
2. Implement all 10 methods of the `discovery.Stage` interface
3. Create `internal/discovery/<stagename>/stage_test.go` using `internal/platform/mock`
4. Register the stage in `internal/discovery/builtin/stages.go`
5. Add the artifact type to `internal/domain/models/` if new data is returned
6. Update `internal/capabilities/builder.go` if the artifact should generate capabilities
7. Update `docs/ai/REPOSITORY_STATE.md`
8. Update `docs/discovery/stages.md`

---

## How to Add a New OS Platform

1. Create `internal/platform/<osname>/` directory
2. Implement all required providers: `HardwareProvider`, `OSProvider`, `FilesystemProvider`, `NetworkProvider`, `EnvironmentProvider`, `SoftwareProvider`
3. Implement the `Platform` interface (`internal/platform/platform.go:54`)
4. Register in `internal/platform/detector/detector.go`
5. Update `docs/platform/environments.md`
6. Update `docs/ai/REPOSITORY_STATE.md`
