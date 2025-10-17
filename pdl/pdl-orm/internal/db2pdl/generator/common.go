package generator

import (
	"strings"

	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

type AttributeConfig struct {
	PrimaryTemplate string
	ColumnTemplate  string
}

type AttributeSetter func(*shared.FieldInfo, string)

func ApplyAttributes(table shared.TableData, config AttributeConfig, setter AttributeSetter) shared.TableData {
	result := table
	result.FieldsInfo = make([]shared.FieldInfo, len(table.FieldsInfo))
	copy(result.FieldsInfo, table.FieldsInfo)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		tags := make([]string, 0, 2)
		if config.PrimaryTemplate != "" && field.IsPrimaryKey {
			tags = append(tags, shared.GenerateAttribute(*field, config.PrimaryTemplate))
		}
		if config.ColumnTemplate != "" && shared.CantConvertBackAndForth(field.Original) {
			tags = append(tags, shared.GenerateAttribute(*field, config.ColumnTemplate))
		}
		setter(field, strings.Join(tags, "\n"))
	}
	return result
}

func CloneTable(table shared.TableData) shared.TableData {
	result := table
	result.FieldsInfo = make([]shared.FieldInfo, len(table.FieldsInfo))
	copy(result.FieldsInfo, table.FieldsInfo)
	return result
}
