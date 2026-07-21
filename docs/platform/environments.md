# Target Environments

The Base Infrastructure Platform is intentionally designed to be cross-platform, allowing seamless bootstrapping on any modern system.

## Supported OS Matrix

| Operating System | Detector Token | Status | Implementation Directory |
|------------------|----------------|--------|--------------------------|
| Linux            | `linux`        | Stable | `internal/platform/linux`|
| Windows          | `windows`      | Stable | `internal/platform/windows`|
| macOS (Darwin)   | `darwin`       | Beta   | `internal/platform/darwin`|
| Android (Termux) | `android`      | Beta   | `internal/platform/android`|
| BSD / pfSense    | `bsd`          | Alpha  | `internal/platform/bsd`|

## Mock Platform
In addition to physical environments, a special `MockPlatform` (`internal/platform/mock`) exists strictly for Unit Testing. 

The Mock Platform implements all interfaces using stubbed maps and hardcoded values, ensuring CI/CD can test the Discovery Engine logic without requiring physical hardware detection capabilities.

