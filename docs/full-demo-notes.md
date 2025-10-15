# Full Demo Notes (online.ordering)

This project demonstrates the original PHP + JavaScript runtime built around PDL.
Key observations for mirroring behaviour on the Go side:

- **Generated Assets**
  - PHP rows live under `full-demo/online.ordering/mimanjar/gen/php/src/Com/Mh/Mimanjar/Domain/Data/`.
  - JavaScript helpers reside in `gen/js/src/com/mh/mimanjar/domain/…`, with a barrel file `gen/js/src/MiManjarPdl.js`.
  - Source PDL files are in `pdl-project/src/com/mh/mimanjar/domain/**`.

- **Row Consumption Patterns (PHP)**
  - Rows expose a static `where()` builder that returns fluent query helpers. Example: `AddressRow::where()->userId($userId)->loadRows()`.
  - `::createInstance()` seeds an empty row object that is then populated and persisted via `->create()`.
  - Domain services keep the row types in their namespace, e.g. `app/Com/Mh/Mimanjar/Application/AddressesController.php` orchestrates persistence exclusively through generated rows.

- **Supporting Helpers**
  - Column definitions (`AddressColumns`, etc.) are used when composing SQL fragments or selections.
  - Trait files such as `ColumnsTraits` provide fluent helpers for building select lists.
  - Where / OrderBy classes are thin wrappers that call `$this->addField('column')` during setup.

- **Runtime Wiring**
  - `PdlServiceProvider` registers the generated assets, ensuring Laravel’s container resolves row factories and DB utilities.
  - The PHP runtime relies on `Com\Mh\Ds\Infrastructure` base classes (Row, Where, OrderBy, FieldList, ColumnsDefinition) – our Go equivalents need to offer the same surface area (Row metadata, fluent query builder, multi insert/delete helpers).

This overview should guide parity work when creating Go/Kotlin/Java/Rust flavours: the generated code must expose Row structs with `Create/Update/Delete`, fluent `Where` builders, and static column metadata mirroring the PHP implementation.
