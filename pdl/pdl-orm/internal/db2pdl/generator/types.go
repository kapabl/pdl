package generator

import "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"

type Renderer interface {
	Render(templatePath string, payload interface{}) (string, error)
}

type CodeWriter func(language string, fileName string, source string) error

type Generator interface {
	Generate(shared.TableData) error
	InjectAttributes(shared.TableData) shared.TableData
}
