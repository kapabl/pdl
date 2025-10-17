package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type AttributeTemplates struct {
	DbID       string `json:"dbId"`
	ColumnName string `json:"columnName"`
}

type PHPConfig struct {
	EmitHelpers bool               `json:"emitHelpers"`
	Attributes  AttributeTemplates `json:"attributes"`
}

type PDLConfig struct {
	DB2PDLSourceDest  string             `json:"db2PdlSourceDest"`
	EntitiesNamespace string             `json:"entitiesNamespace"`
	UseNamespaces     []string           `json:"useNamespaces"`
	Attributes        AttributeTemplates `json:"attributes"`
}

type TypeScriptConfig struct {
	Emit       bool   `json:"emit"`
	OutputFile string `json:"outputFile"`
}

type CSConfig struct {
	Emit bool `json:"emit"`
}

type GoConfig struct {
	Emit    bool   `json:"emit"`
	Package string `json:"package"`
}

type JavaConfig struct {
	Emit    bool   `json:"emit"`
	Package string `json:"package"`
}

type KotlinConfig struct {
	Emit    bool   `json:"emit"`
	Package string `json:"package"`
}

type PythonConfig struct {
	Emit    bool   `json:"emit"`
	Package string `json:"package"`
}

type RustConfig struct {
	Emit bool `json:"emit"`
}

type CppConfig struct {
	Emit bool `json:"emit"`
}

type ConnectionConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type DB2PDLConfig struct {
	Enabled         bool             `json:"enabled"`
	Verbose         bool             `json:"verbose"`
	Connection      ConnectionConfig `json:"connection"`
	OutputDir       string           `json:"outputDir"`
	TemplatesDir    string           `json:"templatesDir"`
	TypeScript      TypeScriptConfig `json:"ts"`
	Java            JavaConfig       `json:"java"`
	Kotlin          KotlinConfig     `json:"kotlin"`
	Python          PythonConfig     `json:"python"`
	Rust            RustConfig       `json:"rust"`
	Cpp             CppConfig        `json:"cpp"`
	CSharp          CSConfig         `json:"cs"`
	Go              GoConfig         `json:"go"`
	PHP             PHPConfig        `json:"php"`
	PDL             PDLConfig        `json:"pdl"`
	ExcludedTables  []string         `json:"excludedTables"`
	ExcludedColumns []string         `json:"excludedColumns"`
}

type RootConfig struct {
	CompanyName      string       `json:"companyName"`
	Project          string       `json:"project"`
	Version          string       `json:"version"`
	OutputDir        string       `json:"outputDir"`
	Verbose          bool         `json:"verbose"`
	Db2PdlSourceDest string       `json:"db2PdlSourceDest"`
	DB2PDL           DB2PDLConfig `json:"db2pdl"`
}

func LoadConfig(configPath string) (RootConfig, error) {
	var result RootConfig
	absolutePath, absErr := filepath.Abs(configPath)
	if absErr != nil {
		return result, absErr
	}
	userConfig, userErr := readJSONConfig(absolutePath)
	if userErr != nil {
		return result, userErr
	}
	mergedConfig, mergeErr := buildMergedConfig(absolutePath, userConfig)
	if mergeErr != nil {
		return result, mergeErr
	}
	if decodeErr := decodeRootConfig(mergedConfig, &result); decodeErr != nil {
		return result, decodeErr
	}
	finalizeDb2Pdl(absolutePath, &result)
	return result, nil
}

func buildMergedConfig(configPath string, userConfig map[string]interface{}) (map[string]interface{}, error) {
	dir := filepath.Dir(configPath)
	external, extErr := readExternalDb2PdlConfig(dir)
	if extErr != nil {
		return nil, extErr
	}
	if external != nil {
		deepMergeMaps(userConfig, external)
	}
	basePath := filepath.Join(dir, "config", "common.pdl.config.json")
	return mergeWithBase(basePath, userConfig)
}

func readExternalDb2PdlConfig(configDir string) (map[string]interface{}, error) {
	candidate := filepath.Join(configDir, "pdl.db2pdl.config.json")
	rawBytes, readErr := os.ReadFile(candidate)
	if readErr != nil {
		if errors.Is(readErr, os.ErrNotExist) {
			return nil, nil
		}
		return nil, readErr
	}
	var payload interface{}
	if err := json.Unmarshal(rawBytes, &payload); err != nil {
		return nil, err
	}
	mapped, ok := payload.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("db2pdl config must be a JSON object")
	}
	if _, exists := mapped["db2pdl"]; exists {
		return mapped, nil
	}
	return map[string]interface{}{"db2pdl": mapped}, nil
}

func decodeRootConfig(data map[string]interface{}, target *RootConfig) error {
	adjustDb2Pdl(data)
	expandEnvInPlace(data)
	payload, encodeErr := json.Marshal(data)
	if encodeErr != nil {
		return encodeErr
	}
	if err := json.Unmarshal(payload, target); err != nil {
		return err
	}
	return nil
}

