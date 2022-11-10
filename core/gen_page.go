package core

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
)

func (c *GenContext) pageInputName() string {
	return fmt.Sprintf("%sPageInput", c.modelName())
}

func (c *GenContext) genPageFilter(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.pageInputName()
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
	def.Description = c.Name + "分页参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "pageNum",
		Type:         ast.NonNullListType(&ast.Type{NamedType: SCALAR_INT, NonNull: true}, nil),
		DefaultValue: &ast.Value{Kind: ast.IntValue, Raw: "1"},
		Directives:   []*ast.Directive{c.validateDirective("min=1,max=1000", "页数")},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "pageSize",
		Type:         ast.NonNullListType(&ast.Type{NamedType: SCALAR_INT, NonNull: true}, nil),
		DefaultValue: &ast.Value{Kind: ast.IntValue, Raw: "10"},
		Directives:   []*ast.Directive{c.validateDirective("min=1,max=200", "分页大小")},
	})

	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
