# Plugin SDK Overview

The SDK (`pkg/sdk`) provides a standardized interface for external plugins to interact with the core Base Infrastructure Platform.

## Purpose
While the core platform is written in Go, infrastructure tooling is often written in bash, PowerShell, or Python. The SDK enables polyglot plugin execution via a JSON-RPC interface.

## Current Status (v0.1.0)
Currently, the SDK contains the foundational configurations required to parse and validate `manifest.yaml` files created by plugins.

### `sdk.go` CLI Command
The `sdk` command integrated into the `platform` CLI allows developers to validate their plugin manifests natively without loading the entire execution engine.

```bash
./platform sdk validate /path/to/myplugin/manifest.yaml
```

## Plugin Manifest (`manifest.yaml`)
A plugin must define a `manifest.yaml`. 

```yaml
id: "git"
name: "Git Version Control"
version: "1.0.0"
capabilities:
  provides:
    - "runtime.git"
```

