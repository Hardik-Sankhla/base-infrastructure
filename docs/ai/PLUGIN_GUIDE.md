# PLUGIN_GUIDE.md
# Plugin Development Guide

> Evidence-verified from `internal/runtime/plugin/manifest.go` and `plugins/*/`.
> **Last verified**: 2026-07-21

---

## What Is a Plugin?

A plugin is a directory containing:
1. A `manifest.yaml` — declares identity, compatibility, and capabilities provided
2. `detect.sh` / `detect.ps1` — fast check if the plugin is already installed
3. `install.sh` / `install.ps1` — performs the actual installation

The Go runtime does NOT execute plugins directly. It reads the manifest. Plugin execution via STDIN/STDOUT JSON-RPC is **planned but not yet implemented** (see `KNOWN_ISSUES.md#KI-002`).

---

## Manifest Schema

Defined in `internal/runtime/plugin/manifest.go:10`.

```yaml
schema_version: "1.0"    # REQUIRED — must not be empty
name: myplugin            # REQUIRED — must not be empty
version: "1.0.0"          # REQUIRED — must not be empty
description: "My plugin description"
compatibility:
  os: [linux, windows, darwin]
  arch: [amd64, arm64]
dependencies:
  - name: git
    version: "2.0.0"
provides:
  - "runtime.myplugin"
checksums:
  install.sh: "sha256:abc123..."
```

> [!WARNING]
> The bundled plugins in `plugins/*/manifest.yaml` are missing `schema_version` and will FAIL validation. This is tracked as KI-001. Do not use them as templates — use the schema above instead.

---

## Validate a Plugin Manifest

```bash
./platform sdk validate /path/to/myplugin/manifest.yaml
```

> [!WARNING]
> The `sdk validate` command is currently a stub (KI-002) — it prints the path but does NOT call `LoadManifest()`. Full validation is planned.

---

## Bundled Plugins (Source-Verified)

| Plugin | Directory | Has detect.sh | Has install.sh | manifest.yaml valid? |
|---|---|---|---|---|
| Docker | `plugins/docker/` | ✅ | ✅ | ❌ Missing `schema_version` |
| Git | `plugins/git/` | ✅ | ✅ | ❌ Missing `schema_version` |
| Python | `plugins/python/` | ✅ | ✅ | ❌ Missing `schema_version` |
| Pocketbase | `plugins/pocketbase/` | ✅ | ✅ | ❌ Missing `schema_version` |

---

## How to Create a New Plugin

1. Create `plugins/<name>/`
2. Write `manifest.yaml` following the schema above (include `schema_version: "1.0"`)
3. Write `detect.sh` — exit 0 if already installed, exit 1 if not
4. Write `install.sh` — install the software, exit 0 on success
5. For Windows support, add `detect.ps1` and `install.ps1` equivalents
6. Validate: `./platform sdk validate plugins/<name>/manifest.yaml`
7. Update `docs/plugins/supported.md`
