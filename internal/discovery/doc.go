// Package discovery implements the Discovery Engine as a composable stage
// pipeline. Each stage probes one facet of the host environment (hardware, OS,
// filesystem, network, etc.) and produces an immutable core.StageResult. The
// pipeline aggregates all stage results into a single Result that downstream
// engines (Planner, Executor) consume.
//
// Architecture:
//
//	Registry   ──registers──▶  core.Stage (interface)
//	    │                         │
//	    ▼                         ▼
//	Pipeline   ──executes──▶  core.StageResult
//	    │
//	    ▼
//	Result     ──consumed by──▶  Planner
//
// Stages are independent and composable. They communicate only through
// immutable domain models and never mutate previous stages' outputs.
// The pipeline supports fail-fast and continue-on-error execution modes.
package discovery
