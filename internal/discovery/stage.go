package discovery

import (
	"context"
	"time"
)

// StageStatus represents the execution state of a stage.
type StageStatus string

const (
	StatusPending StageStatus = "Pending"
	StatusRunning StageStatus = "Running"
	StatusSuccess StageStatus = "Success"
	StatusFailed  StageStatus = "Failed"
	StatusSkipped StageStatus = "Skipped"
)

// Stage is the contract every discovery stage must implement.
type Stage interface {
	// Metadata
	Name() string
	Version() string
	Description() string
	Priority() int
	DependsOn() []string
	Timeout() time.Duration

	// Lifecycle
	Initialize(dctx Context) error
	Run(ctx context.Context, dctx Context) (DiscoveryArtifact, error)
	Validate(artifact DiscoveryArtifact) error
	Cleanup(ctx context.Context) error
}

// StageResult is the immutable output of a single Stage execution.
type StageResult struct {
	StageName string            `json:"stage_name"`
	Status    StageStatus       `json:"status"`
	Artifact  DiscoveryArtifact `json:"artifact,omitempty"`
	Error     string            `json:"error,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Duration  time.Duration     `json:"duration"`
	Timestamp time.Time         `json:"timestamp"`
}
