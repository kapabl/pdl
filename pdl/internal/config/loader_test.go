package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMergesConfigs(t *testing.T) {
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, "config")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		t.Fatalf("failed to create config directory: %v", err)
	}
	outputDir := filepath.Join(tempDir, "output")
	os.Setenv("PDL_OUTPUT", outputDir)
	os.Setenv("PDL_TEMPLATES_DIR", filepath.Join(tempDir, "templates"))

	baseConfig := `{
  "companyName": "BaseCo",
  "project": "Base",
  "version": "1.0.0",
  "outputDir": "${PDL_OUTPUT}",
  "db2pdl": {"enabled": false, "db2PdlSourceDest": "base/domain/data"},
  "templates": {"dir": "", "name": "classTemplate1"},
  "src": ["src/base"],
  "profiles": {
    "js": {"language": "js", "enabled": false, "generateAsObject": false},
    "php": {"language": "php", "enabled": false, "psr4": true}
  },
  "sections": [
    {"name": "Frontend", "files": {"js": ["base/**/*.pdl"], "jsExclude": []}}
  ]
}`

	userConfig := `{
  "project": "Derived",
  "db2pdl": {"enabled": true, "db2PdlSourceDest": "custom/domain/data"},
  "profiles": {
    "js": {"language": "js", "enabled": true, "generateAsObject": true, "namespaceFile": "customNamespace.js"},
    "go": {"enabled": true}
  },
  "sections": [
    {"name": "Frontend", "files": {"js": ["custom/**/*.pdl"], "jsExclude": []}},
    {"name": "Go", "files": {"go": ["domain/**/*.pdl"], "goExclude": []}}
  ]
}`

	basePath := filepath.Join(configDir, "common.pdl.config.json")
	userPath := filepath.Join(tempDir, "pdl.config.json")
	if err := os.WriteFile(basePath, []byte(baseConfig), 0o644); err != nil {
		t.Fatalf("failed to write base config: %v", err)
	}
	if err := os.WriteFile(userPath, []byte(userConfig), 0o644); err != nil {
		t.Fatalf("failed to write project config: %v", err)
	}

	cfg, err := Load(userPath)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if cfg.Project != "Derived" {
		t.Fatalf("expected project to be Derived, got %s", cfg.Project)
	}
	profileJS, ok := cfg.Profiles["js"]
	if !ok {
		t.Fatalf("expected js profile")
	}
	if !profileJS.Enabled {
		t.Fatalf("expected js profile enabled")
	}
	if profileJS.Language != "js" {
		t.Fatalf("expected js language to be js, got %s", profileJS.Language)
	}
	if profileJS.OutputDir != filepath.Join(outputDir, "js") {
		t.Fatalf("expected js outputDir %s, got %s", filepath.Join(outputDir, "js"), profileJS.OutputDir)
	}
	if value, ok := profileJS.Options["generateAsObject"].(bool); !ok || !value {
		t.Fatalf("expected generateAsObject option to be true, got %#v", profileJS.Options["generateAsObject"])
	}
	if value, ok := profileJS.Options["namespaceFile"].(string); !ok || value != "customNamespace.js" {
		t.Fatalf("expected namespaceFile override, got %#v", profileJS.Options["namespaceFile"])
	}
	profileGo, ok := cfg.Profiles["go"]
	if !ok {
		t.Fatalf("expected go profile")
	}
	if profileGo.Language != "go" {
		t.Fatalf("expected go language default, got %s", profileGo.Language)
	}
	if profileGo.OutputDir != filepath.Join(outputDir, "go") {
		t.Fatalf("expected go outputDir %s, got %s", filepath.Join(outputDir, "go"), profileGo.OutputDir)
	}
	if len(cfg.ActiveSections) != 2 {
		t.Fatalf("expected two active sections, got %d", len(cfg.ActiveSections))
	}
	if !cfg.Db2Pdl.Enabled {
		t.Fatalf("expected db2pdl enabled")
	}
	if cfg.Db2Pdl.Db2PdlSourceDest != "custom/domain/data" {
		t.Fatalf("expected db2pdl db2PdlSourceDest custom/domain/data, got %s", cfg.Db2Pdl.Db2PdlSourceDest)
	}
	if cfg.Db2Pdl.OutputDir != filepath.Join(outputDir, "pdl") {
		t.Fatalf("expected db2pdl outputDir %s, got %s", filepath.Join(outputDir, "pdl"), cfg.Db2Pdl.OutputDir)
	}
	if cfg.Db2PdlSourceDest != cfg.Db2Pdl.Db2PdlSourceDest {
		t.Fatalf("expected legacy db2PdlSourceDest to match db2pdl, got %s vs %s", cfg.Db2PdlSourceDest, cfg.Db2Pdl.Db2PdlSourceDest)
	}
}

func TestLoadUsesEmbeddedDefaults(t *testing.T) {
	tempDir := t.TempDir()
	outputDir := filepath.Join(tempDir, "output")
	os.Setenv("PDL_OUTPUT", outputDir)
	os.Setenv("PDL_TEMPLATES_DIR", filepath.Join(tempDir, "templates"))

	payload := `{
  "project": "Embedded",
  "profiles": {
    "js": {"enabled": true, "generateAsObject": true}
  },
  "sections": [
    {"name": "Frontend", "files": {"js": ["domain/**/*.pdl"], "jsExclude": []}}
  ]
}`
	configPath := filepath.Join(tempDir, "pdl.config.json")
	if err := os.WriteFile(configPath, []byte(payload), 0o644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	profileJS, ok := cfg.Profiles["js"]
	if !ok {
		t.Fatalf("expected js profile in embedded defaults")
	}
	if profileJS.Language != "js" {
		t.Fatalf("expected js language, got %s", profileJS.Language)
	}
	if profileJS.OutputDir != filepath.Join(outputDir, "js") {
		t.Fatalf("expected js outputDir %s, got %s", filepath.Join(outputDir, "js"), profileJS.OutputDir)
	}
	if value, ok := profileJS.Options["generateAsObject"].(bool); !ok || !value {
		t.Fatalf("expected generateAsObject to be true, got %#v", profileJS.Options["generateAsObject"])
	}
}
