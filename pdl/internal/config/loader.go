package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/config/defaults"
)

func Load(configPath string) (RootConfig, error) {
	var result RootConfig
	absolutePath, absErr := filepath.Abs(configPath)
	if absErr != nil {
		return result, absErr
	}
	userConfig, userErr := readJSONConfig(absolutePath)
	if userErr != nil {
		return result, userErr
	}
	explicitSections := extractSectionNames(userConfig["sections"])
	basePath := filepath.Join(filepath.Dir(absolutePath), "config", "common.pdl.config.json")
	mergedConfig, mergeErr := mergeWithBase(basePath, userConfig)
	if mergeErr != nil {
		return result, mergeErr
	}
	expandEnvInPlace(mergedConfig)
	payload, encodeErr := json.Marshal(mergedConfig)
	if encodeErr != nil {
		return result, encodeErr
	}
	decodeErr := json.Unmarshal(payload, &result)
	if decodeErr != nil {
		return result, decodeErr
	}
	if len(explicitSections) > 0 {
		result.ActiveSections = make(map[string]struct{}, len(explicitSections))
		for _, name := range explicitSections {
			result.ActiveSections[name] = struct{}{}
		}
	} else {
		result.ActiveSections = make(map[string]struct{})
	}
	setDefaultPaths(&result)
	setDefaultProfileOutputs(&result)
	return result, nil
}

func readJSONConfig(configPath string) (map[string]interface{}, error) {
	var result map[string]interface{}
	rawBytes, readErr := os.ReadFile(configPath)
	if readErr != nil {
		return result, readErr
	}
	decodeErr := json.Unmarshal(rawBytes, &result)
	if decodeErr != nil {
		return result, decodeErr
	}
	return result, nil
}

func mergeWithBase(basePath string, userConfig map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	baseConfig, loadErr := readBaseConfig(basePath)
	if loadErr != nil {
		return result, loadErr
	}
	mergedConfig, mergeErr := mergeConfigData(baseConfig, userConfig)
	if mergeErr != nil {
		return result, mergeErr
	}
	result = mergedConfig
	return result, nil
}

func mergeConfigData(baseConfig map[string]interface{}, userConfig map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	result = deepCopyMap(baseConfig)
	deepMergeMaps(result, userConfig)
	baseSections := extractSectionSlice(baseConfig["sections"])
	userSections := extractSectionSlice(userConfig["sections"])
	if len(baseSections) > 0 || len(userSections) > 0 {
		mergedSections := combineSections(baseSections, userSections)
		result["sections"] = mergedSections
		if len(userSections) == 0 {
			adjustErr := adjustSections(result, userSections)
			if adjustErr != nil {
				return result, adjustErr
			}
		}
	}
	adjustGlobalIndexFilename(result)
	adjustDb2Pdl(result)
	return result, nil
}

func deepCopyMap(input map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(input))
	for key, value := range input {
		result[key] = deepCopyValue(value)
	}
	return result
}

func deepCopyValue(value interface{}) interface{} {
	var result interface{}
	switch typed := value.(type) {
	case map[string]interface{}:
		result = deepCopyMap(typed)
	case []interface{}:
		copied := make([]interface{}, len(typed))
		for index, item := range typed {
			copied[index] = deepCopyValue(item)
		}
		result = copied
	default:
		result = typed
	}
	return result
}

func deepMergeMaps(destination map[string]interface{}, source map[string]interface{}) {
	for key, value := range source {
		if valueMap, ok := value.(map[string]interface{}); ok {
			if existing, exists := destination[key].(map[string]interface{}); exists {
				deepMergeMaps(existing, valueMap)
				continue
			}
			destination[key] = deepCopyMap(valueMap)
			continue
		}
		if valueSlice, ok := value.([]interface{}); ok {
			destination[key] = deepCopyValue(valueSlice)
			continue
		}
		destination[key] = value
	}
}

func extractSectionSlice(value interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	rawSlice, ok := value.([]interface{})
	if !ok {
		result = make([]map[string]interface{}, 0)
		return result
	}
	result = make([]map[string]interface{}, 0, len(rawSlice))
	for _, item := range rawSlice {
		if mapped, mapOk := item.(map[string]interface{}); mapOk {
			result = append(result, deepCopyMap(mapped))
		}
	}
	return result
}

func extractSectionNames(value interface{}) []string {
	sections := extractSectionSlice(value)
	result := make([]string, 0, len(sections))
	for _, section := range sections {
		if name, ok := section["name"].(string); ok && name != "" {
			result = append(result, name)
		}
	}
	return result
}

