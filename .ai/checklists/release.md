# Release Checklist

Before cutting a new release, the Release Manager MUST verify:

- [ ] Is the default branch (`main`) completely stable and 100% Green on all required CI checks?
- [ ] Has a Repository Hygiene Audit been executed to catch trailing or generated files?
- [ ] Are the release notes generated according to `templates/release_notes.md`?
- [ ] Has the semantic version tag been properly bumped according to breaking changes vs features?
- [ ] Has the `website/` documentation been rebuilt and verified?
