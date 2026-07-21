# CODING_RULES.md
# Coding Standards & Conventions

> Source-verified from `.golangci.yml`, CI workflow, and existing codebase patterns.
> **Last verified**: 2026-07-21

---

## Formatting (Non-Negotiable)

All Go files MUST pass both formatters before every commit:

```bash
gofmt -s -w .        # Standard Go formatter + simplification
gofumpt -extra -w .  # Stricter superset of gofmt
```

CI runs both and diffs the result. Any formatting deviation fails the build.
Source: `.github/workflows/ci.yml:28-54`

---

## Linting

`golangci-lint v1.56.2` is the linter. Configuration in `.golangci.yml`.

```bash
golangci-lint run ./...
```

Do not suppress lint errors with `//nolint` without a documented justification comment.

---

## Error Handling

**Always wrap errors** with context:

```go
// ✅ Correct
return fmt.Errorf("failed to read manifest at %s: %w", path, err)

// ❌ Wrong
return err
```

Never swallow errors silently. Every `err != nil` branch must either return the error, log it with sufficient context, or handle it deliberately with a comment explaining why it is safe to ignore.

---

## Naming Conventions

Follow standard Go conventions:

| Type | Convention | Example |
|---|---|---|
| Exported types | `PascalCase` | `DiscoveryManifest` |
| Unexported | `camelCase` | `stageMap` |
| Interfaces | noun or adjective | `Platform`, `Stage`, `Detector` |
| Constructors | `New<Type>` | `NewPlatformContext`, `NewBuilder` |
| Test files | `<file>_test.go` | `stage_test.go` |
| Test functions | `Test<Subject>_<Scenario>` | `TestValidator_CircularDependency` |

---

## Package Boundaries

| Package | Allowed to import | Forbidden |
|---|---|---|
| `internal/domain/models` | stdlib only | Anything from `internal/` |
| `internal/domain/contracts` | stdlib, `models` | Platform packages |
| `internal/discovery` | `domain/*`, `runtime/*` | `platform/*` directly |
| `internal/platform/*` | `domain/models`, stdlib | `discovery/*`, `capabilities/*` |
| `internal/capabilities` | `domain/models` | `platform/*`, `discovery/*` |
| `cmd/platform` | All internal | No direct `platform/*` calls |

---

## Comments & Documentation

All **exported** functions, types, methods, and constants MUST have a Go doc comment:

```go
// Builder translates a DiscoveryManifest into a set of functional Capabilities.
type Builder struct { ... }

// Build evaluates the discovery manifest and generates a list of capabilities.
func (b *Builder) Build() []models.Capability { ... }
```

Unexported helpers should have comments if their purpose is non-obvious.

---

## Context Propagation

Always thread `context.Context` through functions that do I/O or could block:

```go
// ✅ Correct
func (s *Stage) Run(ctx context.Context, dctx discovery.Context) (DiscoveryArtifact, error)

// ❌ Wrong — no context
func (s *Stage) Run() (DiscoveryArtifact, error)
```

Respect context cancellation — check `ctx.Err()` in loops.

---

## Avoiding Common Mistakes

```go
// ❌ Never use runtime.GOOS outside internal/platform/
if runtime.GOOS == "windows" { ... }  // FORBIDDEN in discovery/capabilities/cmd

// ❌ Never create capabilities outside capabilities/builder.go
caps = append(caps, models.Capability{...})  // Only valid in builder.go

// ❌ Never mutate a DiscoveryArtifact after it is returned from Run()
artifact.SomeField = "modified"  // Artifacts are immutable post-validation

// ✅ Always use the stage interface — never cast to concrete types
stage.Name()  // ✅
stage.(*os.Stage).InternalField  // ❌
```
