package gofiles

import (
	"path/filepath"
	"strings"

	gen "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/mysql"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/postgres"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

var _ gen.Generator = (*Generator)(nil)

type Generator struct {
	renderer gen.Renderer
	write    gen.CodeWriter
	config   shared.DB2PDLConfig
}

func New(renderer gen.Renderer, writer gen.CodeWriter, config shared.DB2PDLConfig) Generator {
	return Generator{
		renderer: renderer,
		write:    writer,
		config:   config,
	}
}

func (generator Generator) Generate(table shared.TableData) error {
	renderTable := generator.InjectAttributes(table)
	additionalImports := make([]string, 0)
	for _, importPath := range renderTable.GoImports {
		if importPath == "github.com/kapablanka/pdl/pdl/infra/go" {
			continue
		}
		additionalImports = append(additionalImports, importPath)
	}
	goNamespacePath := shared.NamespaceToGoPath(renderTable.PdlEntitiesNamespace)
	goPackageName := shared.NamespaceToGoPackage(renderTable.PdlEntitiesNamespace)
	if goPackageName == "" {
		goPackageName = strings.ToLower(generator.config.Go.Package)
	}
	relativeFileName := renderTable.TableName
	if goNamespacePath != "" {
		relativeFileName = filepath.Join(goNamespacePath, renderTable.TableName)
	}
	payload := struct {
		shared.TableData
		GoPackage         string   `json:"goPackage"`
		AdditionalImports []string `json:"additionalImports"`
	}{
		TableData:         renderTable,
		GoPackage:         goPackageName,
		AdditionalImports: additionalImports,
	}
	source, renderErr := generator.renderer.Render(filepath.Join("go", "row"), payload)
	if renderErr != nil {
		return renderErr
	}
	return generator.write("go", relativeFileName, source)
}

func (generator Generator) InjectAttributes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.GoType = resolveGoType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolveGoType(driver string, dbType string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		return mysql.GoType(dbType)
	case "postgres", "postgresql":
		return postgres.GoType(dbType)
	default:
		return "string"
	}
}
