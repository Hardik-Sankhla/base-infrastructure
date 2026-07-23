package pocketbase

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/base-infrastructure/platform/internal/infrastructure/pocketbase/migrations"
)

var App *pocketbase.PocketBase

// Init initializes and configures the PocketBase application
func Init() (*pocketbase.PocketBase, error) {
	slog.Info("Initializing PocketBase")

	// Since PBDataDir isn't in Config yet, we'll read it from env directly or default
	dataDir := os.Getenv("PLATFORM_PB_DATA_DIR")
	if dataDir == "" {
		dataDir = ".pb_data" // Default
	}

	// Ensure absolute path
	absDataDir, err := filepath.Abs(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to determine absolute path for PocketBase data dir: %w", err)
	}

	// Make sure the directory exists
	if err := os.MkdirAll(absDataDir, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create PocketBase data directory: %w", err)
	}

	App = pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: absDataDir,
	})

	migratecmd.MustRegister(App, App.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	// Add basic events or middleware if needed
	App.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		slog.Info("PocketBase server started", "url", "http://127.0.0.1:8090/_/")
		return nil
	})

	return App, nil
}

// Start runs the PocketBase server
func Start() error {
	if App == nil {
		if _, err := Init(); err != nil {
			return err
		}
	}

	// Start PocketBase (this will block)
	slog.Info("Starting PocketBase...")
	if err := App.Start(); err != nil {
		return fmt.Errorf("pocketbase failed to start: %w", err)
	}

	return nil
}
