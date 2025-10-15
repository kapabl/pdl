package initcmd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func writeMainConfig(root string, options Options) (bool, error) {
	path := filepath.Join(root, "pdl.config.json")
	if exists(path) {
		return false, nil
	}
	data, err := buildConfigBytes(options)
	if err != nil {
		return false, err
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return false, err
	}
	return true, nil
}

func buildConfigBytes(options Options) ([]byte, error) {
	backend := FilterSelection(options.BackendTargets, BackendChoices())
	frontend := FilterSelection(options.FrontendTargets, FrontendChoices())
	frameworks, languages := SplitFrontendSelections(frontend)
	payload := map[string]interface{}{
		"companyName": fallbackString(options.CompanyName, "My Company"),
		"project":     fallbackString(options.ProjectName, "My Project"),
		"version":     "1.0.0",
		"db2pdl": map[string]interface{}{
			"enabled":          false,
			"db2PdlSourceDest": "domain/data",
		},
		"sections": []interface{}{},
		"targets": map[string]interface{}{
			"backend": backend,
			"frontend": map[string]interface{}{
				"frameworks": frameworks,
				"languages":  languages,
			},
		},
	}
	return json.MarshalIndent(payload, "", "  ")
}

func fallbackString(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
