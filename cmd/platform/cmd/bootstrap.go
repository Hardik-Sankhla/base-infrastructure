package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Initialize environment from zero",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Starting platform bootstrap...")
		// Will call discovery, planner, executor here
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}
