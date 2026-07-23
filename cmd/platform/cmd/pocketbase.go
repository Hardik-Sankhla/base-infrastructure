package cmd

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/base-infrastructure/platform/internal/infrastructure/pocketbase"
	"github.com/spf13/cobra"
)

var pocketbaseCmd = &cobra.Command{
	Use:   "pocketbase",
	Short: "Manage the integrated PocketBase instance",
	Long:  `Manage, start, stop, and inspect the PocketBase infrastructure component.`,
}

var pbStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start PocketBase server",
	Run: func(cmd *cobra.Command, args []string) {
		// Override args so pocketbase parses "serve"
		os.Args = []string{"platform-pb", "serve"}
		if err := pocketbase.Start(); err != nil {
			slog.Error("PocketBase failed to start", "error", err)
			os.Exit(1)
		}
	},
}

var pbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize PocketBase schema and data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Initializing PocketBase...")
		if _, err := pocketbase.Init(); err != nil {
			slog.Error("Failed to initialize pocketbase", "error", err)
			os.Exit(1)
		}

		// Run migrations up
		os.Args = []string{"platform-pb", "migrate", "up"}
		if err := pocketbase.App.Start(); err != nil {
			slog.Error("PocketBase migrations failed", "error", err)
			os.Exit(1)
		}

		fmt.Println("✅ PocketBase initialized successfully.")
	},
}

var pbStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the background PocketBase server",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Warn("Not implemented. Stop the process running 'platform pocketbase start' manually.")
	},
}

var pbStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the PocketBase backend",
	Run: func(cmd *cobra.Command, args []string) {
		client := http.Client{Timeout: 2 * time.Second}
		resp, err := client.Get("http://127.0.0.1:8090/api/health")
		if err != nil {
			fmt.Println("❌ PocketBase is NOT running or unreachable.")
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("✅ PocketBase is healthy and running.")
		} else {
			fmt.Printf("⚠️ PocketBase returned status: %d\n", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(pocketbaseCmd)
	pocketbaseCmd.AddCommand(pbInitCmd)
	pocketbaseCmd.AddCommand(pbStartCmd)
	pocketbaseCmd.AddCommand(pbStopCmd)
	pocketbaseCmd.AddCommand(pbStatusCmd)

	pbStartCmd.DisableFlagParsing = true
	pbInitCmd.DisableFlagParsing = true
}
