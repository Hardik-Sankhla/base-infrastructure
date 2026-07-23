# Engineering Session Protocol

Every engineering session MUST follow this sequence based on Engineering OS v2.0.

## Phase 1 - Boot
Read:
- `.ai/START_HERE.md`
- `.ai/PROJECT_IDENTITY.md`
- `.ai/ENGINEERING_LIFECYCLE.md`
- `.ai/ARCHITECTURE_RULES.md`
- `.ai/MEMORY_PROTOCOL.md`

Adopt the appropriate Persona (`.ai/personas/*.md`) for the current phase.
Do not modify code yet.

## Phase 2 - Planning (Planner / Architect Persona)
- Reference `.ai/checklists/design_review.md`.
- Produce the `.ai/templates/design_review.md` or `.ai/templates/sprint_report.md`.
- Wait for explicit user approval before execution.

## Phase 3 - Implementation (Builder Persona)
- Reference `.ai/checklists/implementation.md`.
- Implement only the approved scope.
- Do not introduce unrelated changes.
- Stop after implementation.

## Phase 4 - Verification (Verifier / CI Guardian Persona)
- Reference `.ai/checklists/verification.md` and `.ai/QUALITY_GATES.md`.
- Ensure all local linting, testing, and formatting passes.
- Ensure GitHub Actions pipelines are 100% Green.

## Phase 5 - Review & Merge (Reviewer / Release Manager Persona)
- Reference `.ai/checklists/merge.md` or `.ai/checklists/release.md`.
- Verify architectural boundaries and documentation sync.
- Generate `.ai/templates/release_notes.md` if releasing.

## Phase 6 - Repository Memory
- Delegate state storage and historical updates according to `.ai/MEMORY_PROTOCOL.md` and `.ai/DOCUMENTATION_PROTOCOL.md`.
- Produce final reports using `.ai/templates/implementation_report.md` or `.ai/templates/ci_report.md`.
