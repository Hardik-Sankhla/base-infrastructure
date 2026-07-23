# Architecture Rules

These rules represent the non-negotiable architectural boundaries of the system.

## Rule 1: Zero Trust — Never Assume
Every architectural claim must be backed by:
- A source file path + line number, OR
- An executed command + its output, OR
- A test result

If evidence cannot be produced, explicitly mark the claim as `UNVERIFIED` and do not act on it.

## Rule 2: Platform Abstraction is Sacred
**THE MOST CRITICAL RULE.**

- ✅ ALLOWED: `internal/platform/linux/...` → platform specific implementations.
- ✅ ALLOWED: `internal/platform/windows/...` → platform specific implementations.
- ❌ FORBIDDEN: `internal/core/...` or `internal/planner/...` → ANY OS-specific imports, `runtime.GOOS` checks, or hardcoded shell commands.

The `internal/platform/` boundary is the ONLY place where the system may interact with or branch on the OS type.

## Rule 3: Single Dependency Injection Root
All engines receive `*context.PlatformContext` (or domain equivalent). This is the single dependency injection point.
- Do not add new global singletons.
- Do not add new function-level dependencies.
- If an engine needs a new dependency, add it to the Context object appropriately.

## Rule 4: ADRs for Architecture Changes
Any change to:
- Pipeline execution model
- Capability derivation logic
- Platform abstraction boundaries
- Core engine orchestration
...requires a new Architecture Decision Record in `docs/adr/`.
