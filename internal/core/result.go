package core

import (
	"sync"
	"time"
)

// Result is the aggregated output of the entire discovery pipeline.
type Result struct {
	Stages map[string]*StageResult `json:"stages"`
	Errors map[string]string       `json:"errors,omitempty"`

	// Summary statistics
	TotalStages      int           `json:"total_stages"`
	SuccessfulStages int           `json:"successful_stages"`
	FailedStages     int           `json:"failed_stages"`
	SkippedStages    int           `json:"skipped_stages"`
	StartTime        time.Time     `json:"start_time"`
	EndTime          time.Time     `json:"end_time"`
	Duration         time.Duration `json:"duration"`
	Success          bool          `json:"success"`
}

// GetStageData retrieves the typed DiscoveryArtifact from a stage result by name.
// Returns nil and false if the stage is not found or produced no result.
func (r *Result) GetStageData(name string) (DiscoveryArtifact, bool) {
	if r == nil || r.Stages == nil {
		return nil, false
	}
	sr, ok := r.Stages[name]
	if !ok || sr == nil {
		return nil, false
	}
	return sr.Artifact, sr.Artifact != nil
}

// StageNames returns the names of all stages that produced results.
func (r *Result) StageNames() []string {
	if r == nil || r.Stages == nil {
		return nil
	}
	names := make([]string, 0, len(r.Stages))
	for name := range r.Stages {
		names = append(names, name)
	}
	return names
}

// ResultBuilder accumulates stage results and produces an immutable Result.
// It is safe for concurrent use.
type ResultBuilder struct {
	stages    map[string]*StageResult
	errors    map[string]string
	startTime time.Time
	mu        sync.Mutex

	total      int
	successful int
	failed     int
	skipped    int
}

// NewResultBuilder creates a new ResultBuilder and records the start time.
func NewResultBuilder() *ResultBuilder {
	return &ResultBuilder{
		stages:    make(map[string]*StageResult),
		errors:    make(map[string]string),
		startTime: time.Now().UTC(),
	}
}

// SetTotalStages sets the expected total number of stages.
func (b *ResultBuilder) SetTotalStages(total int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.total = total
}

// AddStageResult records a stage result and updates summary statistics.
func (b *ResultBuilder) AddStageResult(sr *StageResult) {
	if sr == nil {
		return
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	b.stages[sr.StageName] = sr
	if sr.Error != "" {
		b.errors[sr.StageName] = sr.Error
	}

	switch sr.Status {
	case StatusSuccess:
		b.successful++
	case StatusFailed:
		b.failed++
	case StatusSkipped:
		b.skipped++
	}
}

// AddStageError records a fatal error for a stage that failed before producing a result.
func (b *ResultBuilder) AddStageError(stageName string, err error) {
	if err == nil {
		return
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	b.errors[stageName] = err.Error()
	b.failed++
}

// Build produces the final immutable Result. After calling Build,
// the builder should not be reused.
func (b *ResultBuilder) Build() *Result {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now().UTC()
	return &Result{
		Stages:           b.stages,
		Errors:           b.errors,
		TotalStages:      b.total,
		SuccessfulStages: b.successful,
		FailedStages:     b.failed,
		SkippedStages:    b.skipped,
		StartTime:        b.startTime,
		EndTime:          now,
		Duration:         now.Sub(b.startTime),
		Success:          b.failed == 0,
	}
}
