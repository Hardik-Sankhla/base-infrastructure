package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/discovery/builtin"
	"github.com/base-infrastructure/platform/internal/logger"
	"github.com/base-infrastructure/platform/internal/platform/detector"
	"github.com/base-infrastructure/platform/internal/runtime/events"
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

		engine := discovery.NewEngine(registry)

		// Setup Context
		bus := events.NewEventBus()
		detector := detector.New()

		plat, err := detector.Detect()
		if err != nil {
			slog.Error("Failed to detect platform", "error", err)
			return
		}

		ctx := discovery.NewContext(logger.NewLogger(), bus, nil, nil, plat)

		// Run Discovery
		manifest, err := engine.Run(context.Background(), ctx)
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
