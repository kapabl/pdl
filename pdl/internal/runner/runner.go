package runner

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bmatcuk/doublestar/v4"

	"github.com/kapablanka/pdl/pdl/internal/config"
	"github.com/kapablanka/pdl/pdl/internal/filetimes"
	"github.com/kapablanka/pdl/pdl/internal/jsprocessor"
	"github.com/kapablanka/pdl/pdl/internal/utils"
)

const compilerBinary = "pdlc2"

type Runner struct {
	ConfigPath         string
	Rebuild            bool
	Verbose            bool
	CompilerConfigPath string
}

type sectionContext struct {
	sectionName string
	profileName string
}

type compileStats struct {
	childInstances int
}

func (runner Runner) Run(ctx context.Context) (int, error) {
	var result int
	cfg, loadErr := config.Load(runner.ConfigPath)
	if loadErr != nil {
		return result, loadErr
	}
	effectiveVerbose := cfg.Verbose || runner.Verbose
	verbosePrinter := utils.NewVerbosePrinter(effectiveVerbose)
	binaryDir, dirErr := runner.resolveBinaryDir(cfg)
	if dirErr != nil {
		return result, dirErr
	}
	compilerPath := filepath.Join(binaryDir, compilerBinary)
	info, statErr := os.Stat(compilerPath)
	if statErr != nil || info.IsDir() {
		return result, fmt.Errorf("compiler not found at %s", compilerPath)
	}
	rebuild := runner.Rebuild
	if rebuild {
		tempErr := utils.CleanDir(cfg.TempDir)
		if tempErr != nil {
			return result, tempErr
		}
		outErr := utils.CleanDir(cfg.OutputDir)
		if outErr != nil {
			return result, outErr
		}
	}
	fileTimesPath := filepath.Join(cfg.TempDir, "pdlFileTimes.json")
	times := filetimes.NewFileTimes(fileTimesPath)
	_ = times.Read()
	fmt.Println(">>>>Begin PDL compilation and generation")
	stats := compileStats{}
	errorAccumulator := &multiError{}
	runner.compileSections(ctx, cfg, compilerPath, verbosePrinter, &times, &stats, errorAccumulator)
	writeErr := times.Write()
	if writeErr != nil {
		errorAccumulator.Append(writeErr)
	}
	fmt.Printf("pdl files compiled: %d file(s)\n", stats.childInstances)
	fmt.Println("<<<<End PDL compilation and generation")
	if errorAccumulator.HasErrors() {
		return result, errorAccumulator.Error()
	}
	fmt.Println(">>>>Begin JS processing")
	jsErr := jsprocessor.Run(cfg, verbosePrinter)
	if jsErr != nil {
		return result, jsErr
	}
	fmt.Println("<<<<END JS processing")
	return result, nil
}

func (runner Runner) resolveBinaryDir(cfg config.RootConfig) (string, error) {
	var result string
	baseDir := strings.TrimSpace(os.Getenv("PDL_BIN_PATH"))
	if baseDir == "" {
		executablePath, execErr := os.Executable()
		if execErr != nil {
			return result, fmt.Errorf("unable to locate CLI directory: %w", execErr)
		}
		baseDir = filepath.Dir(executablePath)
	}
	if absBase, absErr := filepath.Abs(baseDir); absErr == nil {
		baseDir = absBase
	}
	if info, statErr := os.Stat(baseDir); statErr != nil || !info.IsDir() {
		return result, fmt.Errorf("binary directory not found: %s", baseDir)
	}
	result = baseDir
	return result, nil
}

func (runner Runner) compileSections(ctx context.Context, cfg config.RootConfig, compilerPath string, printer utils.VerbosePrinter, times *filetimes.FileTimes, stats *compileStats, accumulator *multiError) {
	var wg sync.WaitGroup
	sectionChan := make(chan sectionContext)
	workerCount := 1
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range sectionChan {
				profile := cfg.Profiles[item.profileName]
				if !profile.Enabled {
					continue
				}
				runner.compileSectionProfile(ctx, cfg, item.sectionName, profile, item.profileName, compilerPath, printer, times, stats, accumulator)
			}
		}()
	}
	for _, section := range cfg.Sections {
		if !runner.isSectionEnabled(section.Name, cfg.ActiveSections) {
			continue
		}
		if !runner.sectionHasProfiles(section, cfg.Profiles) {
			continue
		}
		fmt.Printf("Compiling %s...\n", section.Name)
		for profileName := range section.Files {
			if strings.HasSuffix(profileName, "Exclude") {
				continue
			}
			if _, ok := cfg.Profiles[profileName]; !ok {
				accumulator.Append(fmt.Errorf("invalid profile: %s", profileName))
				continue
			}
			if profile := cfg.Profiles[profileName]; !profile.Enabled {
				continue
			}
			sectionChan <- sectionContext{sectionName: section.Name, profileName: profileName}
		}
	}
	close(sectionChan)
	wg.Wait()
}

func (runner Runner) isSectionEnabled(name string, allowed map[string]struct{}) bool {
	if len(allowed) == 0 {
		return true
	}
	_, ok := allowed[name]
	return ok
}

func (runner Runner) sectionHasProfiles(section config.Section, profiles map[string]config.Profile) bool {
	for profileKey := range section.Files {
		if strings.HasSuffix(profileKey, "Exclude") {
			continue
		}
		if _, exists := profiles[profileKey]; exists {
			return true
		}
	}
	return false
}

