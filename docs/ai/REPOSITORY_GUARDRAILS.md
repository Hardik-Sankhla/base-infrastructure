# Repository Guardrails

The following rules dictate the behavior of all AI Agents operating within the `base-infrastructure` repository. These rules are non-negotiable.

## RULE 1
**Never mark a task complete while GitHub Actions is failing.**
The repository is the source of truth. Not the agent. Not the report. Not the local build. A task is complete only when all workflows are green.

## RULE 2
**Never claim production readiness without successful CI.**
Local static analysis or compilation is insufficient to claim production readiness.

## RULE 3
**Every implementation requires runtime verification.**
You must execute the software and observe its outputs before considering an implementation complete.

## RULE 4
**Every audit requires evidence.**
Reports without evidence are observations, not verification. If you state a feature works, you must provide the exact command executed, the expected behavior, the actual behavior, the exit code, and the stdout/stderr.

## RULE 5
**Documentation must match implementation.**
Never hallucinate features that do not exist. Documentation must accurately describe the current state of the codebase.

## RULE 6
**Never suppress failing checks.**
Do not bypass or disable linting, formatting, or test warnings without explicit, justified engineering reason.

## RULE 7
**Always fix the root cause.**
Do not implement temporary workarounds for CI/CD or compilation failures. Investigate and resolve the underlying issue.

## RULE 8
**Strict Evidence Policy**
Every claim requires evidence. If you say "The issue is fixed", you must provide:
- Commit SHA
- Workflow name
- Job name
- GitHub Actions result
- Exit code
- Logs reviewed
- Evidence that the workflow completed successfully

If you cannot provide evidence, you must state: "I have not verified this." Never infer success. Never assume success. Never fabricate verification. Evidence always takes precedence over confidence.
