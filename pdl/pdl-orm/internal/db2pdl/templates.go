package db2pdl

import (
    "bytes"
    "encoding/json"
    "fmt"
    "path/filepath"
    "sync"
    "text/template"

    "io/fs"
    "os"
    "strings"
    "unicode"
)

type TemplateUtils struct {
	config  RootConfig
	source  fs.FS
	cache   map[string]*template.Template
	mutex   sync.Mutex
	funcMap template.FuncMap
}

func NewTemplateUtils(config RootConfig, templatesDir string, useExternal bool) (*TemplateUtils, error) {
	var source fs.FS
	if useExternal {
		fileSystem := os.DirFS(templatesDir)
		source = fileSystem
	} else {
		sub, subErr := fs.Sub(embeddedTemplates, "templates")
		if subErr == nil {
			source = sub
		} else {
			source = embeddedTemplates
		}
	}
	result := &TemplateUtils{
		config: config,
		source: source,
		cache:  make(map[string]*template.Template),
		funcMap: template.FuncMap{
			"braceWrap": func(value interface{}) string {
				return "{" + fmt.Sprint(value) + "}"
			},
			"not": func(input bool) bool { return !input },
            "trimSuffix": func(value string, suffix string) string {
                return strings.TrimSuffix(value, suffix)
            },
            "lowerFirst": func(value string) string {
                if value == "" {
                    return value
                }
                runes := []rune(value)
                runes[0] = unicode.ToLower(runes[0])
                return string(runes)
            },
        },
	}
	return result, nil
}

func (utils *TemplateUtils) Render(templatePath string, payload interface{}) (string, error) {
	utils.mutex.Lock()
	compiled, found := utils.cache[templatePath]
	utils.mutex.Unlock()
	if !found {
		filePath := filepath.ToSlash(templatePath + ".gotmpl")
		content, readErr := fs.ReadFile(utils.source, filePath)
		if readErr != nil {
			return "", readErr
		}
		tmpl, compileErr := template.New(filePath).Funcs(utils.funcMap).Parse(string(content))
		if compileErr != nil {
			return "", compileErr
		}
		utils.mutex.Lock()
		utils.cache[templatePath] = tmpl
		utils.mutex.Unlock()
		compiled = tmpl
	}
	context, contextErr := utils.mergeData(payload)
	if contextErr != nil {
		return "", contextErr
	}
	var buffer bytes.Buffer
	execErr := compiled.Execute(&buffer, context)
	if execErr != nil {
		return "", execErr
	}
	return buffer.String(), nil
}

func (utils *TemplateUtils) mergeData(payload interface{}) (map[string]interface{}, error) {
	base := map[string]interface{}{
		"companyName": utils.config.CompanyName,
		"project":     utils.config.Project,
		"version":     utils.config.Version,
		"CompanyName": utils.config.CompanyName,
		"Project":     utils.config.Project,
		"Version":     utils.config.Version,
	}
	if payload == nil {
		return base, nil
	}
	raw, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return nil, marshalErr
	}
	var extra map[string]interface{}
	if err := json.Unmarshal(raw, &extra); err != nil {
		return nil, err
	}
	for key, value := range extra {
		base[key] = value
	}
	base["Data"] = payload
	base["data"] = payload
	return base, nil
}
