package pipeline

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/bmatcuk/doublestar/v4"

	"github.com/kapablanka/pdl/pdl/internal/config"
	"github.com/kapablanka/pdl/pdl/internal/defaulttemplates"
	"github.com/kapablanka/pdl/pdl/internal/runner"
	"github.com/kapablanka/pdl/pdl/internal/utils"
)

type Action int

const (
	ActionBuild Action = iota
	ActionRebuild
	ActionRebuildAll
	ActionClean
	ActionDb2Pdl
)

type Options struct {
	ConfigPath string
	CleanStage bool
	Action     Action
	Verbose    bool
}

type runParameters struct {
	runDb2Pdl   bool
	rebuild     bool
	cleanFanout bool
}

type buildSpec struct {
	projectRoot        string
	configPath         string
	outputDir          string
	tempDir            string
	astOutputDir       string
	compilerConfigPath string
	db2PdlOutputDir    string
	db2PdlStageDir     string
	jsCopyTarget       string
	phpCopyTarget      string
	bundleCopyTarget   string
	phpDbClassesTarget string
	goCopyTarget       string
	rootConfig         config.RootConfig
}

type pipelineExecutor struct {
	ctx        context.Context
	spec       buildSpec
	printer    utils.VerbosePrinter
	verbose    bool
	cleanStage bool
}

type copySpec struct {
	pattern     string
	stripPrefix string
	destination string
	cleanDest   bool
	flatten     bool
}

type matchDetails struct {
	spec   copySpec
	match  string
	prefix string
	info   os.FileInfo
}

func Run(ctx context.Context, options Options) error {
	if options.ConfigPath == "" {
		options.ConfigPath = "pdl.config.json"
	}
	spec, err := resolveSpec(options)
	if err != nil {
		return err
	}
	printer := utils.NewVerbosePrinter(options.Verbose)
	printer.Println(fmt.Sprintf("pipeline: action=%d config=%s", options.Action, spec.configPath))
	executor := pipelineExecutor{
		ctx:        ctx,
		spec:       spec,
		printer:    printer,
		verbose:    options.Verbose,
		cleanStage: options.CleanStage,
	}
	switch options.Action {
	case ActionDb2Pdl:
		return executor.runDb2Pdl()
	case ActionClean:
		return executor.cleanOutputs()
	case ActionRebuild:
		params := runParameters{runDb2Pdl: false, rebuild: true, cleanFanout: true}
		return executor.runBuild(params)
	case ActionRebuildAll:
		params := runParameters{runDb2Pdl: true, rebuild: true, cleanFanout: true}
		return executor.runBuild(params)
	case ActionBuild:
		fallthrough
	default:
		params := runParameters{runDb2Pdl: false, rebuild: false, cleanFanout: false}
		return executor.runBuild(params)
	}
}

func (executor pipelineExecutor) cleanOutputs() error {
	var result error
	executor.printer.Println("pipeline: cleaning outputs")
	if err := executor.fanoutOutputs(true); err != nil {
		result = err
		return result
	}
	if executor.spec.outputDir != "" {
		if err := utils.CleanDir(executor.spec.outputDir); err != nil {
			result = err
			return result
		}
	}
	if executor.spec.db2PdlOutputDir != "" {
		if err := utils.CleanDir(executor.spec.db2PdlOutputDir); err != nil {
			result = err
			return result
		}
	}
	if executor.cleanStage && executor.spec.db2PdlStageDir != "" {
		if err := utils.CleanDir(executor.spec.db2PdlStageDir); err != nil {
			result = err
			return result
		}
	}
	return result
}