func combineSections(baseSections []map[string]interface{}, userSections []map[string]interface{}) []interface{} {
	result := make([]interface{}, 0, len(baseSections)+len(userSections))
	nameToIndex := make(map[string]int, len(baseSections))
	for _, section := range baseSections {
		if name, ok := section["name"].(string); ok && name != "" {
			nameToIndex[name] = len(result)
		}
		result = append(result, section)
	}
	for _, section := range userSections {
		name, _ := section["name"].(string)
		if name == "" {
			result = append(result, section)
			continue
		}
		if idx, exists := nameToIndex[name]; exists {
			result[idx] = section
			continue
		}
		nameToIndex[name] = len(result)
		result = append(result, section)
	}
	return result
}

func adjustSections(config map[string]interface{}, userSections []map[string]interface{}) error {
	var result error
	profilesValue, ok := config["profiles"].(map[string]interface{})
	if !ok {
		result = fmt.Errorf("profiles section missing in configuration")
		return result
	}
	sectionsValue, ok := config["sections"].([]interface{})
	if !ok {
		result = fmt.Errorf("sections collection missing in configuration")
		return result
	}
	excludePatterns, collectErr := collectUserGlobPatterns(userSections, profilesValue)
	if collectErr != nil {
		result = collectErr
		return result
	}
	dbSection, _ := findSectionByName(sectionsValue, "DbFiles")
	frontendSection, _ := findSectionByName(sectionsValue, "Frontend")
	goSection, _ := findSectionByName(sectionsValue, "Go")
	kotlinSection, _ := findSectionByName(sectionsValue, "Kotlin")
	javaSection, _ := findSectionByName(sectionsValue, "Java")
	dbFilesPath, dbFilesErr := buildDbFiles(config)
	if dbFilesErr != nil {
		result = dbFilesErr
		return result
	}
	currentDbFiles := make([]string, 0)
	if dbSection != nil && dbFilesPath != "" {
		updateSectionFiles(dbSection, "dbFiles", []string{dbFilesPath})
		appendExclusions(dbSection, "dbFilesExclude", excludePatterns)
		currentDbFiles = extractStringSlice(dbSection, "dbFiles")
	}
	if frontendSection != nil {
		profileKeys := collectSectionProfileKeys(frontendSection)
		for _, profileKey := range profileKeys {
			excludeKey := profileKey + "Exclude"
			appendExclusions(frontendSection, excludeKey, excludePatterns)
			appendExclusions(frontendSection, excludeKey, currentDbFiles)
		}
	}
	if goSection != nil {
		appendExclusions(goSection, "goExclude", excludePatterns)
		appendExclusions(goSection, "goExclude", currentDbFiles)
	}
	if kotlinSection != nil {
		appendExclusions(kotlinSection, "kotlinExclude", excludePatterns)
		appendExclusions(kotlinSection, "kotlinExclude", currentDbFiles)
	}
	if javaSection != nil {
		appendExclusions(javaSection, "javaExclude", excludePatterns)
		appendExclusions(javaSection, "javaExclude", currentDbFiles)
	}
	return result
}

func setDefaultPaths(cfg *RootConfig) {
	if cfg == nil {
		return
	}
	cfg.OutputDir = strings.TrimSpace(cfg.OutputDir)
	if cfg.OutputDir == "" {
		cfg.OutputDir = "output"
	}
	cfg.TempDir = strings.TrimSpace(cfg.TempDir)
	if cfg.TempDir == "" {
		cfg.TempDir = filepath.Join(cfg.OutputDir, "temp")
	}
	cfg.Db2PdlSourceDest = strings.TrimSpace(cfg.Db2PdlSourceDest)
	cfg.Db2Pdl.Db2PdlSourceDest = strings.TrimSpace(cfg.Db2Pdl.Db2PdlSourceDest)
	if cfg.Db2Pdl.Db2PdlSourceDest == "" {
		cfg.Db2Pdl.Db2PdlSourceDest = cfg.Db2PdlSourceDest
	}
	cfg.Db2Pdl.OutputDir = strings.TrimSpace(cfg.Db2Pdl.OutputDir)
	if cfg.Db2Pdl.OutputDir == "" && (cfg.Db2Pdl.Enabled || cfg.Db2Pdl.Db2PdlSourceDest != "") {
		cfg.Db2Pdl.OutputDir = filepath.Join(cfg.OutputDir, "pdl")
	}
	if !cfg.Db2Pdl.Enabled && cfg.Db2Pdl.Db2PdlSourceDest != "" {
		cfg.Db2Pdl.Enabled = true
	}
	if cfg.Db2Pdl.Options == nil {
		cfg.Db2Pdl.Options = make(map[string]interface{})
	}
	cfg.Db2PdlSourceDest = cfg.Db2Pdl.Db2PdlSourceDest
}

