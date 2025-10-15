package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/initcmd"
)

type promptConfig struct {
	Label   string
	Options []string
	Default []string
	Aliases map[string]string
}

func collectInitOptions(targetDir string) (initcmd.Options, error) {
	reader := bufio.NewReader(os.Stdin)
	company, err := promptForText(reader, "Company name", "My Company")
	if err != nil {
		return initcmd.Options{}, err
	}
	project, err := promptForText(reader, "Project name", "My Project")
	if err != nil {
		return initcmd.Options{}, err
	}
	backend, err := promptForMulti(reader, promptConfig{
		Label:   "Backend targets (comma separated) (use ts for TypeScript runtime)",
		Options: initcmd.BackendChoices(),
		Default: []string{"go"},
		Aliases: map[string]string{
			"typescript": "ts",
			"ts":         "ts",
		},
	})
	if err != nil {
		return initcmd.Options{}, err
	}
	frontend, err := promptForMulti(reader, promptConfig{
		Label:   "Frontend targets (comma separated) (use ts for TypeScript, js for JavaScript)",
		Options: initcmd.FrontendChoices(),
		Default: []string{"react", "ts"},
		Aliases: map[string]string{
			"typescript": "ts",
			"ts":         "ts",
			"javascript": "js",
			"js":         "js",
		},
	})
	if err != nil {
		return initcmd.Options{}, err
	}
	return initcmd.Options{
		TargetDir:       targetDir,
		CompanyName:     company,
		ProjectName:     project,
		BackendTargets:  backend,
		FrontendTargets: frontend,
	}, nil
}

func promptForText(reader *bufio.Reader, label string, defaultValue string) (string, error) {
	fmt.Printf("%s [%s]: ", label, defaultValue)
	raw, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	value := strings.TrimSpace(raw)
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}

func promptForMulti(reader *bufio.Reader, config promptConfig) ([]string, error) {
	for {
		fmt.Printf("%s [%s] (default: %s): ", config.Label, strings.Join(config.Options, ", "), strings.Join(config.Default, ", "))
		raw, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		parsed, parseErr := parseSelections(raw, config)
		if parseErr == nil {
			return parsed, nil
		}
		fmt.Println(parseErr)
	}
}

func parseSelections(input string, config promptConfig) ([]string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return initcmd.FilterSelection(normalizeDefaults(config.Default, config), config.Options), nil
	}
	split := strings.Split(trimmed, ",")
	normalized := make([]string, 0, len(split))
	allowedSet := make(map[string]string, len(config.Options)+len(config.Aliases))
	for _, option := range config.Options {
		allowedSet[option] = option
	}
	for alias, canonical := range config.Aliases {
		allowedSet[alias] = canonical
	}
	seen := make(map[string]struct{}, len(split))
	for _, entry := range split {
		key := strings.ToLower(strings.TrimSpace(entry))
		if key == "" {
			continue
		}
		canonical, ok := allowedSet[key]
		if !ok {
			return nil, fmt.Errorf("invalid option: %s", key)
		}
		if _, duplicate := seen[canonical]; duplicate {
			continue
		}
		seen[canonical] = struct{}{}
		normalized = append(normalized, canonical)
	}
	if len(normalized) == 0 {
		return nil, fmt.Errorf("no valid selections provided")
	}
	return normalized, nil
}

func normalizeDefaults(defaults []string, config promptConfig) []string {
	result := make([]string, 0, len(defaults))
	for _, value := range defaults {
		key := strings.ToLower(strings.TrimSpace(value))
		if key == "" {
			continue
		}
		if canonical, ok := config.Aliases[key]; ok {
			result = append(result, canonical)
			continue
		}
		result = append(result, key)
	}
	return result
}
