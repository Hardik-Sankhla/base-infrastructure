# Documentation Protocol

**Documentation is infrastructure.**

## 1. Concurrent Updates
Documentation must be updated concurrently with the code that it describes. 
- Never treat documentation as an afterthought.
- It is enforced via the `implementation.md` checklist.

## 2. When to Update
If a PR changes:
- Any interface in `internal/domain/contracts/` → update the Architecture rules.
- Any platform provider → update provider capabilities documentation.
- Any CLI command behavior → update CLI reference docs.
- Any architectural decision → add an ADR to `docs/adr/`.
- Any known bug → update the known issues tracker.

## 3. Website Synchronization
The `website/` directory contains a VitePress build of the project documentation. Ensure `website/` stays synchronized with `docs/` and that the `docs.yml` GitHub action remains green after updates.
