package core

import (
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
)

func (c *GenContext) modelName() string {
	return c.T.Name()
}

func (c *GenContext) modelSneakName() string {
	return strcase.ToSnake(c.modelName())
}

func (c *GenContext) fullModelName() string {
	return c.T.PkgPath() + "." + c.modelName()

}

func (c GenContext) modelDirective() *ast.Directive {
	return &ast.Directive{
		Name: "goModel",
		Arguments: []*ast.Argument{&ast.Argument{
			Name: "model",
			Value: &ast.Value{
				Raw:  c.fullModelName(),
				Kind: ast.StringValue,
			},
		}},
	}
}

func (c *GenContext) genModel(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Name = c.modelName()
	def.Description = c.Name
	def.Directives = []*ast.Directive{c.modelDirective()}
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {

		namedType := field.Scalar()
		isArray := field.IsArray()
		t := &ast.Type{}
		if isArray {
			t = ast.NonNullListType(&ast.Type{
				NamedType: namedType,
				NonNull:   !field.Nullable,
			}, nil)
		} else {
			t = &ast.Type{
				NamedType: namedType,
				NonNull:   !field.Nullable,
			}
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name:        field.GqlFieldName(),
			Description: field.Description(),
			Type:        t,
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
