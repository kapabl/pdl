package sample

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedGoSourcesTypecheck(t *testing.T) {
	generatedDir := filepath.Join("..", "pdl-project", "output", "db2pdl", "go")
	fileset := token.NewFileSet()
	astFiles := make([]*ast.File, 0)
	walkErr := filepath.WalkDir(generatedDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".go") {
			return nil
		}
		parsedFile, parseErr := parser.ParseFile(fileset, path, nil, parser.AllErrors)
		if parseErr != nil {
			return parseErr
		}
		astFiles = append(astFiles, parsedFile)
		return nil
	})
	if walkErr != nil {
		t.Fatalf("failed to traverse generated go directory: %v", walkErr)
	}
	if len(astFiles) == 0 {
		t.Fatalf("no generated go files found under %s", generatedDir)
	}
	config := types.Config{
		Importer: importer.For("source", nil),
	}
	if _, checkErr := config.Check("github.com/kapablanka/pdl/sample", fileset, astFiles, nil); checkErr != nil {
		t.Fatalf("type checking generated go sources failed: %v", checkErr)
	}
}
