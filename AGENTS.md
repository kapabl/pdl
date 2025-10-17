# Agent Collaboration Guide

## 1. Purpose & Scope
- Coordinate changes across the PDL toolchain (C++ compiler, Go ORM, templates, sample projects).
- Preserve clear layering between metadata discovery, code generation, and downstream runtime packages.
- Use this guide alongside `CLAUDE.md`, `README.md`, and docs under `docs/`.

## 2. Roles & Responsibilities
- **Compiler Agent**: Maintains `pdlc` (C++ compiler) sources, headers, and build scripts.
**ORM Agent**: Owns Go code under `pdl-orm`, ensures generators stay in sync with templates and sample output. Attribute templates stay inside the relevant language generator (match legacy db2Pdl.js); avoid sprawling identifiers—prefer concise, scoped names.
- **Template Agent**: Curates files in `pdl/pdl-orm/internal/db2pdl/templates`, aligns cross-language projections.
- **Sample Agent**: Keeps `sample-project/` variants up to date, verifies configs and generated outputs.
- **QA Agent**: Runs validation suites, compares generated artefacts against legacy baselines, signs off releases.

## 3. Operating Procedures
- **Change Coordination**: Announce refactors that affect shared templates or generator plumbing before landing them.
- **Scope Discipline**: Only touch files requested by the task unless a neighbouring fix is mandatory—flag unexpected work early.
- **Branching & Commits**: Work on the current branch unless the user asks for a dedicated branch or commit.
- **Coding Standards**:
  - Go: descriptive identifiers, no inline `//` comments, keep functions ≤35 lines and ≤3 parameters; return a variable named `result`.
  - C++: prefer RAII, avoid inline comments unless clarifying non-obvious code.
  - Config/JSON: keep ordering stable to minimise churn.
- **Testing**: Run `go test ./...` within touched Go modules; for compiler changes invoke `./pdl/build.sh` as needed. Document any tests you cannot run.
- **Tooling**: Use the existing `build.sh` scripts, `go` tooling, and `cmake` as defined in repo instructions—do not introduce new package managers without approval.
- **Generation Guardrails**: The Go ORM must match legacy behaviour; compare against `sample-project/sample-legacy-orm-output` when altering templates or generators.

## 4. Coordination & Handoffs
- Note cross-cutting template or config updates in PR descriptions or handoff logs.
- Ping the Sample Agent when generated code shape changes so sample projects can be refreshed.
- Escalate conflicts or blocking issues to the QA Agent or repo owner.

## 5. Environments & Secrets
- Local development relies on `.env.local` files within sample projects; never commit secrets.
- Respect sandbox constraints: avoid destructive commands without explicit approval.
- Document new environment variables or required services in `docs/`.

## 6. Quality Gates
- `go test ./...` for any Go package touched.
- `./pdl/build.sh` (or targeted `cmake`/`make` steps) for compiler or runtime changes.
- Regenerate and inspect sample outputs when generator templates change.

## 7. Troubleshooting Quick Reference
- **ORM Generation**: Check DB connectivity, ensure `pdl.db2pdl.config.json` resolves env vars, verify template paths.
- **Template Rendering**: Use `go test` with focused cases or create a scratch binary under `cmd/` to debug payloads.
- **Compiler Build Issues**: Clear `pdl/pdlc/build` artifacts and rerun `./build.sh` with verbose flags.

## 8. Change Log
- 2025-10-14: Re-scoped guide for PDL repository; removed unrelated Node/`pnpm` steps (Codex).