func finalizeDb2Pdl(configPath string, cfg *RootConfig) {
	cfg.OutputDir = strings.TrimSpace(cfg.OutputDir)
	if cfg.OutputDir == "" {
		cfg.OutputDir = "output"
	}
	cfg.Db2PdlSourceDest = strings.TrimSpace(cfg.Db2PdlSourceDest)
	db2PdlSource := strings.TrimSpace(cfg.DB2PDL.PDL.DB2PDLSourceDest)
	if db2PdlSource == "" {
		db2PdlSource = cfg.Db2PdlSourceDest
	}
	cfg.DB2PDL.PDL.DB2PDLSourceDest = db2PdlSource
	cfg.Db2PdlSourceDest = db2PdlSource
	if cfg.Verbose && !cfg.DB2PDL.Verbose {
		cfg.DB2PDL.Verbose = true
	}
	cfg.DB2PDL.OutputDir = strings.TrimSpace(cfg.DB2PDL.OutputDir)
	if cfg.DB2PDL.OutputDir == "" {
		cfg.DB2PDL.OutputDir = filepath.Join(cfg.OutputDir, "pdl")
	}
	configDir := filepath.Dir(configPath)
	if cfg.DB2PDL.OutputDir != "" && !filepath.IsAbs(cfg.DB2PDL.OutputDir) {
		cfg.DB2PDL.OutputDir = filepath.Join(configDir, cfg.DB2PDL.OutputDir)
	}
	if cfg.DB2PDL.TemplatesDir != "" && cfg.DB2PDL.TemplatesDir != "default" && !filepath.IsAbs(cfg.DB2PDL.TemplatesDir) {
		cfg.DB2PDL.TemplatesDir = filepath.Join(configDir, cfg.DB2PDL.TemplatesDir)
	}
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
	rawBytes, readErr := os.ReadFile(basePath)
	if readErr != nil {
		if errors.Is(readErr, os.ErrNotExist) {
			result = deepCopyMap(userConfig)
			return result, nil
		}
		return result, readErr
	}
	var baseConfig map[string]interface{}
	decodeErr := json.Unmarshal(rawBytes, &baseConfig)
	if decodeErr != nil {
		return result, decodeErr
	}
	mergedConfig := deepCopyMap(baseConfig)
	deepMergeMaps(mergedConfig, userConfig)
	result = mergedConfig
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
	switch typed := value.(type) {
	case map[string]interface{}:
		return deepCopyMap(typed)
	case []interface{}:
		copied := make([]interface{}, len(typed))
		for index, item := range typed {
			copied[index] = deepCopyValue(item)
		}
		return copied
	default:
		return typed
	}
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

func adjustDb2Pdl(config map[string]interface{}) {
	sourceValue := trimString(config["db2PdlSourceDest"])
	db2pdlValue, _ := config["db2pdl"].(map[string]interface{})
	if db2pdlValue == nil {
		if legacy, ok := config["db2Pdl"].(map[string]interface{}); ok {
			db2pdlValue = legacy
			config["db2pdl"] = db2pdlValue
		}
	}
	if db2pdlValue != nil {
		if candidate := trimString(db2pdlValue["db2PdlSourceDest"]); candidate != "" {
			sourceValue = candidate
		}
	}
	ormValue, _ := config["orm"].(map[string]interface{})
	if ormValue != nil {
		if candidate := trimString(ormValue["db2PdlSourceDest"]); candidate != "" {
			sourceValue = candidate
		}
		if db2pdlValue == nil {
			db2pdlValue = make(map[string]interface{})
			config["db2pdl"] = db2pdlValue
		}
		if nested, ok := ormValue["db2pdl"].(map[string]interface{}); ok {
			deepMergeMaps(db2pdlValue, nested)
		}
		if _, exists := db2pdlValue["enabled"]; !exists {
			if enabled, ok := ormValue["enabled"].(bool); ok {
				db2pdlValue["enabled"] = enabled
			}
		}
		if _, exists := db2pdlValue["outputDir"]; !exists {
			if output := trimString(ormValue["outputDir"]); output != "" {
				db2pdlValue["outputDir"] = output
			}
		}
	}
	if db2pdlValue == nil {
		db2pdlValue = make(map[string]interface{})
		config["db2pdl"] = db2pdlValue
	}
	if _, exists := db2pdlValue["verbose"]; !exists {
		if topVerbose, ok := config["verbose"].(bool); ok {
			db2pdlValue["verbose"] = topVerbose
		}
	}
	if sourceValue == "" {
		return
	}
	pdlValue, ok := db2pdlValue["pdl"].(map[string]interface{})
	if !ok {
		pdlValue = make(map[string]interface{})
		db2pdlValue["pdl"] = pdlValue
	}
	normalizeUseNamespaces(pdlValue)
	pdlValue["db2PdlSourceDest"] = sourceValue
	pdlValue["entitiesNamespace"] = replaceSlashes(sourceValue)
}

func trimString(value interface{}) string {
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(text)
}

func replaceSlashes(value string) string {
	if value == "" {
		return ""
	}
	return strings.ReplaceAll(value, "/", ".")
}

func normalizeUseNamespaces(pdlValue map[string]interface{}) {
	if pdlValue == nil {
		return
	}
	if _, exists := pdlValue["useNamespaces"]; exists {
		return
	}
	if raw, ok := pdlValue["use"]; ok {
		pdlValue["useNamespaces"] = toStringSlice(raw)
	}
}

func toStringSlice(value interface{}) []string {
	switch typed := value.(type) {
	case []string:
		return typed
	case []interface{}:
		result := make([]string, 0, len(typed))
		for _, item := range typed {
			result = append(result, trimString(item))
		}
		return result
	case string:
		return []string{strings.TrimSpace(typed)}
	default:
		return []string{}
	}
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

func ValidateConfig(config RootConfig) error {
	if config.DB2PDL.Connection.Type == "" {
		return fmt.Errorf("database connection type is required")
	}
	return nil
}
