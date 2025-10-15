# Agent Collaboration Guide

## 1. Purpose & Scope
- Coordinate automation and human contributors across the job-search monorepo (`core`, `ui-next`, tooling).
- Maintain a clean Domain-Driven Design separation (interfaces → services → domain/infrastructure).
- Complement CLAUDE.md (collaboration etiquette) and README.md (project setup). Review those alongside this guide.

## 2. Roles & Responsibilities
- **API Agent**: Owns `core/src` API modules, enforces DDD boundaries, keeps Express routes aligned with services, reviews Prisma-driven data access.
- **Search Agent**: Maintains scrapers, scoring, embeddings, job ingestion utilities. Ensures asynchronous embedding generation and scoring logic remain intact.
- **UI Agent**: Manages `ui-next`, mirrors API contract changes, coordinates schema updates with the API Agent.
- **Ops/Infra Agent**: Handles Docker/dev scripts, `.sops.yaml` secrets flow, runtime database provisioning, and long-lived migrations when required.
- **Review/QA Agent**: Final verification, regression testing, release notes. Blocks merges until quality gates pass.
- **Go Agent**: Leads translation of legacy tooling (e.g., `db2Pdl.js`, ORM helpers) into Go services and CLIs while preserving Domain-Driven Design boundaries and coordinating downstream language generators (PHP → Go → Kotlin → Java → C# → C++).

## 3. Operating Procedures
- **Change Coordination**: AI agents must not introduce high-impact structural or tooling changes (e.g., refactoring shared test harnesses, modifying global mocks, altering build/test pipelines) without explicit user direction. When unforeseen blockers arise, stop and request guidance before proceeding.
- **Explicit Scope**: AI agents must not introduce new files, templates, or behavioral changes beyond the user’s request. When an adjustment seems necessary, pause and confirm with the user before proceeding.
- **Task Intake**: Reference issue tracker for priorities; claim tasks before starting. Surface blockers early.
- **Branching & Commits**: Default to local work without creating branches or commits unless the user explicitly instructs otherwise.
- **Coding Standards**: TypeScript ES2020 modules. Use type-only imports/exports when re-exporting types. Avoid inline `//` comments and trailing comments. Prefer self-explanatory code; add concise docstrings only when behavior is non-obvious. Avoid one-letter identifiers across languages; one-letter variables are prohibited — use descriptive names in all cases. Every function must return a variable named `result`. Exceptions: one-liner, if a function returns a variable called "x", it means the name of the variable is result. Never assign a literal/constant to `result` just to return it—return literals directly without introducing a temporary. Limit functions to a maximum of 35 lines and no more than three parameters; split work into helpers when necessary. Go tests must use descriptive parameter names (e.g., `testState *testing.T`), never just `t`.
- **Comment Discipline**: Do not add inline `//` comments to any language.
- **Tooling**: Use `pnpm` for all Node workflows; avoid `npm` and `npx`.
- **Binary builds**: We ship Linux-only binaries. Container images and `/usr/bin/pdl` installs are the target; skip Windows-specific build steps.
- **DDD Enforcement**: Interface handlers delegate to services; services orchestrate repositories; domain/infrastructure remain persistence-focused. No repository usage directly in handlers.
- **Testing & Builds**: Run `pnpm run build:core` before handing off. Use `pnpm run test:core` where regression risk exists. For UI, run `pnpm --filter ui-next run build` if relevant changes occur.
- **Release Coordination**: Notify QA Agent before tagging releases; ensure API/UI compatibility notes are documented.

## 4. Coordination & Handoffs
- Sync with affected agents whenever schemas, API contracts, or scraping logic change.
- Document cross-cutting updates in PR descriptions or handoff notes.
- Escalate conflicts to the Ops Agent or project owner; defer to CLAUDE.md rules for dispute resolution.

## 5. Environments & Secrets
- Development uses local Postgres via Docker compose (`pnpm run db:start`). `.env.local` is managed through SOPS; never commit decrypted secrets.
- Respect sandbox limits: file writes confined to workspace, network access requires approval, destructive commands need explicit user consent.
- When new secrets or services are required, coordinate through Ops Agent and update this guide once approved.

## 6. Quality Gates & Tooling
- Mandatory checks: `pnpm run build:core`, relevant unit tests, linting (if enabled ensure `no-inline-comments` rule is active).
- Maintain logging for long-running jobs (scoring, scraping) and fail gracefully with informative errors.
- Keep dependencies stable; avoid introducing network-installed packages unless authorized.

## 7. Troubleshooting Playbooks
- **Prisma/DB issues**: Prefer runtime migrations via schema scripts; avoid standalone migrations unless directed.
- **Embedding Failures**: Verify Xenova pipeline caching, ensure resume/job text cleanup (`cleanMarkdown`, `trimToLength`) works. Resume embeddings run asynchronously after upload.
- **Scraper Breakage**: Review search agent utilities, respect existing scoring/scraping behavior; changes require explicit instruction.
- **Secrets Errors**: Regenerate `.env.local` via `pnpm run env:decrypt` (SOPS) and confirm permissions.

## 8. Change Log & Ownership
- Maintain a dated log of modifications to this document (add entries below this section when edits are made).
- Current maintainers: API Agent lead, Search Agent lead, Ops Agent lead.
- Submit proposed updates via PR or direct request to maintainers when roles/processes change.

- 2025-10-09: Documented pnpm-only tooling requirement and refreshed Vitest guidance (Codex).
- 2025-10-09: Added guardrail requiring user approval before high-impact refactors (Codex).
- 2025-10-10: Clarified constant returns must not assign to `result` first; Go tests must use descriptive `testing.T` names (Codex).
- 2025-10-11: Documented function size limits (≤35 lines, ≤3 parameters) to reinforce existing guardrails (Codex).


## 9. Appendices
- **Glossary**: Embedding (vector representation of text), Scoring (resume-job similarity), Company Filters (platform, status, name), DDD (Domain-Driven Design).
- **External Dependencies**: Playwright/Puppeteer, Xenova transformers for embeddings, Prisma/Postgres, Next.js frontend.
- **Reference Docs**: CLAUDE.md (collaboration norms), README.md (project setup), notes.txt (ad-hoc research/tasks).

## Go Agent Quick Reference
- **DDD Alignment**: Mirror the existing handler → service → repository layering when porting Node tooling to Go. Keep repositories out of HTTP/CLI handlers and funnel orchestration through services.
- **Change Guardrails**: Do not introduce sweeping refactors (new build pipelines, test harness rewrites, etc.) without explicit user approval. Surface blockers before proceeding.
- **Coding Conventions**: Honor the global rule that every function returns a variable named `result`. Avoid one-letter identifiers; prefer self-explanatory names and concise docstrings only when behavior is non-obvious.
- **Tooling Interop**: When Go workflows depend on Node assets, use `pnpm` for any auxiliary Node tasks as mandated in Section 3. Keep generated artifacts consistent with existing PHP output until replacement targets (Go/Kotlin/Java/C#/C++) are introduced.
- **Testing & Builds**: Before handoff, run relevant Go unit/integration tests alongside mandated checks (`pnpm run build:core`, etc.) to ensure cross-project stability.
- **Coordination**: Sync with API/UI/Search agents whenever schema or contract changes ripple beyond the Go generator. Document cross-cutting updates in PRs or handoff notes.
- **Environments & Secrets**: Respect sandbox limits and never commit decrypted secrets. Coordinate with Ops for new environment variables or services, mirroring the procedures outlined above.
- functions no long that 35 lines, and max 3 argumnt. Prohibited long functions.
