# Engineering Session Protocol

Every engineering session MUST follow this sequence.

## Phase 1 - Boot
Read:
- `.ai/START_HERE.md`
- `.ai/PROJECT_IDENTITY.md`
- `.ai/GUARDRAILS.md`
- `docs/ai/REPOSITORY_STATE.md`
- `docs/ai/NEXT_TASK.md`
- `docs/ai/TECH_DEBT.md`
- `docs/ai/REPOSITORY_MAP.md`

Summarize your understanding.
Do not modify code yet.

---

## Phase 2 - Planning
Determine:
- Current milestone
- Current blockers
- Files likely to change
- Risks
- Verification strategy

---

## Phase 3 - Implementation
Implement only the approved scope.
Do not introduce unrelated changes.

---

## Phase 4 - Verification
Run:
`make verify`

Verify documentation.
Verify CLI.
Verify CI requirements.

---

## Phase 5 - Evidence
Collect evidence.
Do not infer success.

---

## Phase 6 - Repository Memory
Update:
- `docs/ai/REPOSITORY_STATE.md`
- `docs/ai/NEXT_TASK.md`
- `docs/ai/TECH_DEBT.md`
- `CHANGELOG.md`

---

## Phase 7 - Final Report
Produce:
- What changed
- Evidence
- Known issues
- Recommended next task
