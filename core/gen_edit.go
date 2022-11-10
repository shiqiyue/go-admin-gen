package core

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) editReqName() string {
	return c.modelName() + "EditInput"
}

func (c *GenContext) genEditReq(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.editReqName()
	def.Description = "修改" + c.Name + "参数"
	def.Directives = []*ast.Directive{c.modelDirective()}
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsEdit() {
			continue
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: field.GqlFieldName(),
			Type: &ast.Type{
				NamedType: field.Scalar(),
				NonNull:   !field.Nullable,
			},
			Description: field.Description(),
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
