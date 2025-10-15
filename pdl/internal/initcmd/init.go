package initcmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates/*
var templateFS embed.FS

// Options controls how project scaffolding is generated.
type Options struct {
	// TargetDir is the directory that will receive the generated PDL project.
	// If empty, "pdl" (relative to the current working directory) is used.
	TargetDir       string
	CompanyName     string
	ProjectName     string
	BackendTargets  []string
	FrontendTargets []string
}

// Result captures the outcome of scaffolding.
type Result struct {
	RootDir            string
	Created            []string
	Skipped            []string
	AlreadyInitialized bool
}

// Run generates the starter layout for a project consuming the PDL CLI.
func Run(options Options) (Result, error) {
	var result Result

	targetDir := resolveTargetDir(options.TargetDir)
	absTarget, absErr := filepath.Abs(targetDir)
	if absErr != nil {
		return result, absErr
	}
	result.RootDir = absTarget

	alreadyInitialized := workspaceInitialized(absTarget)

	if err := ensureWorkspaceDirs(absTarget); err != nil {
		return result, err
	}

	staticCreated, staticSkipped, staticErr := writeStaticTemplates(absTarget)
	if staticErr != nil {
		return result, staticErr
	}
	result.Created = append(result.Created, staticCreated...)
	result.Skipped = append(result.Skipped, staticSkipped...)

	configCreated, configSkipped, configErr := configureWorkspace(absTarget, options)
	if configErr != nil {
		return result, configErr
	}
	result.Created = append(result.Created, configCreated...)
	result.Skipped = append(result.Skipped, configSkipped...)
	result.AlreadyInitialized = alreadyInitialized

	return result, nil
}

func resolveTargetDir(input string) string {
	if input == "" {
		return "pdl"
	}
	return input
}

func ensureWorkspaceDirs(root string) error {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(root, "src"), 0o755)
}

func writeStaticTemplates(root string) ([]string, []string, error) {
	created := make([]string, 0)
	skipped := make([]string, 0)
	entries := []struct {
		relative string
		template string
	}{
		{relative: ".env.local", template: "env.local"},
		{relative: ".gitignore", template: "gitignore"},
		{relative: "src/README.md", template: "srcREADME.md"},
	}
	for _, entry := range entries {
		path := filepath.Join(root, entry.relative)
		if exists(path) {
			skipped = append(skipped, path)
			continue
		}
		if err := writeTemplate(path, entry.template); err != nil {
			return nil, nil, err
		}
		created = append(created, path)
	}
	return created, skipped, nil
}

func configureWorkspace(root string, options Options) ([]string, []string, error) {
	created := make([]string, 0)
	skipped := make([]string, 0)
	mainCreated, err := writeMainConfig(root, options)
	if err != nil {
		return nil, nil, err
	}
	if mainCreated {
		created = append(created, filepath.Join(root, "pdl.config.json"))
	} else {
		skipped = append(skipped, filepath.Join(root, "pdl.config.json"))
	}
	return created, skipped, nil
}

func writeTemplate(destination string, templateName string) error {
	content, readErr := templateFS.ReadFile(filepath.Join("templates", templateName))
	if readErr != nil {
		return readErr
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(destination, content, 0o644); err != nil {
		return err
	}
	return nil
}

func exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// FormatResult renders a human-readable summary that can be printed by callers.
func FormatResult(result Result) string {
	summary := ""
	if result.AlreadyInitialized {
		summary += fmt.Sprintf("Reused existing PDL workspace at %s\n", result.RootDir)
	} else {
		summary += fmt.Sprintf("Initialized PDL workspace at %s\n", result.RootDir)
	}
	if len(result.Created) > 0 {
		summary += "Created:\n"
		for _, path := range result.Created {
			summary += fmt.Sprintf("  - %s\n", path)
		}
	}
	if len(result.Skipped) > 0 {
		summary += "Skipped existing files:\n"
		for _, path := range result.Skipped {
			summary += fmt.Sprintf("  - %s\n", path)
		}
	}
	if result.AlreadyInitialized || len(result.Skipped) > 0 {
		summary += "No files were overwritten.\n"
	}
	return summary
}

func workspaceInitialized(root string) bool {
	return exists(filepath.Join(root, "pdl.config.json"))
}
