package discovery

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// RepositoryHealth represents the overall health of the engineering system.
type RepositoryHealth struct {
	RepositoryBrain bool
	Architecture    bool
	Documentation   bool
	CIStatus        string
	Coverage        string
	TechDebtCount   int
	ReleaseReady    bool
}

// AnalyzeHealth performs the repository doctor checks.
func AnalyzeHealth(ctx context.Context, repoRoot string) (RepositoryHealth, error) {
	var health RepositoryHealth
	health.ReleaseReady = true // Assume true, fail below if conditions aren't met

	// 1. Check Repository Brain
	if _, err := os.Stat(filepath.Join(repoRoot, ".ai", "SESSION_PROTOCOL.md")); err == nil {
		health.RepositoryBrain = true
	} else {
		health.ReleaseReady = false
	}

	// 2. Check Architecture
	if _, err := os.Stat(filepath.Join(repoRoot, "docs", "architecture", "DECISION_LOG.md")); err == nil {
		health.Architecture = true
	} else {
		health.ReleaseReady = false
	}

	// 3. Check Documentation (Simplistic check for README and CONTRIBUTING)
	hasReadme := false
	hasContributing := false
	if _, err := os.Stat(filepath.Join(repoRoot, "README.md")); err == nil {
		hasReadme = true
	}
	if _, err := os.Stat(filepath.Join(repoRoot, "CONTRIBUTING.md")); err == nil {
		hasContributing = true
	}
	health.Documentation = hasReadme && hasContributing

	// 4. Extract CI Status from REPOSITORY_STATE.md
	health.CIStatus = "UNKNOWN"
	statePath := filepath.Join(repoRoot, "docs", "ai", "REPOSITORY_STATE.md")
	if b, err := os.ReadFile(statePath); err == nil {
		content := string(b)
		if strings.Contains(content, "ci_status: passing") || strings.Contains(content, "ci_status: green") {
			health.CIStatus = "PASS"
		} else if strings.Contains(content, "ci_status: failing") || strings.Contains(content, "ci_status: red") {
			health.CIStatus = "FAIL"
			health.ReleaseReady = false
		}
	}

	// 5. Calculate Coverage via go test
	cmd := exec.CommandContext(ctx, "go", "test", "-cover", "./...")
	cmd.Dir = repoRoot
	var out bytes.Buffer
	cmd.Stdout = &out
	// Ignore errors since tests might just have no test files in some packages, or maybe test fails
	_ = cmd.Run()

	var coverage string
	lines := strings.Split(out.String(), "\n")
	var totalCov float64
	var pkgCount int
	reCov := regexp.MustCompile(`coverage:\s+([0-9.]+)%`)
	for _, line := range lines {
		if matches := reCov.FindStringSubmatch(line); len(matches) > 1 {
			var cov float64
			fmt.Sscanf(matches[1], "%f", &cov)
			totalCov += cov
			pkgCount++
		}
	}
	if pkgCount > 0 {
		avg := totalCov / float64(pkgCount)
		coverage = fmt.Sprintf("%.1f%%", avg)
		if avg < 50.0 {
			health.ReleaseReady = false // Enforce arbitrary 50% minimum for release readiness in v3
		}
	} else {
		coverage = "N/A"
	}
	health.Coverage = coverage

	// 6. Calculate Tech Debt (Count TODO/FIXME in .go and .md files)
	var debtCount int
	err := filepath.Walk(repoRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			name := info.Name()
			if name == ".git" || name == "vendor" || name == "bin" {
				return filepath.SkipDir
			}
			return nil
		}

		ext := filepath.Ext(path)
		if ext == ".go" || ext == ".md" {
			b, err := os.ReadFile(path)
			if err == nil {
				content := string(b)
				debtCount += strings.Count(content, "TODO")
				debtCount += strings.Count(content, "FIXME")
			}
		}
		return nil
	})
	if err == nil {
		health.TechDebtCount = debtCount
	}

	return health, nil
}
