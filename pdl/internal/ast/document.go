package ast

import "encoding/json"

type Document struct {
	Version   int       `json:"version"`
	Source    Source    `json:"source"`
	Namespace Namespace `json:"namespace"`
}

type Source struct {
	File      string     `json:"file"`
	Namespace Identifier `json:"namespace"`
}

type Namespace struct {
	Name    Identifier   `json:"name"`
	Usings  []Identifier `json:"usings"`
	Classes []Class      `json:"classes"`
}

type Identifier struct {
	Segments      []string `json:"segments"`
	QualifiedName string   `json:"qualifiedName"`
}

type Class struct {
	Name           string      `json:"name"`
	AccessModifier string      `json:"accessModifier"`
	Parent         *Identifier `json:"parent"`
	Members        []Member    `json:"members"`
}

type Member struct {
	Kind           string          `json:"kind"`
	Name           string          `json:"name"`
	AccessModifier string          `json:"accessModifier"`
	TypePayload    json.RawMessage `json:"type,omitempty"`
	Attributes     []Attribute     `json:"attributes,omitempty"`
	Arguments      []Argument      `json:"arguments,omitempty"`
	Access         string          `json:"access,omitempty"`
	ReturnType     *Identifier     `json:"returnType,omitempty"`
	Value          interface{}     `json:"value,omitempty"`
}

type PropertyType struct {
	Type          Identifier `json:"type"`
	ArrayNotation []string   `json:"arrayNotation"`
}

type Argument struct {
	Name string     `json:"name"`
	Type Identifier `json:"type"`
}

type Attribute struct {
	Name   Identifier      `json:"name"`
	Params AttributeParams `json:"params"`
}

type AttributeParams struct {
	Required []interface{}         `json:"required"`
	Optional []AttributeNamedParam `json:"optional"`
}

type AttributeNamedParam struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
