package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/presentation"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Checks the health of the repository and engineering system",
	Long:  `Validates the Repository Brain, CI Status, Architecture docs, Tech Debt, and Coverage.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, _ := os.Getwd()
		
		health, err := discovery.AnalyzeHealth(context.Background(), repoRoot)
		if err != nil {
			fmt.Printf("Failed to run doctor: %v\n", err)
			os.Exit(1)
		}

		presentation.PrintHealth(health)

		if !health.ReleaseReady {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
