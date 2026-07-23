# CURRENT_STATUS.md
# Current Implementation Status Snapshot

> This document is a point-in-time snapshot of what exists, what works, and what doesn't.
> Update this file whenever a major feature ships or a known issue is resolved.
> **Snapshot date**: 2026-07-23 | **Version**: v0.4.0

---

## What Works Right Now

Run these and they will succeed:

```bash
# Build
go build -o platform ./cmd/platform

# Run discovery
./platform bootstrap
# Output: JSON manifest + capabilities array

# Lint + format
gofmt -s -w .
gofumpt -extra -w .
golangci-lint run ./...

# Tests
go test -race ./...
```

### `platform bootstrap` produces:
- Full `DiscoveryManifest` with `os`, `hardware`, `network`, `filesystem`, `environment` artifacts
- `capabilities` array containing `network.connectivity` and/or `runtime.*` entries if software is detected

---

## What Does Not Work Yet

| Feature | Status | Issue |
|---|---|---|
| `platform sdk validate` actually validates | ❌ Stub | KI-002 |
| Plugin execution (install scripts) | ❌ Not implemented | KI-001, KI-002 |
| Bundled plugins pass `sdk validate` | ❌ Schema mismatch | KI-001 |
| Planner, Executor, Validator engines | ❌ Interface only | — |
| `os`/`hardware`/`filesystem` capabilities | ❌ Not evaluated | KI-004 |
| Documentation website deploy | ⚠️ Was failing (fix in progress) | Workflow fixed 2026-07-21 |

---

## Repository Health at a Glance

```
Go Build:          ✅ PASS
Go Tests (race):   ✅ PASS
golangci-lint:     ✅ PASS
gofmt/gofumpt:     ✅ PASS
Cross-builds:      ✅ linux/amd64, windows/amd64, darwin/arm64
Docs CI:           ⚠️ Fix applied 2026-07-21 — pending next push
GitHub Pages:      ⚠️ Requires GitHub Pages → Source: GitHub Actions to be enabled
```

---

## Next Priority Actions

1. **Enable GitHub Pages** (Settings → Pages → Source: GitHub Actions)
2. **Fix bundled plugin manifests** — add `schema_version: "1.0"` to all 4 plugins
3. **Wire up `sdk validate`** — call `plugin.LoadManifest()` in the validate command
4. **Expand Capability Builder** — evaluate `os` and `hardware` artifacts
5. **Increase test coverage** — especially `cmd/platform` (currently 0%)