func resolveSpec(options Options) (buildSpec, error) {
	var result buildSpec
	configAbs, absErr := filepath.Abs(options.ConfigPath)
	if absErr != nil {
		return result, absErr
	}
	cfg, cfgErr := config.Load(configAbs)
	if cfgErr != nil {
		return result, cfgErr
	}
	projectRoot := filepath.Dir(configAbs)
	templateDir := os.Getenv("PDL_TEMPLATES_DIR")
	if templateDir == "" {
		materialized, materializeErr := defaulttemplates.MaterializeInto(projectRoot)
		if materializeErr == nil {
			_ = os.Setenv("PDL_TEMPLATES_DIR", materialized)
			templateDir = materialized
		}
	}
	_ = templateDir

	outputDir := resolveProjectPath(projectRoot, cfg.OutputDir)
	if outputDir == "" {
		outputDir = filepath.Join(projectRoot, "output")
	}
	tempDir := resolveProjectPath(projectRoot, cfg.TempDir)
	if tempDir == "" {
		tempDir = filepath.Join(outputDir, "temp")
	}
	astDir := filepath.Join(outputDir, "ast")
	db2pdlOutput := ""
	if cfg.Db2Pdl.Enabled {
		db2pdlOutput = resolveProjectPath(projectRoot, cfg.Db2Pdl.OutputDir)
		if db2pdlOutput == "" {
			db2pdlOutput = resolveProjectPath(projectRoot, os.Getenv("PDL_DB2PDL_OUTPUT"))
		}
		if db2pdlOutput == "" {
			db2pdlOutput = filepath.Join(outputDir, "pdl")
		}
	}
	stageDir := ""
	if cfg.Db2Pdl.Enabled {
		stageDir = filepath.Join(projectRoot, "src")
		if cfg.Db2Pdl.Db2PdlSourceDest != "" {
			stageDir = filepath.Join(stageDir, filepath.FromSlash(cfg.Db2Pdl.Db2PdlSourceDest))
		}
	}
	jsTarget := resolveProjectPath(projectRoot, os.Getenv("PDL_GEN_OUTPUT_JS"))
	phpTarget := resolveProjectPath(projectRoot, os.Getenv("PDL_GEN_OUTPUT_PHP"))
	bundleTarget := resolveProjectPath(projectRoot, os.Getenv("PDL_GEN_OUTPUT_BUNDLE"))
	phpDbTarget := ""
	if cfg.Db2Pdl.Enabled && phpTarget != "" {
		if cfg.Db2Pdl.Db2PdlSourceDest != "" {
			phpDbTarget = filepath.Join(phpTarget, titleNamespace(cfg.Db2Pdl.Db2PdlSourceDest))
		} else {
			phpDbTarget = phpTarget
		}
	}
	compilerConfigPath, configErr := writeCompilerConfig(projectRoot, cfg)
	if configErr != nil {
		return result, configErr
	}
	result = buildSpec{
		projectRoot:        projectRoot,
		configPath:         configAbs,
		outputDir:          outputDir,
		tempDir:            tempDir,
		astOutputDir:       astDir,
		compilerConfigPath: compilerConfigPath,
		db2PdlOutputDir:    db2pdlOutput,
		db2PdlStageDir:     stageDir,
		jsCopyTarget:       jsTarget,
		phpCopyTarget:      phpTarget,
		bundleCopyTarget:   bundleTarget,
		phpDbClassesTarget: phpDbTarget,
		goCopyTarget:       resolveProjectPath(projectRoot, os.Getenv("PDL_GEN_OUTPUT_GO")),
		rootConfig:         cfg,
	}
	return result, nil
}

func (executor pipelineExecutor) runBuild(params runParameters) error {
	var result error
	executor.printer.Println(fmt.Sprintf("pipeline: runBuild db2pdl=%t rebuild=%t cleanFanout=%t", params.runDb2Pdl, params.rebuild, params.cleanFanout))
	ormEnabled := executor.spec.rootConfig.Db2Pdl.Enabled
	if params.runDb2Pdl && ormEnabled {
		executor.printer.Println("pipeline: executing db2pdl stage")
		dbErr := executor.runDb2Pdl()
		if dbErr != nil {
			result = dbErr
			return result
		}
	}
	if ormEnabled && executor.spec.db2PdlOutputDir != "" && executor.spec.db2PdlStageDir != "" {
		stageErr := executor.stageDb2Pdl()
		if stageErr != nil {
			result = stageErr
			return result
		}
	}
	executor.printer.Println("pipeline: invoking compiler")
	compilerErr := executor.runPdlCompiler(params.rebuild)
	if compilerErr != nil {
		result = compilerErr
		return result
	}
	astErr := executor.runAstGenerators()
	if astErr != nil {
		result = astErr
		return result
	}
	syncErr := executor.fanoutOutputs(params.cleanFanout)
	if syncErr != nil {
		result = syncErr
		return result
	}
	return result
}