func setDefaultProfileOutputs(cfg *RootConfig) {
	if cfg == nil || cfg.Profiles == nil {
		return
	}
	baseDir := strings.TrimSpace(cfg.OutputDir)
	for name, profile := range cfg.Profiles {
		if profile.Options == nil {
			profile.Options = make(map[string]interface{})
		}
		language := profile.LanguageOrDefault(name)
		profile.Language = language
		if strings.TrimSpace(profile.OutputDir) == "" {
			profile.OutputDir = filepath.Join(baseDir, language)
		}
		cfg.Profiles[name] = profile
	}
}

func collectSectionProfileKeys(section map[string]interface{}) []string {
	filesValue, ok := section["files"].(map[string]interface{})
	if !ok {
		return []string{}
	}
	result := make([]string, 0, len(filesValue))
	for key := range filesValue {
		if strings.HasSuffix(key, "Exclude") {
			continue
		}
		result = append(result, key)
	}
	return result
}

func collectUserGlobPatterns(userSections []map[string]interface{}, profiles map[string]interface{}) ([]string, error) {
	var result []string
	result = make([]string, 0)
	for _, section := range userSections {
		filesValue, ok := section["files"].(map[string]interface{})
		if !ok {
			continue
		}
		for profileName, rawPatterns := range filesValue {
			if strings.HasSuffix(profileName, "Exclude") {
				continue
			}
			if _, ok := profiles[profileName]; !ok {
				return result, fmt.Errorf("invalid profile: %s", profileName)
			}
			patterns := interfaceSliceToStrings(rawPatterns)
			result = append(result, patterns...)
		}
	}
	return result, nil
}

func findSectionByName(sections []interface{}, name string) (map[string]interface{}, error) {
	var result map[string]interface{}
	for _, rawSection := range sections {
		sectionMap, ok := rawSection.(map[string]interface{})
		if !ok {
			continue
		}
		sectionName, _ := sectionMap["name"].(string)
		if sectionName == name {
			result = sectionMap
			return result, nil
		}
	}
	return result, fmt.Errorf("section %s not found", name)
}

func buildDbFiles(config map[string]interface{}) (string, error) {
	var result string
	source := trimString(config["db2PdlSourceDest"])
	db2pdlValue, _ := config["db2pdl"].(map[string]interface{})
	if db2pdlValue == nil {
		db2pdlValue, _ = config["orm"].(map[string]interface{})
	}
	if db2pdlValue != nil {
		if candidate := trimString(db2pdlValue["db2PdlSourceDest"]); candidate != "" {
			source = candidate
		}
	}
	if source == "" {
		return result, nil
	}
	result = source + "/*.pdl"
	return result, nil
}

func updateSectionFiles(section map[string]interface{}, key string, values []string) {
	filesValue, ok := section["files"].(map[string]interface{})
	if !ok {
		filesValue = make(map[string]interface{})
		section["files"] = filesValue
	}
	filesValue[key] = stringsToInterfaceSlice(values)
}

func appendExclusions(section map[string]interface{}, key string, additions []string) {
	if len(additions) == 0 {
		return
	}
	filesValue, ok := section["files"].(map[string]interface{})
	if !ok {
		filesValue = make(map[string]interface{})
		section["files"] = filesValue
	}
	existing := interfaceSliceToStrings(filesValue[key])
	existing = append(existing, additions...)
	filesValue[key] = stringsToInterfaceSlice(existing)
}

func extractStringSlice(section map[string]interface{}, key string) []string {
	var result []string
	filesValue, ok := section["files"].(map[string]interface{})
	if !ok {
		result = make([]string, 0)
		return result
	}
	result = interfaceSliceToStrings(filesValue[key])
	return result
}

func interfaceSliceToStrings(value interface{}) []string {
	var result []string
	rawSlice, ok := value.([]interface{})
	if !ok {
		result = make([]string, 0)
		return result
	}
	result = make([]string, 0, len(rawSlice))
	for _, item := range rawSlice {
		if text, ok := item.(string); ok {
			result = append(result, text)
		}
	}
	return result
}

