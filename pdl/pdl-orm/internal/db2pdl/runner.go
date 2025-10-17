package db2pdl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	cpptarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/cpp"
	csharptarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/csharp"
	gotarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/go"
	javatarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/java"
	kotlintarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/kotlin"
	pdltarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/pdl"
	phtarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/php"
	rusttarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/rust"
	tstarget "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator/typescript"
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
	writer := func(language string, fileName string, source string) error {
		return runner.outputCode(source, fileName, language, dbConfig.OutputDir, &stats, verbosePrinter)
	}
	pdlGenerator := pdltarget.New(templateUtils, writer)
	phpGenerator := phtarget.New(templateUtils, writer)
	tsGenerator := tstarget.New(templateUtils, writer)
	goGenerator := gotarget.New(templateUtils, writer, dbConfig)
	csGenerator := csharptarget.New(templateUtils, writer)
	cppGenerator := cpptarget.New(templateUtils, writer)
	javaGenerator := javatarget.New(templateUtils, writer)
	kotlinGenerator := kotlintarget.New(templateUtils, writer)
	rustGenerator := rusttarget.New(templateUtils, writer)
	tables, tablesErr := driver.ListTables(ctx, sqlConnection)
	if tablesErr != nil {
		return tablesErr
	}
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
		if err := pdlGenerator.Generate(tableData); err != nil {
			return err
		}
		if dbConfig.PHP.EmitHelpers {
			if err := phpGenerator.Generate(tableData); err != nil {
				return err
			}
		}
		if dbConfig.TypeScript.Emit {
			if err := tsGenerator.AddTable(tableData); err != nil {
				return err
			}
		}
		if dbConfig.Go.Emit {
			if err := goGenerator.Generate(tableData); err != nil {
				return err
			}
		}
		if dbConfig.Java.Emit {
			if err := javaGenerator.Generate(tableData); err != nil {
				return err
			}
		}
		if dbConfig.Kotlin.Emit {
			if err := kotlinGenerator.Generate(tableData); err != nil {
				return err
			}
		}
		if dbConfig.CSharp.Emit {
			if err := csGenerator.Generate(tableData); err != nil {
				return err
			}
		}
		if dbConfig.Cpp.Emit {
			if err := cppGenerator.Generate(tableData); err != nil {
				return err
			}
		}
		if dbConfig.Rust.Emit {
			if err := rustGenerator.Generate(tableData); err != nil {
				return err
			}
		}
	}
	finishErr := runner.finishGeneration(dbConfig, tsGenerator, &stats, verbosePrinter)
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

func (runner Runner) finishGeneration(mysqlConfig DB2PDLConfig, tsGenerator *tstarget.Generator, stats *executionStats, verbosePrinter VerbosePrinter) error {
	if mysqlConfig.TypeScript.Emit && tsGenerator != nil {
		if err := tsGenerator.Flush(mysqlConfig.TypeScript.OutputFile); err != nil {
			return err
		}
	}
	summary := fmt.Sprintf("Stats: %s files generated, %s lines generated, %s bytes", strconv.Itoa(stats.totalFiles), strconv.Itoa(stats.totalLines), strconv.FormatInt(stats.totalSize, 10))
	fmt.Printf("Output dir: %s\n", mysqlConfig.OutputDir)
	fmt.Println(summary)
	return nil
}

func calculateNamespaces(config DB2PDLConfig) (string, string) {
	namespace := strings.TrimSpace(config.PDL.EntitiesNamespace)
	if namespace == "" {
		panic("db2Pdl.pdl.entitiesNamespace is required")
	}
	return namespace, NamespaceToPHP(namespace)
}
