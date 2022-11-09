package core

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) columnEnumName() string {
	return c.modelName() + "Column"

}

func (c *GenContext) genColumnEnum(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.Enum
	def.Name = c.columnEnumName()
	def.EnumValues = make([]*ast.EnumValueDefinition, 0)
	for _, field := range c.Fields {
		def.EnumValues = append(def.EnumValues, &ast.EnumValueDefinition{
			Name: field.DBFieldName(),
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
