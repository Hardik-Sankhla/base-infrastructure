package presentation

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
	"gopkg.in/yaml.v3"
)

// PrintOptions holds configurations for the presentation layer
type PrintOptions struct {
	Format    string   // "summary", "json", "yaml"
	Verbosity int      // 0 = standard, 1 = detailed, 2 = all
	Filters   []string // e.g. ["hardware", "network"]
	Output    string   // file path
}

// Result binds the DiscoveryManifest and Capabilities together
type Result struct {
	Manifest     *models.DiscoveryManifest `json:"manifest"`
	Capabilities []models.Capability       `json:"capabilities"`
}

// Print formats and optionally saves the discovery results based on options
func Print(res Result, opts PrintOptions) error {
	// 1. Filter the manifest if filters are provided
	filteredManifest := filterManifest(res.Manifest, opts.Filters)

	filteredRes := Result{
		Manifest:     filteredManifest,
		Capabilities: res.Capabilities,
	}

	// 2. Generate the output string
	var output string
	var err error

	switch opts.Format {
	case "json":
		output, err = formatJSON(filteredRes)
	case "yaml":
		output, err = formatYAML(filteredRes)
	default: // "summary"
		output = formatSummary(filteredRes, opts.Verbosity)
	}

	if err != nil {
		return fmt.Errorf("formatting failed: %w", err)
	}

	// 3. Print to console
	fmt.Println(output)

	// 4. Save to file if output path is provided
	if opts.Output != "" {
		fileOutput := output
		if opts.Format == "summary" {
			fileOutput, _ = formatJSON(filteredRes) // Always save structured data
		}

		if err := os.WriteFile(opts.Output, []byte(fileOutput), 0o644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		if opts.Format == "summary" {
			fmt.Printf("\nGenerated %s\n", opts.Output)
		} else {
			fmt.Printf("\nFull report saved to: %s\n", opts.Output)
		}
	}

	return nil
}

func formatJSON(res Result) (string, error) {
	bytes, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func formatYAML(res Result) (string, error) {
	bytes, err := yaml.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PrintHealth(h discovery.RepositoryHealth) {
	passFail := func(ok bool) string {
		if ok {
			return "PASS"
		}
		return "FAIL"
	}

	fmt.Println("\nRepository Doctor")
	fmt.Println("────────────────────────────────────────")
	fmt.Printf("%-24s %s\n", "Repository Brain", passFail(h.RepositoryBrain))
	fmt.Printf("%-24s %s\n", "Architecture", passFail(h.Architecture))
	fmt.Printf("%-24s %s\n", "Documentation", passFail(h.Documentation))
	fmt.Printf("%-24s %s\n", "GitHub Actions", h.CIStatus)
	fmt.Printf("%-24s %s\n", "Coverage", h.Coverage)
	fmt.Printf("%-24s %d items\n", "Tech Debt", h.TechDebtCount)
	
	releaseStr := "YES"
	if !h.ReleaseReady {
		releaseStr = "NO"
	}
	fmt.Printf("%-24s %s\n\n", "Release Ready", releaseStr)
}
