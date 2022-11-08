package core

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) removeReqName() string {
	return c.modelName() + "RemoveInput"
}

func (c *GenContext) genRemoveReq(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.removeReqName()
	def.Description = "删除" + c.Name + "参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name: "ids",
		Type: ast.NonNullListType(&ast.Type{NamedType: SCALAR_INT64, NonNull: true}, nil),
	})

	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
