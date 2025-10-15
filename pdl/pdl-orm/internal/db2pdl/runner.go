package db2pdl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Runner struct {
	ConfigPath   string
	ExitWhenDone bool
}

type executionStats struct {
	totalTables int
	totalFiles  int
	totalLines  int
	totalSize   int64
}

func (runner Runner) Run(ctx context.Context) error {
	var result error
	configuration, loadErr := LoadConfig(runner.ConfigPath)
	if loadErr != nil {
		result = loadErr
		return result
	}
	dbConfig := configuration.DB2PDL
	if !dbConfig.Enabled {
		result = nil
		return result
	}
	databaseType := strings.ToLower(dbConfig.Connection.Type)
	if databaseType == "" {
		result = fmt.Errorf("database connection type is required")
		return result
	}
	entityNamespace, phpDefaultNamespace := calculateNamespaces(dbConfig)
	driver, driverErr := createMetadataDriver(databaseType, dbConfig, entityNamespace, phpDefaultNamespace)
	if driverErr != nil {
		return driverErr
	}
	verbosePrinter := NewVerbosePrinter(dbConfig.Verbose || configuration.Verbose)
	cleanErr := CleanOutputDirectory(dbConfig.OutputDir)
	if cleanErr != nil {
		result = cleanErr
		return result
	}
	templatesDir, useExternal := ResolveTemplatesDir(dbConfig.TemplatesDir)
	templateUtils, templateErr := NewTemplateUtils(configuration, templatesDir, useExternal)
	if templateErr != nil {
		result = templateErr
		return result
	}
	sqlConnection, connectionErr := driver.Open(ctx)
	if connectionErr != nil {
		return connectionErr
	}
	defer sqlConnection.Close()
	stats := executionStats{}
	tables, tablesErr := driver.ListTables(ctx, sqlConnection)
	if tablesErr != nil {
		return tablesErr
	}
	typeScriptBlocks := make([]string, 0)
	for _, tableName := range tables {
		if ContainsCaseSensitive(dbConfig.ExcludedTables, tableName) {
			continue
		}
		baseName := ToPascalCase(Singularize(tableName))
		if baseName == "" {
			baseName = ToPascalCase(tableName)
		}
		tableData, tableErr := driver.BuildTableData(ctx, sqlConnection, tableName, baseName)
		if tableErr != nil {
			result = tableErr
			return result
		}
		stats.totalTables++
		generateErr := runner.generateClasses(templateUtils, dbConfig, tableData, &stats, &typeScriptBlocks, verbosePrinter)
		if generateErr != nil {
			result = generateErr
			return result
		}
	}
	finishErr := runner.finishGeneration(templateUtils, dbConfig, typeScriptBlocks, &stats, verbosePrinter)
	if finishErr != nil {
		result = finishErr
		return result
	}
	if runner.ExitWhenDone {
		os.Exit(0)
	}
	result = nil
	return result
}

