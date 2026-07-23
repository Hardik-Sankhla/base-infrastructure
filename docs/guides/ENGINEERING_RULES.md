# ENGINEERING_RULES.md
# Engineering Invariants — Non-Negotiable

> Derived from `CONSTITUTION.md` and source-verified architectural patterns.
> Any AI agent or human contributor MUST follow these rules without exception.

---

## Rule 1: Zero Trust — Never Assume
Every architectural claim must be backed by:
- A source file path + line number, OR
- An executed command + its output, OR
- A test result

If evidence cannot be produced, explicitly mark the claim as **`UNVERIFIED`** and do not act on it.

---

## Rule 2: Platform Abstraction is Sacred

**THE MOST CRITICAL RULE.**

```
✅ ALLOWED:  internal/platform/linux/os.go → calls exec.Command("uname")
✅ ALLOWED:  internal/platform/windows/os.go → calls wmic commands
❌ FORBIDDEN: internal/discovery/os/stage.go → any runtime.GOOS check
❌ FORBIDDEN: internal/capabilities/builder.go → any os.Getenv("OS") check
```

The `internal/platform/` boundary is the ONLY place where the system may branch on OS type.
Source: `CONSTITUTION.md §2`, verified by reviewing `internal/discovery/*/stage.go` — none contain `runtime.GOOS`.

---

## Rule 3: Stage Interface Compliance

Every discovery stage MUST implement `discovery.Stage` (`internal/discovery/stage.go`).

Required methods:
- `Name() string`
- `Version() string`
- `Description() string`
- `Priority() int`
- `DependsOn() []string`
- `Timeout() time.Duration`
- `Initialize(Context) error`
- `Run(context.Context, Context) (DiscoveryArtifact, error)`
- `Validate(DiscoveryArtifact) error`
- `Cleanup(context.Context) error`

A stage that does not implement all methods will not compile. No exceptions.

---

## Rule 4: Artifacts Are Immutable

Once a stage's `Run()` returns a `DiscoveryArtifact`, that artifact is stored in the result and must not be mutated.

The pipeline at `internal/discovery/pipeline.go:133` stores the artifact in `StageResult.Artifact` immediately after validation.

---

## Rule 5: Capabilities Derive From Discovery

The Capability Builder (`internal/capabilities/builder.go`) is the ONLY place that maps discovery data to `models.Capability` structs. No other component may create capabilities directly.

---

## Rule 6: PlatformContext is the DI Root

All engines receive `*context.PlatformContext`. This is the single dependency injection point.
Do not add new global singletons. Do not add new function-level dependencies.
If an engine needs a new dependency, add it to `PlatformContext`.
Source: `internal/runtime/context/context.go:17-29`

---

## Rule 7: Quality Gate is Mandatory

No code may be merged if any of the following fail:
```bash
go mod tidy
go test -race ./...
golangci-lint run ./...
gofmt -s -w .
gofumpt -extra -w .
```
Source: `.github/workflows/ci.yml:28-63`

---

## Rule 8: Documentation Must Match Code

**Documentation is infrastructure.**

If a PR changes:
- Any interface in `internal/domain/contracts/` → update `.ai/ARCHITECTURE_RULES.md`
- Any platform provider → update `.ai/MEMORY_PROTOCOL.md`
- Any CLI command behavior → update `docs/cli/reference.md`
- Any architectural decision → add an ADR to `docs/adr/`
- Any known bug → update `docs/guides/KNOWN_ISSUES.md`

---

## Rule 9: ADRs for Architecture Changes

Any change to:
- Pipeline execution model
- Capability derivation logic
- Platform abstraction boundaries
- Plugin execution model
- PlatformContext fields

...requires a new Architecture Decision Record in `docs/adr/`.

---

## Rule 10: Plugin Schema Validation

The `platform sdk validate [path]` command uses `internal/runtime/plugin/manifest.go:LoadManifest()`.
Valid manifests MUST contain `schema_version`, `name`, and `version`.
The bundled plugins currently fail this validation. See `KNOWN_ISSUES.md`.
