package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar/v4"

	"github.com/kapablanka/pdl/pdl/internal/config"
)

type astGeneratorBinding struct {
	profileName   string
	generatorName string
	binaryPath    string
	outputDir     string
	templateDir   string
	configPath    string
	astFiles      []string
}

type profileGeneratorSpec struct {
	name string
}

var builtinProfileGenerators = map[string]profileGeneratorSpec{
	"ts":     {name: "typescript"},
	"go":     {name: "go"},
	"js":     {name: "javascript"},
	"php":    {name: "php"},
	"csharp": {name: "csharp"},
	"java":   {name: "java"},
	"kotlin": {name: "kotlin"},
	"rust":   {name: "rust"},
	"cpp":    {name: "cpp"},
	"python": {name: "python"},
}

func (executor pipelineExecutor) runAstGenerators() error {
	var result error
	bindings, collectErr := executor.collectAstGeneratorBindings()
	if collectErr != nil {
		result = collectErr
		return result
	}
	if executor.verbose {
		executor.printer.Println(fmt.Sprintf("pipeline: ast bindings=%d", len(bindings)))
	}
	for _, binding := range bindings {
		executor.printer.Println(fmt.Sprintf("pipeline: running AST generator %s for profile %s", binding.generatorName, binding.profileName))
		for _, astPath := range binding.astFiles {
			executeErr := executor.executeAstGenerator(binding, astPath)
			if executeErr != nil {
				result = executeErr
				return result
			}
		}
	}
	return result
}

func (executor pipelineExecutor) collectAstGeneratorBindings() ([]astGeneratorBinding, error) {
	result := make([]astGeneratorBinding, 0)
	names := executor.sortedProfileNames()
	for _, name := range names {
		spec, ok := builtinProfileGenerators[name]
		if !ok {
			continue
		}
		profile := executor.spec.rootConfig.Profiles[name]
		if !profile.Enabled {
			continue
		}
		binding, buildErr := executor.buildAstBinding(name, spec, profile)
		if buildErr != nil {
			return result, buildErr
		}
		if len(binding.astFiles) == 0 {
			continue
		}
		result = append(result, binding)
	}
	return result, nil
}

func (executor pipelineExecutor) sortedProfileNames() []string {
	result := make([]string, 0, len(executor.spec.rootConfig.Profiles))
	for name := range executor.spec.rootConfig.Profiles {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}

func (executor pipelineExecutor) buildAstBinding(profileName string, spec profileGeneratorSpec, profile config.Profile) (astGeneratorBinding, error) {
	var result astGeneratorBinding
	binaryPath, binaryErr := executor.resolveGeneratorBinary("")
	if binaryErr != nil {
		return result, binaryErr
	}
	outputDir := resolveProjectPath(executor.spec.projectRoot, profile.OutputDir)
	templateDir := resolveProjectPath(executor.spec.projectRoot, profile.Templates.Dir)
	if templateDir == "" {
		templateDir = resolveProjectPath(executor.spec.projectRoot, executor.spec.rootConfig.Templates.Dir)
	}
	configPath := executor.spec.compilerConfigPath
	astFiles, collectErr := executor.discoverAstFiles(outputDir)
	if collectErr != nil {
		return result, collectErr
	}
	result = astGeneratorBinding{
		profileName:   profileName,
		generatorName: spec.name,
		binaryPath:    binaryPath,
		outputDir:     outputDir,
		templateDir:   templateDir,
		configPath:    configPath,
		astFiles:      astFiles,
	}
	return result, nil
}

func (executor pipelineExecutor) resolveGeneratorBinary(preferred string) (string, error) {
	var result string
	candidate := strings.TrimSpace(preferred)
	if candidate == "" {
		candidate = "pdlgen"
	}
	if resolved, ok := resolveBinaryPath(candidate); ok {
		result = resolved
		return result, nil
	}
	path, lookupErr := exec.LookPath(candidate)
	if lookupErr != nil {
		return result, lookupErr
	}
	result = path
	return result, nil
}

func (executor pipelineExecutor) discoverAstFiles(profileOutputDir string) ([]string, error) {
	root := executor.spec.astOutputDir
	if root == "" && executor.spec.outputDir != "" {
		root = filepath.Join(executor.spec.outputDir, "ast")
	}
	if root == "" && profileOutputDir != "" {
		parent := filepath.Dir(profileOutputDir)
		if parent != "" && parent != "." {
			root = filepath.Join(parent, "ast")
		} else {
			root = filepath.Join(profileOutputDir, "ast")
		}
	}
	if root == "" {
		return nil, nil
	}
	pattern := filepath.Join(root, "**", "*.ast.json")
	files, err := doublestar.FilepathGlob(pattern)
	if err != nil {
		return nil, err
	}
	if executor.verbose {
		executor.printer.Println(fmt.Sprintf("pipeline: ast search pattern=%s matches=%d", pattern, len(files)))
	}
	sort.Strings(files)
	return files, nil
}

func (executor pipelineExecutor) executeAstGenerator(binding astGeneratorBinding, astPath string) error {
	var result error
	args := []string{"--generator", binding.generatorName, "--ast", astPath}
	if binding.outputDir != "" {
		args = append(args, "--output", binding.outputDir)
	}
	if binding.templateDir != "" {
		args = append(args, "--templates", binding.templateDir)
	}
	if binding.configPath != "" {
		args = append(args, "--config", binding.configPath)
	}
	cmd := exec.CommandContext(executor.ctx, binding.binaryPath, args...)
	cmd.Dir = executor.spec.projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		result = runErr
		return result
	}
	return result
}
