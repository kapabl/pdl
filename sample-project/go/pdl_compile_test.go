package sample

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestFullPDLCycle(testState *testing.T) {
	if testing.Short() {
		testState.Skip("skipping PDL compilation in short mode")
	}

	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		testState.Fatalf("resolve repo root: %v", err)
	}

	pdlModuleDir := filepath.Join(repoRoot, "pdl")
	configPath := filepath.Join(repoRoot, "sample-project", "pdl-project", "pdl.config.json")
	cacheDir := testState.TempDir()

	outputDir := filepath.Join(repoRoot, "sample-project", "pdl-project", "output")
	envVars := loadEnvFile(testState, filepath.Join(repoRoot, "sample-project", "pdl-project", ".env.pg.local"))
	envVars["GOCACHE"] = cacheDir
	envVars["PDL_OUTPUT"] = outputDir
	envVars["PDL_DB2PDL_OUTPUT"] = filepath.Join(outputDir, "db2pdl")
	envVars["PDL_GEN_OUTPUT_PHP"] = filepath.Join(outputDir, "php")
	envVars["PDL_GEN_OUTPUT_JS"] = filepath.Join(outputDir, "js")
	envVars["PDL_GEN_OUTPUT_BUNDLE"] = filepath.Join(outputDir, "bundle")
	envVars["PDL_GEN_OUTPUT_GO"] = filepath.Join(outputDir, "go")
	envVars["PDL_BIN_PATH"] = filepath.Join(repoRoot, "bin")

	env := os.Environ()
	for key, value := range envVars {
		env = append(env, key+"="+value)
	}
	cmd := exec.CommandContext(context.Background(), "go", "run", "./cmd/pdlbuild", "--config", configPath, "db2pdl")
	cmd.Dir = pdlModuleDir
	cmd.Env = env

	output, runErr := cmd.CombinedOutput()
	if runErr != nil {
		testState.Fatalf("pdlbuild failed: %v\n%s", runErr, string(output))
	}
}

func loadEnvFile(testState *testing.T, envPath string) map[string]string {
	file, err := os.Open(envPath)
	if err != nil {
		return map[string]string{}
	}
	defer file.Close()

	result := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx <= 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		value := strings.TrimSpace(line[idx+1:])
		result[key] = value
	}
	if err := scanner.Err(); err != nil {
		testState.Fatalf("read env file %s: %v", envPath, err)
	}
	return result
}
