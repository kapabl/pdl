package sample

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"testing"
)

// TestGeneratedRowAPI inspects one of the generated Go files and asserts the expected
// helper functions and types exist. This keeps the generator aligned with the runtime
// expectations enforced by our PHP reference implementation.
func TestGeneratedRowAPI(t *testing.T) {
	generatedFile := filepath.Join("..", "pdl-project", "output", "db2pdl", "go", "com", "mh", "mimanjar", "domain", "data", "category_products.go")

	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, generatedFile, nil, parser.AllErrors)
	if err != nil {
		t.Fatalf("parse of generated file failed: %v", err)
	}

	symbols := map[string]bool{
		"CategoryProduct":                          false,
		"CategoryProductModel":                     false,
		"CategoryProductWhereBuilder":              false,
		"CategoryProductOrderByBuilder":            false,
		"CategoryProductOrderByDirectionBuilder":   false,
		"CategoryProductColumns":                   false,
		"CategoryProductOrderBy":                   false,
		"CategoryProductColumnList":                false,
		"CategoryProduct.Create_":                  false,
		"CategoryProduct.Update_":                  false,
		"CategoryProduct.Delete_":                  false,
		"CategoryProductWhereBuilder.Project_":     false,
		"CategoryProductWhereBuilder.FieldList_":   false,
		"CategoryProductWhereBuilder.Offset_":      false,
		"CategoryProductWhereBuilder.Limit_":       false,
		"CategoryProductWhereBuilder.Range_":       false,
		"CategoryProductWhereBuilder.OrderBy_":     false,
		"CategoryProductWhereBuilder.LoadRows_":    false,
		"CategoryProductOrderByBuilder.LoadRows_":  false,
		"CategoryProductOrderByBuilder.Project_":   false,
		"CategoryProductOrderByBuilder.FieldList_": false,
		"CategoryProductOrderByBuilder.Offset_":    false,
		"CategoryProductOrderByBuilder.Limit_":     false,
		"CategoryProductOrderByBuilder.Range_":     false,
		"categoryProductAccessor.New":              false,
		"categoryProductAccessor.Where":            false,
	}

	ast.Inspect(file, func(node ast.Node) bool {
		switch typed := node.(type) {
		case *ast.TypeSpec:
			if _, found := symbols[typed.Name.Name]; found {
				symbols[typed.Name.Name] = true
			}
		case *ast.FuncDecl:
			if typed.Recv == nil {
				if _, found := symbols[typed.Name.Name]; found {
					symbols[typed.Name.Name] = true
				}
				return true
			}
			if len(typed.Recv.List) == 0 {
				return true
			}
			receiver := typed.Recv.List[0].Type
			var receiverName string
			switch expr := receiver.(type) {
			case *ast.Ident:
				receiverName = expr.Name
			case *ast.StarExpr:
				if ident, ok := expr.X.(*ast.Ident); ok {
					receiverName = ident.Name
				}
			}
			key := receiverName + "." + typed.Name.Name
			if _, found := symbols[key]; found {
				symbols[key] = true
			}
		case *ast.GenDecl:
			if typed.Tok == token.VAR {
				for _, spec := range typed.Specs {
					if valueSpec, ok := spec.(*ast.ValueSpec); ok {
						for _, ident := range valueSpec.Names {
							if _, found := symbols[ident.Name]; found {
								symbols[ident.Name] = true
							}
						}
					}
				}
			}
		}
		return true
	})

	for name, seen := range symbols {
		if !seen {
			t.Fatalf("expected generated symbol %s to be present", name)
		}
	}
}
