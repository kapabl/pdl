

Store the following JSON in `pdl.db2pdl.config.json` in the same directory as your `pdl.config.json`. The orchestrator will merge it automatically when the build runs.

```json
{
  "db2pdl": {
    "enabled": true,
    "verbose": false,
    "connection": {
      "type": "mysql",
      "host": "${DB_HOST}",
      "port": "${DB_PORT}",
      "user": "${DB_USERNAME}",
      "password": "${DB_PASSWORD}",
      "database": "${DB_DATABASE}"
    },
    "outputDir": "${PDL_DB2PDL_OUTPUT}",
    "templatesDir": "default",
    "ts": {
      "emit": false,
    },
    "js": {
      "emit": false
    },
    "cs": {
      "emit": false
    },
    "go": {
      "emit": false,
      "package": ""
    },
    "java": {
      "emit": false
    },
    "kotlin": {
      "emit": false
    },
    "rust": {
      "emit": false
    },
    "cpp": {
      "emit": false
    },
    "python": {
      "emit": false,
      "package": ""
    },
    "php": {
      "emit": true
    },
    "pdl": {
      "db2PdlSourceDest": "com/mh/mimanjar/domain/data",
      "entitiesNamespace": "com.mh.mimanjar.domain.data",
      "useNamespaces": []
    },
    "excludedTables": [],
    "excludedColumns": []
  }
}
```

Attributes for DbId and ColumnName will be hardoded inside generators
in the orm/db2pdl\
 -DbId means this is the privary key and should be used for updates
 and put back on create
 -ColumnName is used when you can't translate back and forward the 
 name of the column from the db to a language eg: "time_2" will
 be time2 in TS and Time2 in Go, but you can't fo back to "time_2"
 because "time2" also leads to the same column names


## Runtime usage in `pdl-orm`

**Read by the Go db2pdl runner**

- `db2pdl.enabled`
- `db2pdl.verbose`
- `db2pdl.connection.type`
- `db2pdl.connection.host`
- `db2pdl.connection.port`
- `db2pdl.connection.user`
- `db2pdl.connection.password`
- `db2pdl.connection.database`
- `db2pdl.outputDir`
- `db2pdl.templatesDir`
- `db2pdl.ts.emit`
- `db2pdl.ts.outputFile`
- `db2pdl.js.emit`
- `db2pdl.cs.emit`
- `db2pdl.go.emit`
- `db2pdl.go.package`
- `db2pdl.php.emit`
- `db2pdl.pdl.entitiesNamespace`
  - Required; configuration load will panic if it is missing or empty
- `db2pdl.pdl.useNamespaces`
  - Alias: `db2pdl.pdl.use`
- `db2pdl.excludedTables`
- `db2pdl.excludedColumns`

## Plan

1. Teach `pdl --init` to scaffold `pdl.db2pdl.config.json` next to the generated `pdl.config.json`.
2. Ship a migration checklist for repos that still embed the `db2pdl` block inline.
3. Mirror the documentation updates into release notes so downstream agents coordinate on the new layout.
4. Ensure downstream CI jobs copy both config files when packaging sample projects.


## Execution Step
- Update the init command to emit the new external file and adjust scaffolding tests accordingly.

## Verification Phase

- Run `go test ./...` inside `pdl/pdl-orm` to confirm the generators compile after changes.
- Execute a sample run (e.g., `sample-project/pdl-project/scripts/run-orm-mysql.sh`) to inspect generated output when generator behaviour changes.





## Language specific doc
- TypeScript/JavaScript split between backend (db access) and frontend DTO layers; backend supports JSON encode/decode and likely protobuf in gRPC, frontend supplies typed row DTOs with JSON encode/decode only.
- Go emits packages under `output/db2pdl/go/<namespace>/` with CRUD helpers, import wiring, and query builders.
- PHP produces row classes, column definitions, where/order builders, and helper traits compatible with the legacy infrastructure loader.
- C# generates row set classes and fluent accessors targeting the original .NET integration points.
- Java, Kotlin, Rust, and C++ receive backend-oriented DTOs plus fluent query builders mirroring the MiManjar legacy output contract.
- Python output mirrors backend DTO semantics for service-layer consumption (JSON encode/decode, no frontend bundle).
