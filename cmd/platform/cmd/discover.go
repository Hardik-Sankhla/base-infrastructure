package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/base-infrastructure/platform/internal/capabilities"
	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/discovery/builtin"
	"github.com/base-infrastructure/platform/internal/logger"
	"github.com/base-infrastructure/platform/internal/presentation"
	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/spf13/cobra"
)

var (
	formatOpt string
	outputOpt string
	isJSON    bool
	isYAML    bool
	hwOnly    bool
	netOnly   bool
	fsOnly    bool
	osOnly    bool
	swOnly    bool
	verbosity int
)

var discoverCmd = &cobra.Command{
	Use:     "discover",
	Aliases: []string{"bootstrap"}, // keep bootstrap as an alias for backwards compatibility
	Short:   "Discover host environment capabilities",
	Run: func(cmd *cobra.Command, args []string) {
		// Determine Format
		format := "summary"
		if isJSON || formatOpt == "json" {
			format = "json"
		} else if isYAML || formatOpt == "yaml" {
			format = "yaml"
		}

		// Adjust Logging based on verbosity and format
		if verbosity == 0 && format == "summary" {
			// Suppress raw JSON logs for a clean CLI experience
			logger.Init("error", "text")
		} else if verbosity > 0 {
			// Elevate logging
			if verbosity >= 2 {
				logger.Init("debug", "json")
			} else {
				logger.Init("info", "json")
			}
		}

		slog.Debug("Starting platform discovery...")

		// Initialize Discovery Engine
		registry := discovery.NewRegistry()
		if err := builtin.RegisterCoreStages(registry); err != nil {
			slog.Error("Failed to register core discovery stages", "error", err)
			os.Exit(1)
		}

		engine := discovery.NewDiscoveryEngine(registry, discovery.PipelineConfig{})

		// Attach ProgressHook for clean CLI UX
		if verbosity == 0 && format == "summary" {
			engine.AddHook(NewProgressHook())
		}

		// Setup Context
		pctx := runtime.NewPlatformContext(&config.Cfg, nil)

		// Run Discovery
		manifest, err := engine.Run(pctx)
		if err != nil {
			slog.Error("Discovery pipeline failed", "error", err)
			os.Exit(1)
		}

		// Build Capabilities
		builder := capabilities.NewBuilder(manifest)
		caps := builder.Build()

		// Prepare Result
		res := presentation.Result{
			Manifest:     manifest,
			Capabilities: caps,
		}

		// Determine Filters
		var filters []string
		if hwOnly {
			filters = append(filters, "Hardware")
		}
		if netOnly {
			filters = append(filters, "Network")
		}
		if fsOnly {
			filters = append(filters, "Filesystem")
		}
		if osOnly {
			filters = append(filters, "OS")
		}
		if swOnly {
			filters = append(filters, "Software")
		}

		// Configure Options
		opts := presentation.PrintOptions{
			Format:    format,
			Verbosity: verbosity,
			Filters:   filters,
			Output:    outputOpt,
		}

		// Auto-save behavior if format is summary and no output provided
		if opts.Format == "summary" && opts.Output == "" {
			// Auto save to a reports dir if possible or current dir
			// A production system might save to `~/.base-infrastructure/reports/`
			// We'll just save it to reports/ if we can, or current dir.
			opts.Output = "discovery-" + time.Now().Format("2006-01-02T15-04-05") + ".json"
		}

		if err := presentation.Print(res, opts); err != nil {
			slog.Error("Failed to format output", "error", err)
			os.Exit(1)
		}

		if format == "summary" {
			fmt.Printf("\nDiscovery completed successfully.\n")
		}
		os.Exit(0)
	},
}

func init() {
	discoverCmd.Flags().StringVarP(&formatOpt, "format", "f", "summary", "Output format (summary, json, yaml)")
	discoverCmd.Flags().BoolVar(&isJSON, "json", false, "Alias for --format=json")
	discoverCmd.Flags().BoolVar(&isYAML, "yaml", false, "Alias for --format=yaml")
	discoverCmd.Flags().StringVarP(&outputOpt, "output", "o", "", "Save full report to file")

	discoverCmd.Flags().BoolVar(&hwOnly, "hardware", false, "Show only hardware capabilities")
	discoverCmd.Flags().BoolVar(&netOnly, "network", false, "Show only network capabilities")
	discoverCmd.Flags().BoolVar(&fsOnly, "filesystem", false, "Show only filesystem capabilities")
	discoverCmd.Flags().BoolVar(&osOnly, "os", false, "Show only OS capabilities")
	discoverCmd.Flags().BoolVar(&swOnly, "software", false, "Show only software capabilities")

	discoverCmd.Flags().CountVarP(&verbosity, "verbose", "v", "Verbosity level")

	rootCmd.AddCommand(discoverCmd)
}
