# PDL AST JSON Schema

This document describes the JSON document emitted by `pdlc` when the AST writer is enabled. The schema is stable under a `version`ed envelope so downstream generators can evolve independently from the parser.

Enable or disable the writer in any compiler profile by adding an `ast` block to the compiler configuration passed to `pdlc` (for example `config/compiler/pdl.js.json`):

```json
{
  "ast": {
    "enabled": true,
    "outputDir": "output/ast"
  }
}
```

If `outputDir` is omitted the compiler writes into `<out>/ast`, where `<out>` is the profile output directory supplied to `pdlc`.

## Root Structure

```json
{
  "version": 1,
  "source": { ... },
  "namespace": { ... }
}
```

- `version`: Integer schema revision. The initial release is `1`.
- `source`: Metadata about the compiled file.
- `namespace`: AST payload for the parsed namespace.

## Source Block

| Field | Type | Description |
| ----- | ---- | ----------- |
| `file` | string | Path to the compiled `.pdl` file as provided to the compiler. |
| `namespace` | object | Identifier object matching the namespace declaration. |

## Identifier Object

| Field | Type | Description |
| ----- | ---- | ----------- |
| `segments` | string[] | Ordered identifier segments. |
| `qualifiedName` | string | Dot-joined identifier suitable for display. |

Identifiers appear in namespace names, using directives, type references, parent classes, and attribute names.

## Namespace Block

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | object | Identifier for the namespace declaration. |
| `usings` | object[] | Collection of imported identifiers. |
| `classes` | object[] | Class declarations within the namespace. |

## Class Object

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | string | Declared class name. |
| `accessModifier` | string | Access modifier (`public`, `internal`, `protected`, `private`, or empty if unspecified). |
| `parent` | object or null | Identifier of the base class when present. |
| `members` | object[] | Ordered member declarations. |

## Member Objects

Each member object contains a `kind` discriminator:

- `property`: Full property declaration with attributes and access modifiers.
- `shortProperty`: Compact property syntax without attributes or modifiers.
- `method`: Method declaration.
- `const`: Constant declaration.

All member objects contain:

| Field | Applicable kinds | Type | Description |
| ----- | ---------------- | ---- | ----------- |
| `kind` | all | string | Member category discriminator. |
| `name` | all | string | Member name. |
| `accessModifier` | property, method | string | Access modifier name or empty string. |

### Property Members

| Field | Type | Description |
| ----- | ---- | ----------- |
| `type` | object | Property type descriptor (see below). |
| `attributes` | object[] | Attribute descriptors associated with the property. |
| `arguments` | object[] | Indexer argument list for the property. |
| `access` | string | Access specifier (`read`, `write`, `readWrite`, or empty). |

### Short Property Members

- Same as property members except `attributes` and `arguments` are empty arrays and `accessModifier` is empty.

### Method Members

| Field | Type | Description |
| ----- | ---- | ----------- |
| `returnType` | object | Identifier describing the return type. |
| `arguments` | object[] | Method arguments. |

### Const Members

| Field | Type | Description |
| ----- | ---- | ----------- |
| `type` | object | Identifier describing the constant type. |
| `value` | string, number, or boolean | Literal value as parsed. |

## Type and Argument Helpers

### Property Type

| Field | Type | Description |
| ----- | ---- | ----------- |
| `type` | object | Identifier describing the element type. |
| `arrayNotation` | string[] | Array suffix segments (e.g., `"[]"`, `"[10]"`). |

### Argument

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | string | Argument name. |
| `type` | object | Identifier describing the argument type. |

## Attributes

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | object | Identifier describing the attribute. |
| `params` | object | Attribute parameters. |

The `params` object contains:

| Field | Type | Description |
| ----- | ---- | ----------- |
| `required` | array | Ordered list of positional arguments (strings, numbers, or booleans). |
| `optional` | object[] | Named arguments preserving declaration order. |

Optional parameter entries contain:

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | string | Parameter identifier. |
| `value` | string, number, or boolean | Literal value assigned to the parameter. |

## Literals

Literal values in attributes and constants preserve their native JSON types:

- Strings remain UTF-8 sequences.
- Numbers are emitted as integers or doubles.
- Booleans retain `true`/`false`.

Downstream generators should treat missing strings as empty values rather than null references.
