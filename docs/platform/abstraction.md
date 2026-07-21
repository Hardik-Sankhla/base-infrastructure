# Platform Abstraction

The Platform Abstraction layer (`internal/platform`) is responsible for providing a unified API over disparate Host OS environments.

## Purpose
Go handles cross-compilation incredibly well via `GOOS` and `GOARCH`, but executing system-level queries (e.g. "What is the CPU layout?") requires vastly different command-line interactions depending on the target OS. 

The Platform Abstraction guarantees that higher-level components like the Discovery Engine do not have to write `if runtime.GOOS == "windows" { ... }` in their logic.

## The Detector (`internal/platform/detector`)
When the framework boots up, `detector.DetectPlatform()` is called. 
This determines the host OS and instantiates the correct struct that satisfies the `Platform` interface.

## Provider Groups
The `Platform` interface is a composite of multiple sub-interfaces called "Providers".
- `Hardware()`
- `OS()`
- `Network()`
- `Environment()`
- `Filesystem()`
- `Software()`

By fetching the specific platform context, the system safely routes queries down to native logic.
