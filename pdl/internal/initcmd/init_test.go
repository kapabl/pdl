package initcmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCreatesWorkspace(testState *testing.T) {
	tempDir := testState.TempDir()
	target := filepath.Join(tempDir, "pdl")

	options := Options{
		TargetDir:       target,
		CompanyName:     "Example Co",
		ProjectName:     "Example Project",
		BackendTargets:  []string{"go", "java"},
		FrontendTargets: []string{"react", "js"},
	}
	result, err := Run(options)
	if err != nil {
		testState.Fatalf("unexpected error: %v", err)
	}
	if result.AlreadyInitialized {
		testState.Fatalf("expected fresh workspace to not be marked initialized")
	}

	expectedFiles := []string{
		"pdl.config.json",
		".env.local",
		".gitignore",
		"src/README.md",
	}

	for _, relative := range expectedFiles {
		path := filepath.Join(target, relative)
		if _, err := os.Stat(path); err != nil {
			testState.Fatalf("expected %s to exist: %v", path, err)
		}
	}

	configBytes, readErr := os.ReadFile(filepath.Join(target, "pdl.config.json"))
	if readErr != nil {
		testState.Fatalf("failed to read config: %v", readErr)
	}
	if !strings.Contains(string(configBytes), `"backend":`) {
		testState.Fatalf("expected backend targets in config, got %s", string(configBytes))
	}

	if !strings.Contains(FormatResult(result), "Initialized PDL workspace") {
		testState.Fatalf("format result should include summary, got %q", FormatResult(result))
	}
}

func TestRunSkipsExistingFiles(testState *testing.T) {
	tempDir := testState.TempDir()
	target := filepath.Join(tempDir, "pdl")
	if err := os.MkdirAll(target, 0o755); err != nil {
		testState.Fatalf("mkdir failed: %v", err)
	}
	configPath := filepath.Join(target, "pdl.config.json")
	if err := os.WriteFile(configPath, []byte("{}"), 0o644); err != nil {
		testState.Fatalf("write failed: %v", err)
	}

	options := Options{
		TargetDir:       target,
		BackendTargets:  []string{"go"},
		FrontendTargets: []string{"react"},
	}
	result, err := Run(options)
	if err != nil {
		testState.Fatalf("unexpected error: %v", err)
	}
	if !result.AlreadyInitialized {
		testState.Fatalf("expected workspace to be marked as already initialized")
	}

	skipped := false
	for _, path := range result.Skipped {
		if path == configPath {
			skipped = true
		}
	}
	if !skipped {
		testState.Fatalf("expected existing config to be skipped")
	}
	summary := FormatResult(result)
	if !strings.Contains(summary, "Reused existing PDL workspace") {
		testState.Fatalf("expected reused summary, got %q", summary)
	}
	if !strings.Contains(summary, "No files were overwritten.") {
		testState.Fatalf("expected overwrite notice, got %q", summary)
	}
}