func (runner Runner) compileSectionProfile(ctx context.Context, cfg config.RootConfig, sectionName string, profile config.Profile, profileName string, compilerPath string, printer utils.VerbosePrinter, times *filetimes.FileTimes, stats *compileStats, accumulator *multiError) {
	section, err := findSection(cfg.Sections, sectionName)
	if err != nil {
		accumulator.Append(err)
		return
	}
	if !profile.Enabled {
		return
	}
	templatesDir := cfg.Templates.Dir
	if templatesDir == "" {
		templatesDir = profile.Templates.Dir
	}
	templateName := cfg.Templates.Name
	if templateName == "" {
		templateName = profile.Templates.Name
	}
	if templateName == "" {
		templateName = "classTemplate1"
	}
	if section.OutputDir != "" && profile.OutputDir == "" {
		profile.OutputDir = section.OutputDir
	}
	if profile.OutputDir == "" {
		profile.OutputDir = cfg.OutputDir
	}
	arguments := []string{templatesDir, templateName, profile.OutputDir}
	if runner.CompilerConfigPath != "" {
		arguments = append(arguments, runner.CompilerConfigPath)
	}
	exclusions := runner.collectExclusions(section, profileName, profile, cfg)
	files := runner.resolveFiles(section, profileName, profile, cfg, exclusions, accumulator)
	for _, file := range files {
		processErr := runner.processFile(ctx, file, arguments, compilerPath, printer, times, stats)
		if processErr != nil {
			accumulator.Append(processErr)
		}
	}
}

func (runner Runner) processFile(ctx context.Context, file string, arguments []string, compilerPath string, printer utils.VerbosePrinter, times *filetimes.FileTimes, stats *compileStats) error {
	_, statErr := os.Stat(file)
	if statErr != nil {
		return statErr
	}
	modified, modErr := times.IsFileModified(file)
	if modErr != nil {
		return modErr
	}
	if !modified {
		printer.Println(file + " not modified")
		return nil
	}
	expanded := runner.expandArguments(file, arguments)
	printer.Println(compilerPath + " " + strings.Join(expanded, " "))
	cmd := exec.CommandContext(ctx, compilerPath, expanded...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	stats.childInstances++
	_ = times.AddFile(file)
	return nil
}

func (runner Runner) expandArguments(file string, arguments []string) []string {
	result := make([]string, 0, len(arguments)+1)
	result = append(result, file)
	for _, arg := range arguments {
		value := strings.ReplaceAll(arg, "[inputFile]", file)
		result = append(result, value)
	}
	return result
}

func (runner Runner) resolveFiles(section config.Section, profileName string, profile config.Profile, cfg config.RootConfig, exclusions map[string]struct{}, accumulator *multiError) []string {
	result := make([]string, 0)
	sourcePaths := mergeStringSlices(section.Src, profile.Src, cfg.Src)
	patterns, ok := section.Files[profileName]
	if !ok {
		return result
	}
	for _, pattern := range patterns {
		for _, source := range sourcePaths {
			globPattern := filepath.Join(source, pattern)
			convertedPattern := toRelativeGlob(globPattern)
			matches, globErr := doublestar.Glob(os.DirFS("."), convertedPattern)
			if globErr != nil {
				accumulator.Append(globErr)
				continue
			}
			for _, match := range matches {
				clean := filepath.Clean(match)
				if _, excluded := exclusions[clean]; excluded {
					continue
				}
				abs, absErr := filepath.Abs(clean)
				if absErr != nil {
					accumulator.Append(absErr)
					continue
				}
				result = append(result, abs)
			}
		}
	}
	return result
}

func (runner Runner) collectExclusions(section config.Section, profileName string, profile config.Profile, cfg config.RootConfig) map[string]struct{} {
	result := make(map[string]struct{})
	sourcePaths := mergeStringSlices(section.Src, profile.Src, cfg.Src)
	excludeKey := profileName + "Exclude"
	patterns, ok := section.Files[excludeKey]
	if !ok {
		return result
	}
	for _, source := range sourcePaths {
		for _, pattern := range patterns {
			globPattern := filepath.Join(source, pattern)
			convertedPattern := toRelativeGlob(globPattern)
			matches, _ := doublestar.Glob(os.DirFS("."), convertedPattern)
			for _, match := range matches {
				abs, absErr := filepath.Abs(match)
				if absErr != nil {
					continue
				}
				result[filepath.Clean(abs)] = struct{}{}
			}
		}
	}
	return result
}

func mergeStringSlices(slices ...[]string) []string {
	result := make([]string, 0)
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func toRelativeGlob(pattern string) string {
	cleaned := filepath.ToSlash(pattern)
	if strings.HasPrefix(cleaned, "./") {
		cleaned = cleaned[2:]
	}
	if filepath.IsAbs(pattern) {
		relative, relErr := filepath.Rel(".", pattern)
		if relErr == nil {
			cleaned = filepath.ToSlash(relative)
		}
	}
	return cleaned
}

func findSection(sections []config.Section, name string) (config.Section, error) {
	for _, section := range sections {
		if section.Name == name {
			result := section
			return result, nil
		}
	}
	return config.Section{}, fmt.Errorf("section not found: %s", name)
}

type multiError struct {
	errors []error
	mutex  sync.Mutex
}

func (collector *multiError) Append(err error) {
	if err == nil {
		return
	}
	collector.mutex.Lock()
	defer collector.mutex.Unlock()
	collector.errors = append(collector.errors, err)
}

func (collector *multiError) HasErrors() bool {
	collector.mutex.Lock()
	defer collector.mutex.Unlock()
	return len(collector.errors) > 0
}

func (collector *multiError) Error() error {
	collector.mutex.Lock()
	defer collector.mutex.Unlock()
	if len(collector.errors) == 0 {
		return nil
	}
	messages := make([]string, len(collector.errors))
	for index, err := range collector.errors {
		messages[index] = err.Error()
	}
	result := errors.New(strings.Join(messages, "; "))
	return result
}
