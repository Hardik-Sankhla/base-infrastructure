# Architecture Principles

1. Interfaces Over Implementations.
2. The core domain must not depend on the OS abstraction.
3. Errors must propagate clearly and not panic (unless completely fatal at initialization).
4. CI/CD pipelines cannot be bypassed.
