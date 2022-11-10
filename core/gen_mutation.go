package core

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
)

func (c *GenContext) genMuation(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{
		Kind: ast.Object,
		Name: "Mutation",
	}
	def.Fields = make([]*ast.FieldDefinition, 0)
	// 添加
	addArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "data",
		Type: ast.NonNullNamedType(c.addReqName(), nil),
	}}
	addMutation := &ast.FieldDefinition{
		Description: "添加" + c.Name,
		Name:        "add" + c.modelName(),
		Arguments:   addArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, addMutation)
	// 修改
	editArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "data",
		Type: ast.NonNullNamedType(c.editReqName(), nil),
	}}
	editMutation := &ast.FieldDefinition{
		Description: "修改" + c.Name,
		Name:        "edit" + c.modelName(),
		Arguments:   editArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, editMutation)
	// 删除
	removeArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "data",
		Type: ast.NonNullNamedType(c.removeReqName(), nil),
	}}
	removeMutation := &ast.FieldDefinition{
		Description: "删除" + c.Name,
		Name:        fmt.Sprintf("remove%ss", c.modelName()),
		Arguments:   removeArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, removeMutation)

	SchemaDocument.Extensions = append(SchemaDocument.Extensions, def)
}
