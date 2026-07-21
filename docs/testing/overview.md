# Testing Guidelines

Quality is a core pillar of the Base Infrastructure Platform. 

## Unit Tests
Every specific package must be paired with an adjoining `_test.go` file.
We employ Go's native `testing` module. Tests should be isolated, rapid, and predictable.

## Mocking
Located in `internal/platform/mock`, we provide a complete, statically-defined Mock platform instance.
- **Why?** Discovering hardware constraints via physical OS commands is inherently brittle in a CI environment (e.g. GitHub Actions cannot test the Windows `wmic` output natively from a Linux runner).
- The `MockPlatform` intercepts calls and provides controlled responses to test the Discovery Engine logic independently of physical host calls.

## Race Detector
While the Discovery Engine's pipeline execution is priority-sorted and sequentially executed, other subsystems (like event bus subscribers and background tasks) may operate concurrently. 

All CI pipelines execute `go test -race ./...` to guarantee memory safety across all subsystems.