func (runner Runner) generateClasses(templateUtils *TemplateUtils, mysqlConfig DB2PDLConfig, tableData TableData, stats *executionStats, typeScriptBlocks *[]string, verbosePrinter VerbosePrinter) error {
	var result error
	pdlSource, pdlErr := templateUtils.Render(filepath.Join("pdl", "pdlRowClass"), tableData)
	if pdlErr != nil {
		result = pdlErr
		return result
	}
	pdlOutputErr := runner.outputCode(pdlSource, tableData.PdlRowClass, "pdl", mysqlConfig.OutputDir, stats, verbosePrinter)
	if pdlOutputErr != nil {
		result = pdlOutputErr
		return result
	}
	if mysqlConfig.PHP.EmitHelpers {
		phpTemplates := []struct {
			Template string
			FileName string
		}{
			{Template: filepath.Join("php", "rowClass"), FileName: tableData.RowClass},
			{Template: filepath.Join("php", "columnsDefinitionClass"), FileName: tableData.ColumnsDefinitionClass},
			{Template: filepath.Join("php", "whereClass"), FileName: tableData.WhereClass},
			{Template: filepath.Join("php", "orderByClass"), FileName: tableData.OrderByClass},
			{Template: filepath.Join("php", "columnsListTraits"), FileName: tableData.ColumnsListTraits},
			{Template: filepath.Join("php", "fieldListClass"), FileName: tableData.FieldListClass},
		}
		for _, entry := range phpTemplates {
			phpSource, phpErr := templateUtils.Render(entry.Template, tableData)
			if phpErr != nil {
				result = phpErr
				return result
			}
			phpOutputErr := runner.outputCode(phpSource, entry.FileName, "php", mysqlConfig.OutputDir, stats, verbosePrinter)
			if phpOutputErr != nil {
				result = phpOutputErr
				return result
			}
		}
	}
	if mysqlConfig.TypeScript.Emit {
		tsSource, tsErr := templateUtils.Render(filepath.Join("ts", "dbBlock"), tableData)
		if tsErr != nil {
			result = tsErr
			return result
		}
		*typeScriptBlocks = append(*typeScriptBlocks, tsSource)
	}
	if mysqlConfig.Go.Emit {
		additionalImports := make([]string, 0)
		for _, importPath := range tableData.GoImports {
			if importPath == "github.com/kapablanka/pdl/pdl/infra/go" {
				continue
			}
			additionalImports = append(additionalImports, importPath)
		}
		goNamespacePath := NamespaceToGoPath(tableData.PdlEntitiesNamespace)
		goPackageName := NamespaceToGoPackage(tableData.PdlEntitiesNamespace)
		if goPackageName == "" {
			goPackageName = strings.ToLower(mysqlConfig.Go.Package)
		}
		relativeFileName := tableData.TableName
		if goNamespacePath != "" {
			relativeFileName = filepath.Join(goNamespacePath, tableData.TableName)
		}
		goPayload := struct {
			TableData
			GoPackage         string   `json:"goPackage"`
			AdditionalImports []string `json:"additionalImports"`
		}{
			TableData:         tableData,
			GoPackage:         goPackageName,
			AdditionalImports: additionalImports,
		}
		goSource, goErr := templateUtils.Render(filepath.Join("go", "row"), goPayload)
		if goErr != nil {
			result = goErr
			return result
		}
		goOutputErr := runner.outputCode(goSource, relativeFileName, "go", mysqlConfig.OutputDir, stats, verbosePrinter)
		if goOutputErr != nil {
			result = goOutputErr
			return result
		}
	}
	if mysqlConfig.CSharp.Emit {
		csSource, csErr := templateUtils.Render(filepath.Join("cs", "csharpRowSetClass"), tableData)
		if csErr != nil {
			result = csErr
			return result
		}
		csOutputErr := runner.outputCode(csSource, tableData.CsharpRowSetClass, "cs", mysqlConfig.OutputDir, stats, verbosePrinter)
		if csOutputErr != nil {
			result = csOutputErr
			return result
		}
	}
	result = nil
	return result
}

func (runner Runner) outputCode(source string, fileName string, language string, outputDir string, stats *executionStats, verbosePrinter VerbosePrinter) error {
	var result error
	targetDir := outputDir
	if language != "" {
		targetDir = filepath.Join(outputDir, language)
	}
	createErr := os.MkdirAll(targetDir, 0o755)
	if createErr != nil {
		result = createErr
		return result
	}
	targetName := fileName
	if language != "" {
		targetName = targetName + "." + language
	}
	outputPath := filepath.Join(targetDir, targetName)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		result = err
		return result
	}
	writeErr := os.WriteFile(outputPath, []byte(source), 0o644)
	if writeErr != nil {
		result = writeErr
		return result
	}
	verbosePrinter.Println(outputPath + " generated.")
	stats.totalFiles++
	stats.totalSize += int64(len(source))
	stats.totalLines += strings.Count(source, "\n")
	result = nil
	return result
}

func (runner Runner) finishGeneration(templateUtils *TemplateUtils, mysqlConfig DB2PDLConfig, typeScriptBlocks []string, stats *executionStats, verbosePrinter VerbosePrinter) error {
	var result error
	if mysqlConfig.TypeScript.Emit && len(typeScriptBlocks) > 0 {
		payload := map[string]interface{}{
			"source": strings.Join(typeScriptBlocks, "\n"),
		}
		tsSource, tsErr := templateUtils.Render(filepath.Join("ts", "dbModule"), payload)
		if tsErr != nil {
			result = tsErr
			return result
		}
		tsOutputErr := runner.outputCode(tsSource, mysqlConfig.TypeScript.OutputFile, "ts", mysqlConfig.OutputDir, stats, verbosePrinter)
		if tsOutputErr != nil {
			result = tsOutputErr
			return result
		}
	}
	summary := fmt.Sprintf("Stats: %s files generated, %s lines generated, %s bytes", strconv.Itoa(stats.totalFiles), strconv.Itoa(stats.totalLines), strconv.FormatInt(stats.totalSize, 10))
	fmt.Printf("Output dir: %s\n", mysqlConfig.OutputDir)
	fmt.Println(summary)
	result = nil
	return result
}

func calculateNamespaces(config DB2PDLConfig) (string, string) {
	namespace := strings.TrimSpace(config.PDL.EntitiesNamespace)
	if namespace == "" {
		panic("db2Pdl.pdl.entitiesNamespace is required")
	}
	return namespace, NamespaceToPHP(namespace)
}
