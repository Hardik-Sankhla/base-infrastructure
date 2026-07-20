package context

import (
	gocontext "context"
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

	// goCtx is the cancellable Go context for this platform run.
	goCtx gocontext.Context
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
		goCtx:      gocontext.Background(),
	}
}

// GoContext returns the cancellable Go context for this platform run.
// If no context was set, it returns context.Background().
func (p *PlatformContext) GoContext() gocontext.Context {
	if p.goCtx != nil {
		return p.goCtx
	}
	return gocontext.Background()
}

// WithGoContext returns a shallow copy of the PlatformContext with the
// given Go context set. This is useful for injecting cancellation.
func (p *PlatformContext) WithGoContext(ctx gocontext.Context) *PlatformContext {
	cp := *p
	cp.goCtx = ctx
	return &cp
}
