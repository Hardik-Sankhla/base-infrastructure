# Implementation Checklist

Before completing an implementation sprint, the Builder MUST verify:

- [ ] Does the implementation strictly match the approved `design_review.md`?
- [ ] Are all new interfaces and public methods documented?
- [ ] Were exhaustive table-driven unit tests written?
- [ ] Has the local verification suite passed (`go fmt`, `go vet`, `go test -race ./...`)?
- [ ] Have all transient/temporary files been removed from the working directory?
- [ ] Is the `implementation_report.md` artifact generated for the user?
- [ ] Has `docs/` been synchronized with these changes?