func (executor pipelineExecutor) runDb2Pdl() error {
	var result error
	binaryArgs := []string{"--run", "--exit", "--config", executor.spec.configPath}
	if binaryPath, ok := resolveBinaryPath("db2pdl"); ok {
		executor.printer.Println("pipeline: using prebuilt db2pdl binary")
		result = executor.runCommand(executor.spec.projectRoot, binaryPath, binaryArgs)
		return result
	}
	return fmt.Errorf("db2pdl binary not found; ensure it is on PATH or set PDL_BIN_PATH")
}

func (executor pipelineExecutor) runPdlCompiler(rebuild bool) error {
	var result error
	binaryArgs := []string{"--config", executor.spec.configPath}
	if rebuild {
		binaryArgs = append(binaryArgs, "--rebuild")
	}
	if binaryPath, ok := resolveBinaryPath("pdl"); ok {
		if !sameExecutable(binaryPath) {
			executor.printer.Println("pipeline: using prebuilt pdl binary")
			result = executor.runCommand(executor.spec.projectRoot, binaryPath, binaryArgs)
			return result
		}
		executor.printer.Println("pipeline: detected self-invocation, running compiler inline")
	}
	compiler := runner.Runner{
		ConfigPath:         executor.spec.configPath,
		CompilerConfigPath: executor.spec.compilerConfigPath,
		Rebuild:            rebuild,
		Verbose:            executor.verbose,
	}
	_, runErr := compiler.Run(executor.ctx)
	if runErr != nil {
		result = runErr
		return result
	}
	return result
}

func (executor pipelineExecutor) stageDb2Pdl() error {
	executor.printer.Println("pipeline: staging db2pdl outputs")
	if err := executor.prepareStageDir(); err != nil {
		return err
	}
	pattern := filepath.Join(executor.spec.db2PdlOutputDir, "pdl", "*.pdl")
	copies, err := doublestar.FilepathGlob(pattern)
	if err != nil {
		return err
	}
	return executor.copyStageFiles(copies)
}

func (executor pipelineExecutor) prepareStageDir() error {
	if executor.cleanStage {
		return utils.CleanDir(executor.spec.db2PdlStageDir)
	}
	return utils.EnsureDir(executor.spec.db2PdlStageDir)
}

func (executor pipelineExecutor) copyStageFiles(paths []string) error {
	for _, src := range paths {
		info, err := os.Stat(src)
		if err != nil {
			return err
		}
		if info.IsDir() {
			continue
		}
		target := filepath.Join(executor.spec.db2PdlStageDir, filepath.Base(src))
		if err := copyFile(src, target); err != nil {
			return err
		}
	}
	return nil
}

func (executor pipelineExecutor) fanoutOutputs(cleanTargets bool) error {
	var result error
	specs := executor.buildCopySpecs(cleanTargets)
	for _, item := range specs {
		if err := executor.syncPattern(item); err != nil {
			result = err
			return result
		}
	}
	return result
}

