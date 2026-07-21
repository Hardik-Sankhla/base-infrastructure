# Platform Providers

Providers form the granular interface boundaries that isolate specific environmental functionality. 

## Implemented Providers

### OSProvider
Responsible for determining `models.OSInfo`.
- **Properties**: `Family`, `Name`, `Version`, `Architecture`, `Kernel`.
- **Linux Impl**: Parses `/etc/os-release` and `uname`.
- **Windows Impl**: Uses `wmic os get` and registry keys.

### HardwareProvider
Responsible for determining `models.HardwareInfo`.
- **Properties**: `CPU` (Model, Cores, Threads), `Memory` (Total, Free).
- **Linux Impl**: Parses `/proc/cpuinfo` and `/proc/meminfo`.
- **Windows Impl**: Uses `wmic cpu` and `wmic memorychip`.

### NetworkProvider
Responsible for determining `models.NetworkInfo`.
- **Properties**: Network Interfaces (`Name`, `MAC`, `IPv4`, `IPv6`, `IsUp`).
- **Standard Impl**: Relies heavily on Go's native `net.Interfaces()` which is inherently cross-platform.

### EnvironmentProvider
Responsible for determining `models.EnvironmentInfo`.
- **Properties**: `Variables` (map of Env Vars), `Path` (parsed `$PATH`), `User` (current executing user ID).

### FilesystemProvider
Responsible for determining `models.FilesystemInfo`.
- **Properties**: `Mounts` (device paths, mount points, available space).
- **Linux Impl**: Parses `/proc/mounts`.
- **Windows Impl**: Uses `wmic logicaldisk`.

### SoftwareProvider
Responsible for determining `models.SoftwareInfo`.
- **Properties**: `Runtimes`, `Tools` (e.g. Docker, Python, Git).
- **Standard Impl**: Uses `exec.LookPath` across OS implementations to locate binaries.

## Planned Providers
- `ProcessProvider`: Running processes and PIDs.
- `SecurityProvider`: Firewall rules, SE-Linux states, Defender status.
- `UserProvider`: User and Group ACL management.
- `ServiceProvider`: Systemd/Windows Services interrogation.
