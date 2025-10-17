# PDL Configuration Guide

The Go CLI now reads **one canonical file** per project: `pdl.config.json`.  
All build settings—templates, profiles, sections, optional services—live inside that file so teams no longer manage a secondary `config/` tree or language‑specific profile files.

---

## 1. Top-Level Settings

```json
{
  "companyName": "Acme",
  "project": "Storefront",
  "version": "1.0.0",
  "outputDir": "${PDL_OUTPUT}",
  "verbose": false,
  "templates": {
    "dir": "${PDL_TEMPLATES_DIR}",
    "name": "classTemplate1"
  },
  "db2pdl": {
    "enabled": true,
    "db2PdlSourceDest": "acme/storefront/domain/data"
  },
  "src": ["src"],
  …
}
```

- `outputDir` is the base for every language (`output/<language>`); if omitted it defaults to `output`.
- `templates` defines the default template directory and class template name. Individual profiles can override either field when necessary.
- `db2pdl` controls the db2pdl and ORM scaffolding; when enabled, set `db2PdlSourceDest` to the namespace you want to populate. The CLI writes ORM artifacts to `${outputDir}/pdl` unless you override `db2pdl.outputDir`.

Environment placeholders (e.g. `${PDL_OUTPUT}`) are expanded before the config is parsed, so the CLI still honours the same `.env` or shell exports as the legacy tooling.

---

## 2. Profiles

The `profiles` object declares each generator you want to run.

```json
"profiles": {
  "ts": {
    "language": "ts",
    "enabled": true
  },
  "tsConsts": {
    "language": "ts",
    "enabled": true,
    "templates": {
      "dir": "${PDL_TEMPLATES_DIR}/js-object"
    }
  },
  "js": {
    "language": "js",
    "enabled": true,
    "generateAsObject": true,
    "namespaceFile": "",
    "useLet": true
  },
  "go": {
    "enabled": true
  },
  "php": {
    "enabled": true
  }
}
```

Key points:

- `language` is optional; when present it groups multiple profiles under the same language short code (`ts` and `tsConsts` above). When omitted, the profile name itself is used.
- Each profile inherits the top-level template settings and `outputDir/<language>` by default. Any value you provide inside the profile (`templates.dir`, `templates.name`, additional flags like `generateAsObject`) overrides just that field.
- Additional profile-specific fields are simply forwarded to the generator—no extra config scaffolding is required.

The CLI currently supports inhibitors for the standard generators (`ts`, `js`, `go`, `php`, `csharp`, `java`). Custom generators can read the same structure.

---

## 3. Sections

Sections control which `.pdl` files map to which profiles. The schema is unchanged, but now references only the profile names declared above:

```json
"sections": [
  {
    "name": "Frontend",
    "files": {
      "ts": ["**/*.pdl"],
      "tsExclude": []
    }
  },
  {
    "name": "Domain",
    "files": {
      "go": ["**/*.pdl"],
      "php": ["**/*.pdl"],
      "java": ["**/*.pdl"],
      "goExclude": [],
      "phpExclude": [],
      "javaExclude": []
    }
  }
]
```

Use `<profile>Name` as the key, plus the optional `<profile>NameExclude` entry for path-based exclusions.

---

## 4. Db2Pdl/ORM and Optional Services

Place the ORM payload in `pdl.db2pdl.config.json`, stored alongside `pdl.config.json`. The runner merges it automatically before normalisation:

```json
{
  "db2pdl": {
    "enabled": true,
    "db2PdlSourceDest": "io/myproject/domain/data",
    "outputDir": "${PDL_DB2PDL_OUTPUT}"
  }
}
```

Existing projects that still embed the block inside `pdl.config.json` continue to work, but new scaffolds should rely on the external file. When `outputDir` is omitted, it resolves to `<outputDir>/pdl` automatically (or `${PDL_DB2PDL_OUTPUT}` when set).

---

## 5. Environment Variables

Only a few environment variables remain mandatory:

| Variable             | Purpose                                |
|----------------------|----------------------------------------|
| `PDL_OUTPUT`         | Base output directory                  |
| `PDL_TEMPLATES_DIR`  | Directory containing generator templates |
| `PDL_DB2PDL_OUTPUT`  | Destination for db2pdl stage (if used)  |
| Database credentials | Required when db2pdl/ORM stages run     |

Language-specific copies (`PDL_GEN_OUTPUT_*`) are no longer needed; the CLI derives them from `outputDir` plus the profile’s language code.

---

## 6. Migration Notes

1. Delete any legacy `config/` profile JSON files and `.pdl/defaults` directories—the CLI will ignore them.
2. Move generator overrides (e.g. TS const templates, JS flags) into the corresponding profile block.
3. Remove obsolete `configFile` references and profile `outputDir` entries unless you need per-profile overrides.
4. If you previously toggled rebuild behaviour via `rebuild` in `common.pdl.config.json`, use CLI flags instead (`pdl --rebuild`).

---

With these changes, `pdl.config.json` and `pdl.db2pdl.config.json` act as a paired source of truth, and every binary in the monorepo (CLI, ORM tools, C++ compiler) reads the same structure. Edit the files once, run `pdl --build`, and the rest of the toolchain follows automatically.
