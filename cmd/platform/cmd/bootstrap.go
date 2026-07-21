package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/discovery/builtin"
	"github.com/base-infrastructure/platform/internal/runtime/context"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Initialize environment from zero",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Starting platform bootstrap...")

		// Initialize Discovery Engine
		registry := discovery.NewRegistry()
		if err := builtin.RegisterCoreStages(registry); err != nil {
			slog.Error("Failed to register core discovery stages", "error", err)
			return
		}

		engine := discovery.NewDiscoveryEngine(registry, discovery.PipelineConfig{})

		// Setup Context
		pctx := context.NewPlatformContext(&config.Cfg, nil)

		// Run Discovery
		manifest, err := engine.Run(pctx)
		if err != nil {
			slog.Error("Discovery pipeline failed", "error", err)
			return
		}

		// Output result
		output, _ := json.MarshalIndent(manifest, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}
