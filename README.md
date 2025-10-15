# pdl

## Roadmap

1. **Capture Current PHP Tooling**
   - Collect the existing PHP generators, ORM helpers, and supporting scripts that serve as the source of truth for desired features.
   - Document data flow, configuration flags, and expected outputs for representative PDL models.

2. **Design Go Codegen + CLI**
   - Define the Go package structure that mirrors the PHP feature set (models, persistence helpers, utility layers).
   - Specify CLI requirements (input arguments, config files, templating) and map them to existing workflows.
   - Establish templates and shared helpers that translate PDL AST nodes into idiomatic Go code.

3. **Implement Go Generator and CLI**
   - Port core generators to Go, covering types, properties, methods, and annotations produced today.
   - Recreate ORM behaviour (e.g., CRUD scaffolding, query helpers) with unit coverage and integration tests on sample PDLs.
   - Translate the Node-based `db2Pdl.js` workflow into Go (residing under `pdl-orm/`), preserving current PHP output while laying groundwork for Go/Kotlin/Java/C#/C++ emitters.
   - Ship a runnable CLI command that wraps the new generator, including usage docs and example projects.

4. **Introduce Kotlin and Java Targets**
   - Design template sets for Kotlin and Java that keep parity with the PHP/Go outputs while respecting language conventions.
   - Implement generators that reuse shared translation logic where possible and add language-specific helpers as needed.
   - Plan future language emitters (Go ORM support, Kotlin, Java, C#, C++), sequencing them after the Go CLI migration.
   - Validate output through compilation checks and smoke tests in sample JVM projects.

5. **Consolidate Tooling & Documentation**
   - Add configuration examples showing how to enable each language target and customize output directories.
   - Expand README/docs with migration guidance from PHP tooling, CLI usage, and troubleshooting tips.
   - Set up CI jobs (build + lint) covering Go, Kotlin, and Java outputs to catch regressions early.
