package bootstrap

import (
	"context"
	"os"
	"os/exec"

	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/core"
	"github.com/base-infrastructure/platform/internal/logger"
	"github.com/base-infrastructure/platform/internal/services/pocketbase"
)

// App is the global application state and dependency orchestrator.
type App struct {
	ConfigLoaded bool
}

// Global bootstrap app
var Current *App

func init() {
	Current = &App{}
}

// Initialize loads configuration and logger.
func (a *App) Initialize(cfgFile string) error {
	if err := config.Load(cfgFile); err != nil {
		return err
	}
	logger.Init(config.Cfg.System.LogLevel, config.Cfg.Logging.Format)
	a.ConfigLoaded = true
	return nil
}

// BootstrapEnvironment initializes work directories and runs migrations.
func (a *App) BootstrapEnvironment() error {
	if err := os.MkdirAll(config.Cfg.System.DataDir, 0o755); err != nil {
		return err
	}

	cmdPath, _ := os.Executable()
	initCmd := exec.Command(cmdPath, "pocketbase", "init")
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr
	return initCmd.Run()
}

// StartDatabase starts the PocketBase background server.
func (a *App) StartDatabase() error {
	return pocketbase.Start()
}

// InitDatabase initializes schemas and data for PocketBase.
func (a *App) InitDatabase() error {
	if _, err := pocketbase.Init(); err != nil {
		return err
	}
	os.Args = []string{"platform-pb", "migrate", "up"}
	return pocketbase.App.Start()
}

// RunDiscovery runs the discovery engine and returns the health result.
func (a *App) RunDiscovery(ctx context.Context, repoRoot string) (core.RepositoryHealth, error) {
	// This would be replaced with actual pipeline execution later
	return core.AnalyzeHealth(ctx, repoRoot)
}
