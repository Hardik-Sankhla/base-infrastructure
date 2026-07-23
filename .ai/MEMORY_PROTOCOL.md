# Memory Protocol

This document outlines the strategy for reading and writing state tracking files, ensuring historical continuity across AI sessions.

## 1. Transient State
- Current sprint progress and temporary tasks are stored locally in session artifacts or a temporary `task.md`.
- Transient state should not pollute the global Git history unless actively collaborating on a long-running PR branch.

## 2. Persistent Memory
- Architectural decisions, complex design nuances, and known issues are codified into `docs/adr/` (Architecture Decision Records).
- The `docs/` folder is the permanent source of truth for the system's memory.

## 3. Pre-Session Loading
Agents should always reference the Root Documents in `.ai/` and relevant ADRs before concluding a research phase. Do not blindly overwrite historical logic without cross-referencing previous design documents.
