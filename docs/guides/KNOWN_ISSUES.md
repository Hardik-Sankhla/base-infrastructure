# KNOWN_ISSUES.md
# Known Issues & Inconsistencies

> This document tracks verified bugs, inconsistencies, and technical debt.
> Every entry is evidence-backed. Do not add unverified issues.

---

## KI-001: Plugin Manifest Schema Mismatch

**Severity**: High — causes CLI validation failures on all bundled plugins  
**Status**: ✅ RESOLVED 2026-07-21 — added `schema_version: "1.0"` to all 4 bundled manifests  
**Discovered**: 2026-07-21 via source audit

### Evidence

The `platform sdk validate` command calls `plugin.LoadManifest()` which requires:
- `schema_version` (source: `internal/runtime/plugin/manifest.go:44`)
- `name` (source: `manifest.go:47`)
- `version` (source: `manifest.go:50`)

The bundled plugins only contain:
```yaml
# plugins/docker/manifest.yaml
name: docker
version: 1.0.0
dependencies: []
```

`schema_version` is missing, causing `LoadManifest()` to return:
```
error: missing schema_version
```

**Note**: The `platform sdk validate` CLI command is currently a stub that only prints the path and does NOT yet call `LoadManifest()`. But when it is wired up, all bundled plugins will fail.

### Resolution
Either:
- Add `schema_version: "1.0"` to all bundled plugin `manifest.yaml` files, OR
- Make `schema_version` optional in `manifest.go`

---

## KI-002: `platform sdk validate` Is a Stub

**Severity**: Medium — misleading documentation  
**Status**: ✅ RESOLVED 2026-07-21 — `sdk validate` now calls `plugin.LoadManifest()` and prints structured output  
**Discovered**: 2026-07-21 via source audit

### Evidence

`cmd/platform/cmd/sdk.go:28`:
```go
Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Validating manifest at: %s\n", args[0])
    // Implementation for manifest validation
},
```

The `validate` command prints the path but does NOT call `plugin.LoadManifest()`. No actual validation occurs.

---

## KI-003: Software Discovery Returns Empty on Some Platforms

**Severity**: Medium  
**Status**: Open (observed in terminal output)  
**Discovered**: 2026-07-21 from user-provided execution log

### Evidence

From user-provided bootstrap output:
```json
"software": {}
```

The software stage returned an empty struct, meaning no runtimes or tools were discovered. This may be expected in sandboxed/Termux environments where standard binary paths differ.

---

## KI-004: Capability Builder Does Not Evaluate `os`, `hardware`, `filesystem`, `environment` Artifacts

**Severity**: Low (planned feature gap, not a bug)  
**Status**: Open — by design, known gap  
**Discovered**: 2026-07-21 via source audit

### Evidence

`internal/capabilities/builder.go:22-28`:
```go
func (b *Builder) Build() []models.Capability {
    var caps []models.Capability
    caps = append(caps, b.evaluateNetwork()...)
    caps = append(caps, b.evaluateSoftware()...)
    return caps
}
```

Only `network` and `software` artifacts are evaluated. The `os`, `hardware`, `filesystem`, and `environment` artifacts exist in the manifest but produce zero capabilities.

---

## KI-005: Test Coverage is Low

**Severity**: Medium  
**Status**: Open  
**Discovered**: 2026-07-21 from CI logs

| Package | Coverage |
|---|---|
| `cmd/platform` | 0.0% |
| `internal/config` | 0.0% |
| `internal/discovery` (overall) | 13.4% |

The validator tests are present but much of the pipeline, engine, and capability logic is untested.
