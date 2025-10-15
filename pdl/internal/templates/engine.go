package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"
	"text/template"

	"io/fs"
)

type Engine struct {
	data    map[string]interface{}
	source  fs.FS
	cache   map[string]*template.Template
	mutex   sync.Mutex
	funcMap template.FuncMap
}

func NewEngine(baseData map[string]interface{}) *Engine {
	result := &Engine{
		data:   baseData,
		source: FS,
		cache:  make(map[string]*template.Template),
		funcMap: template.FuncMap{
			"braceWrap": func(value interface{}) string {
				return "{" + fmt.Sprint(value) + "}"
			},
			"not": func(input bool) bool { return !input },
		},
	}
	return result
}

func (engine *Engine) Render(templatePath string, payload interface{}) (string, error) {
	tmpl, compileErr := engine.loadTemplate(templatePath)
	if compileErr != nil {
		return "", compileErr
	}
	merged, mergeErr := engine.mergeData(payload)
	if mergeErr != nil {
		return "", mergeErr
	}
	var buffer bytes.Buffer
	if execErr := tmpl.Execute(&buffer, merged); execErr != nil {
		return "", execErr
	}
	return buffer.String(), nil
}

func (engine *Engine) loadTemplate(templatePath string) (*template.Template, error) {
	engine.mutex.Lock()
	cached, found := engine.cache[templatePath]
	engine.mutex.Unlock()
	if found {
		return cached, nil
	}
	path := filepath.ToSlash(templatePath + ".gotmpl")
	content, readErr := fs.ReadFile(engine.source, path)
	if readErr != nil {
		return nil, readErr
	}
	compiled, compileErr := template.New(path).Funcs(engine.funcMap).Parse(string(content))
	if compileErr != nil {
		return nil, compileErr
	}
	engine.mutex.Lock()
	engine.cache[templatePath] = compiled
	engine.mutex.Unlock()
	return compiled, nil
}

func (engine *Engine) mergeData(payload interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{}, len(engine.data))
	for key, value := range engine.data {
		result[key] = value
	}
	if payload == nil {
		return result, nil
	}
	raw, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return result, marshalErr
	}
	var extra map[string]interface{}
	unmarshalErr := json.Unmarshal(raw, &extra)
	if unmarshalErr != nil {
		return result, unmarshalErr
	}
	for key, value := range extra {
		result[key] = value
	}
	return result, nil
}
