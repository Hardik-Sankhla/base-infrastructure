# Platform Usage Guide

The **Universal Bootstrap Framework** operates primarily through the `platform` CLI tool. As of this milestone, the engine's foundational architecture and the **Discovery Pipeline** are fully functional.

## 1. Core Concepts

* **Platform Runtime**: The execution environment that handles configuration, logging, event buses, and dependency injection.
* **Platform Detector**: Automatically identifies the underlying operating system and hardware architecture (Linux, Windows, macOS, Android/Termux) and binds the correct system abstraction layer.
* **Discovery Pipeline**: A series of immutable, prioritized stages (Hardware, OS, Network, Filesystem, Environment) that scan the host and generate a `DiscoveryManifest`.

## 2. Building the CLI

To build the `platform` binary locally:

```bash
# From the repository root
go build -o bin/platform ./cmd/platform/main.go
```

## 3. Running Discovery

To initialize the environment and run the full discovery pipeline, use the `discover` command (formerly `bootstrap`):

```bash
./bin/platform discover
```

### Expected Output
The command will execute the core discovery stages and output a human-readable summary of the detected capabilities.

```text
Platform Discovery Summary
────────────────────────────────────────

Operating System
  ubuntu 24.04
  Kernel 6.17.0-PRoot-Distro
  Architecture aarch64

Hardware
  CPU: Cortex-A53 (8 logical cores)
  Memory: 5.9 GB Total, 2.4 GB Available

Storage
  Root: 121.2 GB (102.8 GB Free)

Network
  Interfaces: 24
  Active: wlan0 (IPv4: 192.168.1.47)

Capabilities
  ✓ runtime.os
  ✓ runtime.hardware
  ✓ runtime.network
  ✓ runtime.filesystem

Discovery completed successfully.
```

If you need the raw JSON or YAML output, you can use the `--json` or `--yaml` flags:
```bash
./bin/platform discover --json
```

## 4. Current Discovery Stages

1. **Hardware (Priority 10)**: Detects CPU topology, RAM usage, Disks, and overall system architecture using `gopsutil`.
2. **OS (Priority 20)**: Detects the operating system, distribution (e.g. Debian, Ubuntu, Termux), kernel versions, package managers, and init systems (e.g. systemd).
3. **Network (Priority 30)**: Enumerates network interfaces, IP addresses (v4/v6), MAC addresses, DNS resolvers, and Proxy settings.
4. **Environment (Priority 40)**: Extracts execution context to determine if the platform is running inside a Container (Docker/LXC), a Virtual Machine (KVM/WSL/VMware), a Cloud Provider (AWS/GCP/Azure), or a CI/CD environment (GitHub Actions/GitLab CI).
5. **Filesystem (Priority 50)**: Maps mounts, capacities, search paths (`$PATH`), configuration directories (`XDG_CONFIG_HOME`), and system limits.

## 5. Integrating with the SDK

If you are writing custom plugins or modules, you can access the platform abstraction natively in Go:

```go
import (
    "context"
    "github.com/base-infrastructure/platform/internal/platform/detector"
)

func main() {
    // Detect the host platform
    plat, err := detector.New().Detect()
    if err != nil {
        panic(err)
    }

    // Access strongly-typed provider models
    cpuInfo, _ := plat.Hardware().GetCPU(context.Background())
    println("Running on", cpuInfo.Model, "with", cpuInfo.LogicalCores, "cores")
    
    envInfo, _ := plat.Environment().GetEnvironmentInfo(context.Background())
    if envInfo.IsVirtualMachine {
        println("Running inside VM:", envInfo.Virtualization)
    }
}
```
