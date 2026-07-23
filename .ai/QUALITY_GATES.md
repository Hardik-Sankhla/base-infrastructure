# Quality Gates

This document defines the strict CI and verification rules that all code must pass.

## 1. Local Verification (Pre-Commit)
No code may be committed or pushed without first passing the following checks locally:
- `gofmt -s -w .`
- `gofumpt -extra -w .`
- `golangci-lint run ./...`
- `go vet ./...`
- `go test -v -race ./...`

## 2. GitHub Actions (CI)
The `ci.yml` pipeline strictly enforces the local verification steps in a clean Ubuntu environment. 

### Rules:
- The pipeline MUST remain 100% Green.
- If a CI pipeline fails, **NO FEATURE WORK** may proceed until the CI is fixed. 
- Fixing CI takes absolute precedence over all other engineering tasks.
- If staticcheck or golangci-lint fails (e.g. empty branches, unused variables, shadowing), the code must be refactored to pass. Warnings are treated as Errors.

## 3. Dependency Management
- `go mod tidy` must be run before committing.
- Dependencies must flow one way: `CLI -> Bootstrap -> Planners/Engines -> Platform -> Core Data Models`. Circular dependencies will break the build.
