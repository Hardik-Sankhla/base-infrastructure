# Deployment Guide

## Distributing the Binary
The core `platform` artifact compiles down to a single, statically-linked Go binary. 

Because the Go compiler inherently supports cross-compilation, deploying the application to a new environment does not require having Go installed on the target machine.

## Compiling for Targets
To compile the binary for specific operating systems:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o platform_linux ./cmd/platform

# Windows
GOOS=windows GOARCH=amd64 go build -o platform.exe ./cmd/platform

# macOS (Darwin)
GOOS=darwin GOARCH=arm64 go build -o platform_mac ./cmd/platform
```

## Running remotely (Future)
Future iterations of the platform (v1.0.0+) will support bootstrapping remote nodes natively.
The master instance will connect via SSH or WinRM, copy the compiled binary into the remote `/tmp` equivalent, execute it, read the JSON capabilities over stdout, and automatically orchestrate state.
