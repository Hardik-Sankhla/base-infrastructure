# Verification Checklist

Before considering an implementation ready for merge, the Verifier MUST confirm:

- [ ] Does `go test -v -race ./...` pass?
- [ ] Are there any unaddressed linting errors (`golangci-lint run ./...`)?
- [ ] Is the formatting completely clean (`gofmt -s -w .` and `gofumpt -extra -w .`)?
- [ ] Has the GitHub Actions CI pipeline returned a completely Green status for the latest commit?
- [ ] If CI is red, has the `ci_guardian` persona been engaged to fix it?
