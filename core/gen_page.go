package core

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

func (c *GenContext) pageInputName() string {
	return fmt.Sprintf("%sPageInput", c.modelName())
}

func (c *GenContext) pageFilterName() string {
	return fmt.Sprintf("%sPageFilter", c.modelName())
}

func (c GenContext) sortKeyEnumName() string {
	return fmt.Sprintf("%sSortKeys", c.modelName())
}

func (c *GenContext) genSortKey(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.Enum
	def.Name = c.sortKeyEnumName()
	def.Description = c.Name + " 排序"
	def.EnumValues = make([]*ast.EnumValueDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsSortKey() {
			continue
		}
		def.EnumValues = append(def.EnumValues, &ast.EnumValueDefinition{
			Description: field.Description(),
			Name:        strings.ToUpper(strcase.ToSnake(field.Name)),
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)

}

func (c *GenContext) genPageFilter(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.pageFilterName()
	def.Description = c.Name + "分页过滤参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsFilter() {
			continue
		}
		scalar := field.Scalar()
		if scalar == SCALAR_TIME {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description() + "-最小值",
				Name:        field.GqlFieldName() + "Min",
				Type:        ast.NamedType(scalar, nil),
			})
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description() + "-最大值",
				Name:        field.GqlFieldName() + "Max",
				Type:        ast.NamedType(scalar, nil),
			})
		}
		if scalar == SCALAR_STRING {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description(),
				Name:        field.GqlFieldName(),
				Type:        ast.NamedType(scalar, nil),
			})
		}
		if scalar == SCALAR_BOOLEAN {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description(),
				Name:        field.GqlFieldName(),
				Type:        ast.NamedType(scalar, nil),
			})
		}
		if scalar == SCALAR_INT32 || scalar == SCALAR_INT || scalar == SCALAR_INT64 {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description(),
				Name:        field.GqlFieldName() + "s",
				Type:        ast.ListType(&ast.Type{NamedType: scalar, NonNull: false}, nil),
			})
		}
	}

	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}

func (c *GenContext) genPageInput(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.pageInputName()
	def.Directives = []*ast.Directive{c.goModelDirective(c.queryDtoFullName())}
	def.Description = c.Name + "分页参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "pageNum",
		Description:  "页数",
		Type:         ast.NonNullNamedType(SCALAR_INT, nil),
		DefaultValue: &ast.Value{Kind: ast.IntValue, Raw: "1"},
		Directives:   []*ast.Directive{c.validateDirective("min=1,max=1000", "页数")},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "pageSize",
		Description:  "分页大小",
		Type:         ast.NonNullNamedType(SCALAR_INT, nil),
		DefaultValue: &ast.Value{Kind: ast.IntValue, Raw: "10"},
		Directives:   []*ast.Directive{c.validateDirective("min=1,max=200", "分页大小")},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:        "filter",
		Description: "过滤条件",
		Type:        ast.NamedType(c.pageFilterName(), nil),
		Directives:  []*ast.Directive{c.goModelDirective(c.filterDtoFullName())},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "reverse",
		Description:  "排序方向；true:asc;false:desc",
		Type:         ast.NamedType(SCALAR_BOOLEAN, nil),
		DefaultValue: &ast.Value{Kind: ast.BooleanValue, Raw: "false"},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "sortKey",
		Description:  "排序方向；true:asc;false:desc",
		Type:         ast.NamedType(c.sortKeyEnumName(), nil),
		DefaultValue: &ast.Value{Kind: ast.EnumValue, Raw: "ID"},
	})
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
