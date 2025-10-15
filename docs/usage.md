# Generated ORM Usage

## CLI Quickstart

```
dnf install pdl    # or apt install pdl / docker run ...
pdl --init         # creates ./pdl with config, src/, and .env.local
pdl                # equivalent to pdl --build --config pdl/pdl.config.json
```

The scaffolded `.env.local` supplies the output directories and database placeholders; update those values before generating code against a real schema.

## Go

The Go ORM output lives under `output/db2pdl/go` using the project namespace as the package path. For example, the MiManjar sample project produces:

```
output/db2pdl/go/com/mh/mimanjar/domain/data/category_products.go
```

Import the package using its namespace-derived path and work through the exported accessor:

```go
import (
    orm "github.com/kapablanka/pdl/sample/com/mh/mimanjar/domain/data"
)

func persistCategoryProduct() error {
    record := orm.CategoryProduct.New()
    record.CategoryId = 123
    record.ProductId = 456
    return record.Create_()
}

func loadTypedRows() ([]*orm.CategoryProductRow, error) {
    return orm.CategoryProduct.Where().
        StoreId(789).
        OrderBy_().StoreId().Asc_().CategoryId().Desc_().
        Range_(5, 20).
        FieldList_(orm.CategoryProductColumns.StoreId, orm.CategoryProductColumns.ProductId).
        LoadRows_()
}
```

The accessor exposes:

- `New()` – constructs a `*<Entity>Row` with CRUD helpers wired up.
- `Where()` – returns a fluent builder backed by `pdlgo.QueryBuilder`.
- `LoadRows_()` – materialises typed records while keeping the lower-level `Load()` method for raw maps.
- `Create_()` / `Update_()` / `Delete_()` – row-level helpers that stay out of the column-name namespace.
- `OrderBy_()` / `FieldList_()` / `Limit_()` / `Offset_()` / `Range_()` – fluent helpers for sorting, projection, and pagination.

The generated module declares `module github.com/kapablanka/pdl/sample` inside `output/db2pdl/go`, so consumers can target the namespace path directly or use a `replace` directive during local development:

```
replace github.com/kapablanka/pdl/sample => ../pdl-project/output/db2pdl/go
```

Remember to call `pdlgo.SetDefaultStore` (or pass a store via `Where().WithStore(...)`) before issuing queries so the accessor knows how to talk to your persistence layer.
