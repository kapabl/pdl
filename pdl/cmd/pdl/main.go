package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/initcmd"
	"github.com/kapablanka/pdl/pdl/internal/pipeline"
	"github.com/kapablanka/pdl/pdl/internal/runner"
)

type stringFlag struct {
	value string
	set   bool
}

func (flagValue *stringFlag) Set(value string) error {
	flagValue.value = value
	flagValue.set = true
	return nil
}

func (flagValue *stringFlag) String() string {
	return flagValue.value
}

func main() {
	configFlag := &stringFlag{value: "pdl.config.json"}
	flag.Var(configFlag, "config", "configuration file")
	initFlag := flag.Bool("init", false, "initialize a PDL workspace in the current directory")
	initDirFlag := flag.String("dir", "", "target directory for --init (defaults to ./pdl)")
	buildFlag := flag.Bool("build", false, "run the full build pipeline (default)")
	rebuildFlag := flag.Bool("rebuild", false, "rebuild generated outputs and refresh fanout targets")
	cleanFlag := flag.Bool("clean", false, "clean generated outputs and staged directories")
	db2pdlFlag := flag.Bool("db2pdl", false, "run db2pdl stage only")
	compileFlag := flag.Bool("compile", false, "run the PDL compiler only")
	cleanStageFlag := flag.Bool("clean-stage", false, "clean staged directories before copying")
	verboseFlag := flag.Bool("verbose", false, "enable verbose logging")
	flag.Parse()

	if *initFlag {
		options, promptErr := collectInitOptions(*initDirFlag)
		if promptErr != nil {
			fmt.Fprintf(os.Stderr, "init cancelled: %v\n", promptErr)
			os.Exit(1)
		}
		result, initErr := initcmd.Run(options)
		if initErr != nil {
			fmt.Fprintf(os.Stderr, "init failed: %v\n", initErr)
			os.Exit(1)
		}
		fmt.Print(initcmd.FormatResult(result))
		os.Exit(0)
	}

	configPath := resolveConfigPath(configFlag)
	loadEnvFiles(configPath)

	if *compileFlag {
		if *cleanFlag || *db2pdlFlag || *rebuildFlag {
			fmt.Fprintln(os.Stderr, "--compile cannot be combined with --clean, --db2pdl, or --rebuild")
			os.Exit(1)
		}
		runr := runner.Runner{ConfigPath: configPath, Rebuild: false, Verbose: *verboseFlag}
		exitCode, err := runr.Run(context.Background())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(exitCode)
	}

	action := pipeline.ActionBuild
	if *cleanFlag {
		if *rebuildFlag || *db2pdlFlag {
			fmt.Fprintln(os.Stderr, "--clean cannot be combined with --rebuild or --db2pdl")
			os.Exit(1)
		}
		action = pipeline.ActionClean
	} else if *rebuildFlag {
		if *db2pdlFlag {
			action = pipeline.ActionRebuildAll
		} else {
			action = pipeline.ActionRebuild
		}
	} else if *db2pdlFlag {
		action = pipeline.ActionDb2Pdl
	} else if *buildFlag {
		action = pipeline.ActionBuild
	}

	opts := pipeline.Options{
		ConfigPath: configPath,
		CleanStage: *cleanStageFlag,
		Action:     action,
		Verbose:    *verboseFlag,
	}

	if err := pipeline.Run(context.Background(), opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func resolveConfigPath(configFlag *stringFlag) string {
	path := configFlag.value
	if path == "" {
		path = "pdl.config.json"
	}
	if configFlag.set {
		return path
	}
	if fileExists(path) {
		return path
	}
	workspacePath := filepath.Join("pdl", "pdl.config.json")
	if fileExists(workspacePath) {
		return workspacePath
	}
	return path
}

func loadEnvFiles(configPath string) {
	configDir := filepath.Dir(configPath)
	candidates := []string{
		filepath.Join(configDir, ".env"),
		filepath.Join(configDir, ".env.local"),
	}
	for _, candidate := range candidates {
		if err := loadEnvFile(candidate); err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to load %s: %v\n", candidate, err)
		}
	}
}

func loadEnvFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if key == "" {
			continue
		}
		if _, exists := os.LookupEnv(key); exists {
			continue
		}
		_ = os.Setenv(key, os.ExpandEnv(value))
	}
	return scanner.Err()
}

func fileExists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}
