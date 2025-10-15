package generator

import (
	"context"
	"fmt"
	"sync"

	"github.com/kapablanka/pdl/pdl/internal/ast"
)

type Options struct {
	GeneratorName string
	ASTPath       string
	OutputDir     string
	TemplateDir   string
	ConfigPath    string
}

type Generator interface {
	Name() string
	Generate() error
}

type Factory func(context.Context, ast.Document, Options) Generator

var (
	registryMutex sync.RWMutex
	registryMap   = map[string]Factory{}
)

func Register(name string, factory Factory) {
	if name == "" {
		panic("generator name cannot be empty")
	}
	if factory == nil {
		panic(fmt.Sprintf("generator %s factory cannot be nil", name))
	}
	registryMutex.Lock()
	defer registryMutex.Unlock()
	if _, exists := registryMap[name]; exists {
		panic(fmt.Sprintf("generator %s already registered", name))
	}
	registryMap[name] = factory
}

func Create(options Options) (Generator, error) {
	var result Generator
	document, loadErr := ast.LoadDocumentFile(options.ASTPath)
	if loadErr != nil {
		return result, loadErr
	}
	registryMutex.RLock()
	factory := registryMap[options.GeneratorName]
	registryMutex.RUnlock()
	if factory == nil {
		return result, fmt.Errorf("generator %s not registered", options.GeneratorName)
	}
	instance := factory(context.Background(), document, options)
	return instance, nil
}

func Names() []string {
	result := make([]string, 0)
	registryMutex.RLock()
	for name := range registryMap {
		result = append(result, name)
	}
	registryMutex.RUnlock()
	return result
}
