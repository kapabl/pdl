package sample

import (
	"testing"
	"time"

	orm "github.com/kapablanka/pdl/sample/com/mh/mimanjar/domain/data"
)

func TestCategoryProductLifecycle(t *testing.T) {
	if testStore == nil || !tableLookupExists("category_products") {
		t.Skip("category_products integration not available")
	}

	var createdID int64
	defer func() {
		if createdID != 0 {
			mustExec("DELETE FROM category_products WHERE id = ?", createdID)
		}
	}()

	row := orm.CategoryProductModel.New()
	row.CategoryId = 123
	row.ProductId = 456
	row.StoreId = 789
	row.Position = 1
	now := time.Now().UTC().Truncate(time.Second)
	row.CreatedAt = now
	row.UpdatedAt = now

	if err := row.Create_(); err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if row.Id == 0 {
		t.Fatalf("expected autoincremented id to be set after Create_, got 0")
	}
	createdID = row.Id

	baseBuilder := func() orm.CategoryProductWhereBuilder {
		return orm.CategoryProductModel.Where().
			CategoryId(row.CategoryId).
			ProductId(row.ProductId).
			StoreId(row.StoreId).
			CreatedAt(row.CreatedAt)
	}
	testCases := []struct {
		scenarioName string
		loadRows     func() ([]*orm.CategoryProduct, error)
	}{
		{
			scenarioName: "base",
			loadRows: func() ([]*orm.CategoryProduct, error) {
				return baseBuilder().LoadRows_()
			},
		},
		{
			scenarioName: "field_list",
			loadRows: func() ([]*orm.CategoryProduct, error) {
				return baseBuilder().
					FieldList_(orm.CategoryProductColumns.Id, orm.CategoryProductColumns.CategoryId).
					LoadRows_()
			},
		},
		{
			scenarioName: "order_limit",
			loadRows: func() ([]*orm.CategoryProduct, error) {
				return baseBuilder().
					OrderBy_().Id().Desc_().
					Limit_(1).
					LoadRows_()
			},
		},
		{
			scenarioName: "range_order",
			loadRows: func() ([]*orm.CategoryProduct, error) {
				return baseBuilder().
					OrderBy_().Id().Asc_().
					Range_(0, 1).
					LoadRows_()
			},
		},
	}

	for _, scenario := range testCases {
		rows, fetchErr := scenario.loadRows()
		if fetchErr != nil {
			t.Fatalf("%s fetch failed: %v", scenario.scenarioName, fetchErr)
		}
		if len(rows) != 1 {
			t.Fatalf("%s expected single row, got %d", scenario.scenarioName, len(rows))
		}
		if rows[0].Id != createdID {
			t.Fatalf("%s returned unexpected id %d (expected %d)", scenario.scenarioName, rows[0].Id, createdID)
		}
	}

	if loaded := fetchCategoryProduct(t, createdID); loaded.Position != 1 {
		t.Fatalf("expected position to be 1, got %d", loaded.Position)
	}

	row.Position = 2
	row.UpdatedAt = now.Add(time.Hour)
	if err := row.Update_(); err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if loaded := fetchCategoryProduct(t, createdID); loaded.Position != 2 {
		t.Fatalf("expected position to be 2, got %d", loaded.Position)
	}

	if err := orm.CategoryProductModel.Where().Id(createdID).Delete_(); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if fetched, err := orm.CategoryProductModel.Where().Id(createdID).OrderBy_().Id().Asc_().LoadRows_(); err != nil {
		t.Fatalf("Load after delete failed: %v", err)
	} else if len(fetched) != 0 {
		t.Fatalf("expected 0 rows after delete, got %d", len(fetched))
	}
}

func TestCategoryProductMetadata(t *testing.T) {
	if orm.CategoryProductColumns.CategoryId != "category_id" {
		t.Fatalf("unexpected column mapping: %+v", orm.CategoryProductColumns)
	}
	columns := orm.CategoryProductColumnList()
	for _, must := range []string{"id", "category_id", "product_id"} {
		if !contains(columns, must) {
			t.Fatalf("expected column list to contain %s, got %#v", must, columns)
		}
	}
}

func fetchCategoryProduct(t *testing.T, id int64) *orm.CategoryProduct {
	t.Helper()
	records, err := orm.CategoryProductModel.Where().Id(id).LoadRows_()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("expected single row, got %d", len(records))
	}
	return records[0]
}

func contains(collection []string, target string) bool {
	for _, item := range collection {
		if item == target {
			return true
		}
	}
	return false
}
