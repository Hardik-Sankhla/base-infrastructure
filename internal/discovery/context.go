package discovery

import (
	"database/sql"
	"log/slog"

	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/platform"
	"github.com/base-infrastructure/platform/internal/runtime"
)

// Context provides the execution environment and dependencies for discovery stages.
// It acts as a facade, hiding the complexity of the underlying systems.
type Context interface {
	Logger() *slog.Logger
	EventBus() runtime.EventBus
	Config() *config.Config
	DB() *sql.DB
	Platform() platform.Platform
	Cache() Cache
}

type defaultContext struct {
	logger   *slog.Logger
	bus      runtime.EventBus
	cfg      *config.Config
	db       *sql.DB
	platform platform.Platform
	cache    Cache
}

// NewContext creates a new discovery context with the provided dependencies.
func NewContext(logger *slog.Logger, bus runtime.EventBus, cfg *config.Config, db *sql.DB, p platform.Platform) Context {
	return &defaultContext{
		logger:   logger,
		bus:      bus,
		cfg:      cfg,
		db:       db,
		platform: p,
		cache:    NewCache(),
	}
}

func (c *defaultContext) Logger() *slog.Logger        { return c.logger }
func (c *defaultContext) EventBus() runtime.EventBus  { return c.bus }
func (c *defaultContext) Config() *config.Config      { return c.cfg }
func (c *defaultContext) DB() *sql.DB                 { return c.db }
func (c *defaultContext) Platform() platform.Platform { return c.platform }
func (c *defaultContext) Cache() Cache                { return c.cache }

// removed