func (executor pipelineExecutor) buildCopySpecs(cleanTargets bool) []copySpec {
	specs := make([]copySpec, 0, 4)
	addSpec := func(pattern string, strip string, destination string, flatten bool) {
		specs = append(specs, copySpec{
			pattern:     pattern,
			stripPrefix: strip,
			destination: destination,
			cleanDest:   cleanTargets,
			flatten:     flatten,
		})
	}
	if executor.spec.jsCopyTarget != "" {
		base := filepath.Join(executor.spec.outputDir, "js")
		if !strings.EqualFold(filepath.Clean(base), filepath.Clean(executor.spec.jsCopyTarget)) {
			addSpec(filepath.Join(base, "**", "*.js"), base, executor.spec.jsCopyTarget, false)
		}
	}
	if executor.spec.phpCopyTarget != "" {
		base := filepath.Join(executor.spec.outputDir, "php")
		if !strings.EqualFold(filepath.Clean(base), filepath.Clean(executor.spec.phpCopyTarget)) {
			addSpec(filepath.Join(base, "**", "*.php"), base, executor.spec.phpCopyTarget, false)
		}
		if executor.spec.db2PdlOutputDir != "" && executor.spec.phpDbClassesTarget != "" {
			dbBase := filepath.Join(executor.spec.db2PdlOutputDir, "php")
			addSpec(filepath.Join(dbBase, "**", "*.php"), dbBase, executor.spec.phpDbClassesTarget, true)
		}
	}
	if executor.spec.goCopyTarget != "" {
		base := filepath.Join(executor.spec.outputDir, "go")
		if !strings.EqualFold(filepath.Clean(base), filepath.Clean(executor.spec.goCopyTarget)) {
			addSpec(filepath.Join(base, "**", "*.go"), base, executor.spec.goCopyTarget, false)
		}
	}
	return specs
}
func (executor pipelineExecutor) syncPattern(spec copySpec) error {
	if spec.destination == "" {
		return nil
	}
	matches, err := doublestar.FilepathGlob(spec.pattern)
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		if spec.cleanDest && spec.destination != "" {
			return os.RemoveAll(spec.destination)
		}
		return nil
	}
	executor.printer.Println(fmt.Sprintf("pipeline: syncing %s -> %s", spec.pattern, spec.destination))
	if err := executor.prepareDestination(spec); err != nil {
		return err
	}
	prefix := withTrailingSeparator(filepath.Clean(spec.stripPrefix))
	for _, match := range matches {
		if err := executor.copyMatchedFile(spec, match, prefix); err != nil {
			return err
		}
	}
	return nil
}

func (executor pipelineExecutor) prepareDestination(spec copySpec) error {
	if spec.cleanDest {
		if err := os.RemoveAll(spec.destination); err != nil {
			return err
		}
	}
	return utils.EnsureDir(spec.destination)
}

func (executor pipelineExecutor) copyMatchedFile(spec copySpec, match string, prefix string) error {
	info, err := os.Stat(match)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	relative := relativePath(matchDetails{spec: spec, match: match, prefix: prefix, info: info})
	destinationPath := filepath.Join(spec.destination, relative)
	if err := utils.EnsureDir(filepath.Dir(destinationPath)); err != nil {
		return err
	}
	return copyFile(match, destinationPath)
}

func relativePath(details matchDetails) string {
	if details.spec.flatten {
		return details.info.Name()
	}
	if details.prefix != "" && strings.HasPrefix(details.match, details.prefix) {
		return strings.TrimPrefix(details.match, details.prefix)
	}
	return filepath.Base(details.match)
}

func withTrailingSeparator(value string) string {
	if value == "" {
		return value
	}
	if strings.HasSuffix(value, string(filepath.Separator)) {
		return value
	}
	return value + string(filepath.Separator)
}

func copyFile(src string, dest string) error {
	var result error
	sourceFile, openErr := os.Open(src)
	if openErr != nil {
		result = openErr
		return result
	}
	defer sourceFile.Close()
	targetFile, createErr := os.Create(dest)
	if createErr != nil {
		result = createErr
		return result
	}
	defer targetFile.Close()
	_, copyErr := io.Copy(targetFile, sourceFile)
	if copyErr != nil {
		result = copyErr
		return result
	}
	syncErr := targetFile.Sync()
	if syncErr != nil {
		result = syncErr
		return result
	}
	chErr := os.Chmod(dest, 0o644)
	if chErr != nil {
		result = chErr
		return result
	}
	return result
}

func titleNamespace(value string) string {
	if value == "" {
		return ""
	}
	parts := strings.Split(value, "/")
	for index, part := range parts {
		parts[index] = titleWord(part)
	}
	result := filepath.Join(parts...)
	return result
}

func titleWord(input string) string {
	if input == "" {
		return ""
	}
	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])
	result := string(runes)
	return result
}

