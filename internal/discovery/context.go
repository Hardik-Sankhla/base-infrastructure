package discovery

import (
	"database/sql"
	"log/slog"

	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/platform"
	"github.com/base-infrastructure/platform/internal/runtime/events"
)

// Context provides a read-only view of platform runtime services to
// discovery stages.
type Context interface {
	Logger() *slog.Logger
	EventBus() events.Bus
	Config() *config.Config
	DB() *sql.DB
	Platform() platform.Platform
}

type defaultContext struct {
	logger   *slog.Logger
	eventBus events.Bus
	config   *config.Config
	db       *sql.DB
	plat     platform.Platform
}

func NewContext(logger *slog.Logger, bus events.Bus, cfg *config.Config, db *sql.DB, plat platform.Platform) Context {
	return &defaultContext{
		logger:   logger,
		eventBus: bus,
		config:   cfg,
		db:       db,
		plat:     plat,
	}
}

func (c *defaultContext) Logger() *slog.Logger        { return c.logger }
func (c *defaultContext) EventBus() events.Bus        { return c.eventBus }
func (c *defaultContext) Config() *config.Config      { return c.config }
func (c *defaultContext) DB() *sql.DB                 { return c.db }
func (c *defaultContext) Platform() platform.Platform { return c.plat }

// removed