func stringsToInterfaceSlice(values []string) []interface{} {
	var result []interface{}
	result = make([]interface{}, len(values))
	for index, value := range values {
		result[index] = value
	}
	return result
}

func trimString(value interface{}) string {
	var result string
	text, ok := value.(string)
	if !ok {
		result = ""
		return result
	}
	result = strings.TrimSpace(text)
	return result
}

func adjustGlobalIndexFilename(config map[string]interface{}) {
	projectName, ok := config["project"].(string)
	if !ok || projectName == "" {
		return
	}
	jsValue, ok := config["js"].(map[string]interface{})
	if !ok {
		return
	}
	globalIndex, ok := jsValue["globalIndex"].(map[string]interface{})
	if !ok {
		return
	}
	filename := projectName + "Pdl.js"
	globalIndex["filename"] = filename
}

func adjustDb2Pdl(config map[string]interface{}) {
	source := trimString(config["db2PdlSourceDest"])
	db2pdlValue, _ := config["db2pdl"].(map[string]interface{})
	if db2pdlValue == nil {
		if legacy, ok := config["db2Pdl"].(map[string]interface{}); ok {
			db2pdlValue = legacy
			config["db2pdl"] = db2pdlValue
		}
	}
	if db2pdlValue != nil {
		if candidate := trimString(db2pdlValue["db2PdlSourceDest"]); candidate != "" {
			source = candidate
		}
	}
	legacyOrm, _ := config["orm"].(map[string]interface{})
	if legacyOrm != nil {
		if candidate := trimString(legacyOrm["db2PdlSourceDest"]); candidate != "" {
			source = candidate
		}
		if db2pdlValue == nil {
			db2pdlValue = make(map[string]interface{})
			config["db2pdl"] = db2pdlValue
		}
		if nested, ok := legacyOrm["db2pdl"].(map[string]interface{}); ok {
			deepMergeMaps(db2pdlValue, nested)
		}
		if _, exists := db2pdlValue["enabled"]; !exists {
			if enabled, ok := legacyOrm["enabled"].(bool); ok {
				db2pdlValue["enabled"] = enabled
			}
		}
		if _, exists := db2pdlValue["outputDir"]; !exists {
			if output := trimString(legacyOrm["outputDir"]); output != "" {
				db2pdlValue["outputDir"] = output
			}
		}
	}
	if source == "" {
		return
	}
	if db2pdlValue == nil {
		db2pdlValue = make(map[string]interface{})
		config["db2pdl"] = db2pdlValue
	}
	pdlValue, ok := db2pdlValue["pdl"].(map[string]interface{})
	if !ok {
		pdlValue = make(map[string]interface{})
		db2pdlValue["pdl"] = pdlValue
	}
	pdlValue["db2PdlSourceDest"] = source
	pdlValue["entitiesNamespace"] = replaceSlashes(source)
}

func replaceSlashes(value string) string {
	var result string
	if value == "" {
		result = ""
		return result
	}
	builder := make([]rune, 0, len(value))
	for _, char := range value {
		if char == '/' {
			builder = append(builder, '.')
			continue
		}
		builder = append(builder, char)
	}
	result = string(builder)
	return result
}

func expandEnvInPlace(value interface{}) {
	switch typed := value.(type) {
	case map[string]interface{}:
		for key, inner := range typed {
			expandEnvInPlace(inner)
			if text, ok := inner.(string); ok {
				typed[key] = os.ExpandEnv(text)
			} else {
				typed[key] = inner
			}
		}
	case []interface{}:
		for index, inner := range typed {
			expandEnvInPlace(inner)
			if text, ok := inner.(string); ok {
				typed[index] = os.ExpandEnv(text)
			} else {
				typed[index] = inner
			}
		}
	}
}

func readBaseConfig(basePath string) (map[string]interface{}, error) {
	var result map[string]interface{}
	rawBytes, readErr := os.ReadFile(basePath)
	if readErr == nil {
		decodeErr := json.Unmarshal(rawBytes, &result)
		if decodeErr != nil {
			return nil, decodeErr
		}
		return result, nil
	}
	if !errors.Is(readErr, os.ErrNotExist) {
		return nil, readErr
	}
	return loadEmbeddedBaseConfig()
}

func loadEmbeddedBaseConfig() (map[string]interface{}, error) {
	bytes, err := defaults.Common()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	decodeErr := json.Unmarshal(bytes, &result)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return result, nil
}
