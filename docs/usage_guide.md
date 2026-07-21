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

To initialize the environment and run the full discovery pipeline, use the `bootstrap` command:

```bash
./bin/platform bootstrap
```

### Expected Output
The command will execute the core discovery stages and output a JSON representation of the `DiscoveryManifest`.

```json
{
  "id": "e3b0c442...",
  "start_time": "2026-07-20T10:00:00Z",
  "end_time": "2026-07-20T10:00:02Z",
  "duration": 2000000000,
  "platform": "linux",
  "stages": [
    {
      "name": "hardware",
      "status": "success",
      "duration": 15000000
    },
    {
      "name": "os",
      "status": "success",
      "duration": 5000000
    }
  ],
  "artifacts": {
    "Hardware": {
      "cpu": {
        "vendor": "GenuineIntel",
        "model": "Intel Core i7",
        "physical_cores": 4,
        "logical_cores": 8
      },
      "ram": {
        "total_bytes": 17179869184,
        "available_bytes": 8589934592
      }
    },
    "OS": {
      "operating_system": "linux",
      "distribution": "ubuntu"
    },
    "Filesystem": {
      "root_capacity": {
        "total_bytes": 500000000000,
        "free_bytes": 250000000000
      }
    },
    "Network": {
      "interfaces": [
        {
          "name": "eth0",
          "ipv4": ["192.168.1.100"]
        }
      ]
    },
    "Environment": {
      "is_virtual_machine": true,
      "virtualization": "wsl",
      "is_container": false
    }
  }
}
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
