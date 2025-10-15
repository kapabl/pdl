package config

import (
	"encoding/json"
	"strings"
)

type TemplateConfig struct {
	Dir  string `json:"dir"`
	Name string `json:"name"`
}

type Profile struct {
	Language  string                 `json:"language"`
	Enabled   bool                   `json:"enabled"`
	Src       []string               `json:"src"`
	OutputDir string                 `json:"outputDir"`
	Templates TemplateConfig         `json:"templates"`
	Options   map[string]interface{} `json:"-"`
}

type Db2PdlConfig struct {
	Enabled          bool                   `json:"enabled"`
	OutputDir        string                 `json:"outputDir"`
	Db2PdlSourceDest string                 `json:"db2PdlSourceDest"`
	Options          map[string]interface{} `json:"-"`
}

func (configEntry *Db2PdlConfig) UnmarshalJSON(data []byte) error {
	type rawDb2Pdl struct {
		Enabled          *bool  `json:"enabled"`
		OutputDir        string `json:"outputDir"`
		Db2PdlSourceDest string `json:"db2PdlSourceDest"`
	}
	aux := rawDb2Pdl{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Enabled != nil {
		configEntry.Enabled = *aux.Enabled
	}
	configEntry.OutputDir = strings.TrimSpace(aux.OutputDir)
	configEntry.Db2PdlSourceDest = strings.TrimSpace(aux.Db2PdlSourceDest)
	configEntry.Options = make(map[string]interface{})
	if err := json.Unmarshal(data, &configEntry.Options); err != nil {
		return err
	}
	delete(configEntry.Options, "enabled")
	delete(configEntry.Options, "outputDir")
	delete(configEntry.Options, "db2PdlSourceDest")
	return nil
}

func (profile *Profile) UnmarshalJSON(data []byte) error {
	type rawProfile struct {
		Language  string         `json:"language"`
		Enabled   *bool          `json:"enabled"`
		Src       []string       `json:"src"`
		OutputDir string         `json:"outputDir"`
		Templates TemplateConfig `json:"templates"`
	}
	aux := rawProfile{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	profile.Language = strings.TrimSpace(aux.Language)
	if aux.Enabled != nil {
		profile.Enabled = *aux.Enabled
	} else {
		profile.Enabled = true
	}
	profile.Src = aux.Src
	profile.OutputDir = aux.OutputDir
	profile.Templates = aux.Templates
	profile.Options = make(map[string]interface{})
	if err := json.Unmarshal(data, &profile.Options); err != nil {
		return err
	}
	delete(profile.Options, "language")
	delete(profile.Options, "enabled")
	delete(profile.Options, "src")
	delete(profile.Options, "outputDir")
	delete(profile.Options, "templates")
	return nil
}

func (profile Profile) LanguageOrDefault(name string) string {
	lang := strings.TrimSpace(profile.Language)
	if lang == "" {
		return name
	}
	return lang
}

func (profile Profile) ConfigSettings() map[string]interface{} {
	settings := make(map[string]interface{}, len(profile.Options)+1)
	for key, value := range profile.Options {
		settings[key] = value
	}
	if _, exists := settings["enabled"]; !exists {
		settings["enabled"] = profile.Enabled
	} else {
		if enabledValue, ok := settings["enabled"].(bool); ok {
			settings["enabled"] = enabledValue || profile.Enabled
		} else {
			settings["enabled"] = profile.Enabled
		}
	}
	return settings
}

type Section struct {
	Name      string              `json:"name"`
	OutputDir string              `json:"outputDir"`
	Src       []string            `json:"src"`
	Files     map[string][]string `json:"files"`
}

type JSIndexConfig struct {
	Enabled    bool             `json:"enabled"`
	Filename   string           `json:"filename"`
	Template   string           `json:"template"`
	OutputDir  string           `json:"outputDir"`
	Namespaces JSNamespaceDepth `json:"namespaces"`
}

type JSNamespaceDepth struct {
	Depth int `json:"depth"`
}

type JSNamespacesConfig struct {
	Enabled   bool   `json:"enabled"`
	Filename  string `json:"filename"`
	OutputDir string `json:"outputDir"`
	Template  string `json:"template"`
}

type JSTypescriptConfig struct {
	Generate       bool   `json:"generate"`
	GenerateIndex  bool   `json:"generateIndex"`
	IndexFilename  string `json:"indexFilename"`
	BarrelFilename string `json:"barrelFilename"`
	ClassTemplate  string `json:"classTemplate"`
	IndexTemplate  string `json:"indexTemplate"`
	BarrelTemplate string `json:"barrelTemplate"`
	OutputDir      string `json:"outputDir"`
}

type JSConfig struct {
	Index        JSIndexConfig      `json:"index"`
	GlobalIndex  JSIndexConfig      `json:"globalIndex"`
	TemplatesDir string             `json:"templatesDir"`
	Dirs         []string           `json:"dirs"`
	Namespaces   JSNamespacesConfig `json:"namespaces"`
	Typescript   JSTypescriptConfig `json:"typescript"`
}

type RootConfig struct {
	CompanyName      string              `json:"companyName"`
	Project          string              `json:"project"`
	Version          string              `json:"version"`
	Verbose          bool                `json:"verbose"`
	Profiles         map[string]Profile  `json:"profiles"`
	Sections         []Section           `json:"sections"`
	Src              []string            `json:"src"`
	OutputDir        string              `json:"outputDir"`
	TempDir          string              `json:"tempDir"`
	Templates        TemplateConfig      `json:"templates"`
	JS               JSConfig            `json:"js"`
	Db2Pdl           Db2PdlConfig        `json:"db2pdl"`
	Db2PdlSourceDest string              `json:"db2PdlSourceDest"`
	ActiveSections   map[string]struct{} `json:"-"`
}
