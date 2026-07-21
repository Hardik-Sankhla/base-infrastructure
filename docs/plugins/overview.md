# Plugins Architecture

Plugins enable the Platform Runtime to execute imperative operations (like installing a package) while keeping the core Go runtime entirely decoupled from specific package managers.

## Anatomy of a Plugin
A standard plugin directory contains:
- `manifest.yaml`: Describes what Capabilities the plugin can provide.
- `detect.sh` / `detect.ps1`: Optional scripts to quickly determine if the software is already installed locally.
- `install.sh` / `install.ps1`: Execution scripts that receive a JSON payload via STDIN and perform the installation.

## Execution Model (Planned)
The future execution engine will spawn the plugin's `install` script in a sub-process, pipe the context configuration to `STDIN`, and read `STDOUT` for JSON-RPC status updates.

