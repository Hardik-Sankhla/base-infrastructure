# CLI Reference

The `platform` CLI is the primary entrypoint for interacting with the Base Infrastructure runtime.

## Global Flags
- `--debug`: Enable verbose logging and debug information.

## Commands

### `bootstrap`
Executes the environment discovery pipeline and constructs platform capabilities.

**Usage:**
```bash
./platform bootstrap
```
**Expected Output:**
Logs indicating the DAG execution of stages, followed by a printed list of identified capabilities.

---

### `sdk`
Subcommands related to plugin development and validation.

#### `sdk validate`
Validates a plugin `manifest.yaml` file to ensure it conforms to the platform SDK schemas.

**Usage:**
```bash
./platform sdk validate /path/to/manifest.yaml
```

**Arguments:**
- `path`: The absolute or relative path to the `manifest.yaml` file.

**Expected Output:**
```
INFO[0000] Manifest validated successfully
```

## Future Commands (Planned)
- `apply`: Reads a configuration file and executes tasks to match desired state.
- `plan`: Reads a configuration file and outputs a drift report without making changes.

