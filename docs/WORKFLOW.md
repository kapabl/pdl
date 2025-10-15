
- [x] convert all hbs inside /pdl to go templates
- [x] fix pdl-orm to fill this templates
- [x] generate the pdl files with this templates and make sure they are correct, the grammar
of pdl is expressed inside pdlc using boots/spirit
- [x] change all references to io.pdl.... to io.pdl...

- [x] same with this use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition
    lets convert it to Io\Pdl\Infra....

---

- [x] read full-demo and learn how the previous pdl generated code was used around the orm classes, 
    fully typed, we need to acomplish the same objective in go, kotlin, java and rust.
- [x] eliminate all dependencies of sample-project from gulp and node
- [x] inside sample-project/go/ create testing for the generated classes

---

- [x] align generated Go row constructors with the roadmap “perfect code” API (e.g., `row := AddressRow.New()` semantics without helper clutter).
- [x] derive Go package/module names from the company/project namespace so imports read like `com/mh/mimanjar/...` instead of generic `generated`.
- [x] mirror the PHP `_loadRows()` pattern on Go query builders to avoid collisions with column names while keeping fluent chaining.
- [x] document the Go usage story in `docs/usage.md`, mirroring the PHP examples from full-demo.
- [x] drop the `Row` suffix from generated Go structs so constructors look like `Category.New()` and usage mirrors the roadmap examples.
- [x] Add annotation in Go structs for PK/auto-increment metadata so downstream stores can introspect keys without manual wiring.

- [x] Default CLI workflow for downstream projects
    1. Install the PDL CLI via your package manager (e.g. `dnf install pdl`, `apt install pdl`, or pull the official container image).
    2. From the repository root run `pdl --init`. This scaffolds `pdl/` with a starter `pdl.config.json`, a `src/` folder, and `.env.local` containing example variables.
    3. Subsequent invocations of `pdl` (with no extra arguments) run `pdl --build --config pdl/pdl.config.json`, so everyday builds are `pdl` from the repo root.

- [x] Write driver for infra/go for postgres
- [x] Make pdl-orm generate different code if the target is postgres. THe orm should read the config and instantiate the correct generator based on the db
    - [x] Testing by having a second config in sample-projects for postgres
    - [x] Support posgress vector to allow using it for embeddings
- [ ] Whenever touching legacy Go code, split any functions longer than 35 lines or taking more than three parameters into smaller helpers to comply with the shared coding standards.

- [ ] Improving init
    -I would say thinkling more about this that a user may choose like a stack ( 
    -Backeend( make then select from a list of language, can be more than 1 )
    here they can select for example( Go, Java, Kotlin, Rust, C#, maybe GraphQL and proto)

    -make then select From a list( React or Vue, Ts and/or JS ) 

    -This is during init we ask this, including the values we have how in the almost empty pdl.config.json

    -From there we generate the config(s) for the user

    -So this happens when the user does --init

- [ ] AST output pipeline
    - [ ] Create AST output generator.
    - [ ] Document the AST nodes.
    - [ ] Convert all generators to generate from the AST.
    - [ ] Make each generator a CLI so anyone can add another CLI for their language.

- [ ] convert all generator to go and make them convert from json ast
    - [ ] Go
    - [ ] JS
    - [ ] PHP
    - [ ] CSharp
    - [ ] Rust
    - [ ] Cpp
    - Introduce new generators
        - [ ] java
        - [ ] kotlin
        - [ ] Rust
- [ ] create infras for
    - [ ] PHP <--already exists, extract from old pdl projec
    - [ ] Go, this one is amost done
    - [ ] Kotlin
    - [ ] Java
    - [ ] Rust
    - [ ] Cpp
    - [ ] Ts backend side
    - [ ] Ts frontend side

- [ ] pdl and/or pdlbuild should act as orquestrators to
    call the generators after pdlc generate the ast
    - [ ] when all external generators are working
        remove generators from pdlc, clean pdlc.
    - [ ] new generators should use Go templates    
