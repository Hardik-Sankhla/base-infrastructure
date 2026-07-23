package cmd

import (
	"fmt"
	"os"

	"github.com/base-infrastructure/platform/internal/bootstrap"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Initialize the complete development environment",
	Long:  `Bootstrap sets up the infrastructure platform for a new developer. It initializes directories, sets up configuration, runs migrations, and prepares the backend for development.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting Platform Bootstrap Process...")

		// Run bootstrap
		if err := bootstrap.Current.BootstrapEnvironment(); err != nil {
			fmt.Println("❌ Failed to initialize environment:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Environment configuration loaded.")

		fmt.Println("")
		fmt.Println("🎉 Bootstrap complete! Your development environment is ready.")
		fmt.Println("To start the backend infrastructure, run:")
		fmt.Println("    platform pocketbase start")
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}
