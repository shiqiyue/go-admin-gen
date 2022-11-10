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
	// 分页查询
	listQueryArguments := []*ast.ArgumentDefinition{
		NewListAndItemNotNullArgument(c.columnEnumName(), "distinct_on", ""),
		NewArgument(SCALAR_INT, "limit", ""),
		NewArgument(SCALAR_INT, "offset", ""),
		NewListAndItemNotNullArgument(c.orderByName(), "order_by", ""),
		NewArgument(c.boolExpName(), "where", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: fmt.Sprintf("%s分页", c.Name),
		Name:        c.modelSneakName() + "s",
		Arguments:   listQueryArguments,
		Type:        ast.NonNullListType(ast.NonNullNamedType(c.modelName(), nil), nil),
		Directives:  nil,
		Position:    nil,
	})

	// 数量查询
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: "数量查询",
		Name:        c.modelSneakName() + "_count",
		Arguments:   listQueryArguments,
		Type:        NewNotNullType(SCALAR_INT),
	})

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

	SchemaDocument.Extensions = append(SchemaDocument.Extensions, def)

}
