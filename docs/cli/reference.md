# CLI Reference

The `platform` CLI is the primary entrypoint for interacting with the Base Infrastructure runtime.

## Global Flags
- `--debug`: Enable verbose logging and debug information.

## Commands

### `discover` (alias: `bootstrap`)
Executes the environment discovery pipeline and constructs platform capabilities. Features a robust presentation layer with filtering and multiple output formats.

**Usage:**
```bash
./platform discover
./platform discover --json
./platform discover --hardware --network
```
**Expected Output:**
A formatted console summary (or structured JSON/YAML) detailing the identified capabilities and artifacts.

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

