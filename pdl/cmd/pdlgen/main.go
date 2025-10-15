package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kapablanka/pdl/pdl/internal/generator"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/cpp"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/csharp"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/golang"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/java"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/javascript"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/kotlin"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/php"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/rust"
	_ "github.com/kapablanka/pdl/pdl/internal/generator/typescript"
)

func main() {
	options, parseErr := parseOptions()
	if parseErr != nil {
		fmt.Fprintln(os.Stderr, parseErr)
		os.Exit(1)
	}
	runErr := runGenerator(options)
	if runErr != nil {
		fmt.Fprintln(os.Stderr, runErr)
		os.Exit(1)
	}
}

func parseOptions() (generator.Options, error) {
	var result generator.Options
	generatorFlag := flag.String("generator", "", "generator name")
	astFlag := flag.String("ast", "", "path to AST document")
	outputFlag := flag.String("output", "", "output directory")
	templateFlag := flag.String("templates", "", "templates directory")
	configFlag := flag.String("config", "", "generator configuration file")
	flag.Parse()
	result = generator.Options{
		GeneratorName: *generatorFlag,
		ASTPath:       *astFlag,
		OutputDir:     *outputFlag,
		TemplateDir:   *templateFlag,
		ConfigPath:    *configFlag,
	}
	validateErr := validateOptions(result)
	if validateErr != nil {
		return result, validateErr
	}
	return result, nil
}

func validateOptions(options generator.Options) error {
	if options.GeneratorName == "" {
		return fmt.Errorf("--generator is required")
	}
	if options.ASTPath == "" {
		return fmt.Errorf("--ast is required")
	}
	return nil
}

func runGenerator(options generator.Options) error {
	instance, createErr := generator.Create(options)
	if createErr != nil {
		return createErr
	}
	return instance.Generate()
}
