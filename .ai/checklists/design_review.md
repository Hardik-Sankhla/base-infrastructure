# Design Review Checklist

Before approving a design, the Architect MUST verify:

- [ ] Does the design strictly preserve the Architecture Freeze boundaries?
- [ ] Are all dependencies pointing inwards (`CLI -> Bootstrap -> Planners/Engines -> Platform -> Core`)?
- [ ] Are Platform-specific operations strictly confined to `internal/platform/`?
- [ ] Have all alternatives been thoroughly considered and documented?
- [ ] Is the design small enough to be executed in a single sprint?
- [ ] Has an ADR been generated if this changes core behavior?
