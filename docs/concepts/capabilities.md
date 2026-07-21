# Capabilities

## What is a Capability?

A Capability represents a proven, active feature of the environment where the Base Infrastructure runtime is executing.

Rather than assuming an environment has specific tools (like `docker` or `git`) simply because the OS is Ubuntu, the system interrogates the host. If it finds the `docker` executable, verifies it works, and detects it in the PATH, it translates that raw discovery artifact into a standard `models.Capability`.

## Why Capabilities?

Traditional infrastructure tools are imperative. They attempt to install a package and fail if the OS isn't supported. 

Capabilities enable a **declarative, state-driven** architecture. 
If the desired state requires the `runtime.docker` Capability, the engine checks if it exists. If it does not, the Execution Engine triggers the appropriate plugin to fulfill that Capability.

## How it works

1. The **Discovery Engine** outputs a massive, unstructured JSON map called the `DiscoveryManifest`.
2. The **Capability Builder** (`internal/capabilities/builder.go`) iterates over the manifest.
3. The Builder converts strict structs (`models.SoftwareInfo`) into simple capability tags.

Example of a standard generated Capability:
```json
{
  "ID": "runtime.docker",
  "Provider": "docker",
  "Version": "24.0.5"
}
```

## When are they generated?
Capabilities are generated entirely in-memory at the very end of the `platform bootstrap` lifecycle, right before determining execution tasks.

## Where do they live?
They are defined in `internal/domain/models/capability.go`.

