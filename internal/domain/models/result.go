package models

import "time"

// Result is the unified response from all execution engines
type Result struct {
	Success   bool              `json:"success"`
	Warnings  []error           `json:"warnings"`
	Errors    []error           `json:"errors"`
	Duration  time.Duration     `json:"duration"`
	Rollback  bool              `json:"rollback"`
	Artifacts map[string]string `json:"artifacts"`
	Metadata  map[string]string `json:"metadata"`
}
