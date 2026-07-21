# Quickstart Guide

This guide covers how to execute the platform binary to run the Discovery Engine and interrogate your local system for its active Capabilities.

## Prerequisite
Ensure you have successfully built the `platform` binary as described in the [Installation Guide](install.md).

## Running Discovery
The core functionality available in `v0.1.0` is the environment discovery mechanism.

Run the `bootstrap` command to trigger the pipeline:

```bash
./platform bootstrap
```

### Expected Output
The platform will asynchronously run discovery stages (Environment, Hardware, Network, OS, Filesystem, and Software). It will then process the resultant raw manifestations into an abstract array of `Capabilities`.

Example output:
```
INFO[0000] Discovery pipeline completed successfully
INFO[0000] Generated Platform Capabilities:
- capability: environment.os (Version: linux)
- capability: hardware.cpu (Arch: amd64, Cores: 16)
- capability: network.connectivity (IPv4: 192.168.1.100)
- capability: runtime.docker (Provider: docker, Version: 24.0.5)
```

## Validating Plugins
The platform supports modular plugin execution. If you are developing a plugin, you can validate its manifest payload against the SDK context via the `sdk` command.

```bash
./platform sdk validate --path /path/to/plugin/manifest.yaml
```

## Next Steps
Now that you have run the discovery pipeline, you can learn more about how the system builds these capabilities by reading the [Architecture Overview](../architecture/overview.md).

