package core

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) addReqName() string {
	return c.modelName() + "AddInput"
}

func (c *GenContext) genAddReq(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.addReqName()
	def.Description = "添加" + c.Name + "参数"
	def.Directives = []*ast.Directive{c.modelDirective()}
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsAdd() {
			continue
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: field.GqlFieldName(),
			Type: &ast.Type{
				NamedType: field.Scalar(),
				NonNull:   !field.Nullable,
			},
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
