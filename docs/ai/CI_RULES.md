# CI_RULES.md
# CI/CD Pipeline Rules

> Source-verified from `.github/workflows/ci.yml`.
> **Last verified**: 2026-07-21

---

## Active CI Pipelines

### `CI` — `.github/workflows/ci.yml`
Triggers on push to: `main`, `fix/*`, `feature/*`, `chore/*`
Triggers on PR to: `main`
Runner: `ubuntu-latest`

**Steps in order** (source: `ci.yml:12-75`):
1. Checkout (`actions/checkout@v4`, `fetch-depth: 0`)
2. Setup Go 1.22 (`actions/setup-go@v5`)
3. `go mod download`
4. Install tools: `gofumpt` + `golangci-lint v1.56.2`
5. **Verify gofmt** — fails if any file is not formatted
6. **Verify gofumpt** — fails if any file is not formatted with extra rules
7. `go build -v ./...` (pre-cache for linter)
8. `golangci-lint run ./...`
9. `go test -v -race ./...`
10. `go test -v -coverprofile=coverage.out ./...`
11. Cross-build Linux amd64
12. Cross-build Windows amd64
13. Cross-build macOS arm64

### `Deploy Documentation` — `.github/workflows/docs.yml`
Triggers on push to `main` when `docs/**`, `website/**`, or `package.json` changes.
Triggers on `workflow_dispatch`.

**Steps**:
1. Checkout with full history
2. Setup Node.js 20
3. `npm install`
4. `npm run docs:build` (VitePress → `website/.vitepress/dist/`)
5. Upload pages artifact
6. Deploy to GitHub Pages

---

## CI Rules for Contributors

### CR1: Never Push Code that Breaks Any CI Step
All 13 CI steps are blocking. A failing lint, format, or build is a blocker for merge.

### CR2: Format Before Every Commit
```bash
gofmt -s -w .
gofumpt -extra -w .
```
The CI will fail with a diff if these are not applied.

### CR3: Documentation Changes Deploy Automatically
Any push to `main` touching `docs/` or `website/` triggers a new GitHub Pages deployment. Do not push broken Markdown that will cause VitePress build errors.

### CR4: Cross-Platform Builds Must Pass
The CI verifies builds for `linux/amd64`, `windows/amd64`, and `darwin/arm64`. Platform-specific code that compiles on one OS but not another will surface here.

---

## Adding a New Workflow

Before adding a new GitHub Actions workflow:
1. Confirm it does not conflict with the `pages` concurrency group
2. Use `actions/checkout@v4`, `actions/setup-go@v5` — keep action versions consistent
3. Document the new workflow in this file
