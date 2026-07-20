package context

import (
	"database/sql"
	"log/slog"

	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/runtime/events"
	"github.com/base-infrastructure/platform/internal/runtime/fs"
	"github.com/base-infrastructure/platform/internal/runtime/http"
	"github.com/base-infrastructure/platform/internal/runtime/plugin"
	"github.com/base-infrastructure/platform/internal/runtime/tasks"
)

// PlatformContext is the unified context passed to all engines
type PlatformContext struct {
	Logger     *slog.Logger
	Config     *config.Config
	DB         *sql.DB
	EventBus   events.Bus
	TaskEngine tasks.Engine
	FS         fs.Manager
	Downloader http.Downloader
	Registry   plugin.Registry
}

func NewPlatformContext(cfg *config.Config, db *sql.DB) *PlatformContext {
	return &PlatformContext{
		Logger:     slog.Default(),
		Config:     cfg,
		DB:         db,
		EventBus:   events.NewBus(),
		TaskEngine: tasks.NewEngine(),
		FS:         fs.NewManager(cfg.System.DataDir),
		Downloader: http.NewDownloader(),
		Registry:   plugin.NewRegistry(),
	}
}