func resolveProjectPath(projectRoot string, value string) string {
	result := value
	if value == "" {
		return result
	}
	if filepath.IsAbs(value) {
		result = value
		return result
	}
	result = filepath.Join(projectRoot, value)
	return result
}

func writeCompilerConfig(projectRoot string, cfg config.RootConfig) (string, error) {
	languages := buildLanguageConfig(cfg)
	if len(languages) == 0 {
		return "", nil
	}
	targetDir := filepath.Join(projectRoot, ".pdl")
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return "", err
	}
	payload := map[string]interface{}{
		"file": languages,
	}
	bytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "", err
	}
	path := filepath.Join(targetDir, "compiler.config.json")
	if err := os.WriteFile(path, bytes, 0o644); err != nil {
		return "", err
	}
	return path, nil
}

func buildLanguageConfig(cfg config.RootConfig) map[string]interface{} {
	languages := make(map[string]map[string]interface{})
	for name, profile := range cfg.Profiles {
		settings := profile.ConfigSettings()
		if _, ok := settings["enabled"]; !ok {
			settings["enabled"] = profile.Enabled
		}
		language := profile.LanguageOrDefault(name)
		existing, exists := languages[language]
		if !exists {
			languages[language] = settings
			continue
		}
		if existingEnabled, ok := existing["enabled"].(bool); ok {
			existing["enabled"] = existingEnabled || profile.Enabled
		} else {
			existing["enabled"] = profile.Enabled
		}
		for key, value := range settings {
			if key == "enabled" {
				continue
			}
			existing[key] = value
		}
	}
	result := make(map[string]interface{}, len(languages))
	for language, payload := range languages {
		if _, ok := payload["enabled"]; !ok {
			payload["enabled"] = false
		}
		result[language] = payload
	}
	return result
}

func (executor pipelineExecutor) runCommand(workingDir string, command string, arguments []string) error {
	var result error
	executor.printer.Println(fmt.Sprintf("pipeline: exec %s %s (cwd=%s)", command, strings.Join(arguments, " "), workingDir))
	cmd := exec.CommandContext(executor.ctx, command, arguments...)
	cmd.Dir = workingDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	result = cmd.Run()
	return result
}

func resolveBinaryPath(name string) (string, bool) {
	binRoot := os.Getenv("PDL_BIN_PATH")
	if binRoot != "" {
		candidates := []string{filepath.Join(binRoot, name)}
		if runtime.GOOS == "windows" {
			candidates = append(candidates, filepath.Join(binRoot, name+".exe"))
		} else {
			candidates = append(candidates, filepath.Join(binRoot, name+".exe"))
		}
		for _, candidate := range candidates {
			info, err := os.Stat(candidate)
			if err == nil && !info.IsDir() {
				return candidate, true
			}
		}
	}
	if self, err := os.Executable(); err == nil {
		selfDir := filepath.Dir(self)
		candidates := []string{filepath.Join(selfDir, name)}
		if runtime.GOOS == "windows" {
			candidates = append(candidates, filepath.Join(selfDir, name+".exe"))
		} else {
			candidates = append(candidates, filepath.Join(selfDir, name+".exe"))
		}
		for _, candidate := range candidates {
			info, err := os.Stat(candidate)
			if err == nil && !info.IsDir() {
				return candidate, true
			}
		}
	}
	if resolved, err := exec.LookPath(name); err == nil {
		return resolved, true
	}
	return "", false
}

func sameExecutable(candidate string) bool {
	selfPath, err := os.Executable()
	if err != nil {
		return false
	}
	selfAbs, err := filepath.Abs(selfPath)
	if err != nil {
		return false
	}
	candidateAbs, err := filepath.Abs(candidate)
	if err != nil {
		return false
	}
	if selfAbs == candidateAbs {
		return true
	}
	selfInfo, err := os.Stat(selfAbs)
	if err != nil {
		return false
	}
	candidateInfo, err := os.Stat(candidateAbs)
	if err != nil {
		return false
	}
	return os.SameFile(selfInfo, candidateInfo)
}
