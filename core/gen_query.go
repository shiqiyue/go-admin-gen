package core

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
)

func (c *GenContext) genQuery(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{
		Kind: ast.Object,
		Name: "Query",
	}
	def.Fields = make([]*ast.FieldDefinition, 0)

	// 主键查询
	pkArguments := []*ast.ArgumentDefinition{
		NewNotNullArgument(SCALAR_INT64, "id", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: c.Name,
		Name:        c.modelSneakName(),
		Arguments:   pkArguments,
		Type:        NewType(c.modelName()),
	})

	// 分页查询
	listQueryArguments := []*ast.ArgumentDefinition{
		NewNotNullArgument(c.pageInputName(), "data", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: fmt.Sprintf("%s分页", c.Name),
		Name:        c.modelSneakName() + "s",
		Arguments:   listQueryArguments,
		Type:        ast.NonNullNamedType(c.pageResultName(), nil),
		Directives:  nil,
		Position:    nil,
	})

	SchemaDocument.Extensions = append(SchemaDocument.Extensions, def)

}
