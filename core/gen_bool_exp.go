package core

import "github.com/vektah/gqlparser/v2/ast"

func (c *GenContext) boolExpName() string {
	return c.modelName() + "BoolExp"
}

func (c *GenContext) genBoolExp(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.boolExpName()
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:       "_and",
		Type:       NewListType(c.boolExpName()),
		Directives: nil,
		Position:   nil,
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:       "_not",
		Type:       NewType(c.boolExpName()),
		Directives: nil,
		Position:   nil,
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:       "_or",
		Type:       NewListType(c.boolExpName()),
		Directives: nil,
		Position:   nil,
	})
	for _, field := range c.Fields {
		if !field.IsWhere() {
			continue
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: field.GqlFieldName(),
			Type: NewType(field.Scalar() + "ComparisonExp"),
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
