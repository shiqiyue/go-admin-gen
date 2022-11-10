package core

import (
	"bytes"
	"github.com/iancoleman/strcase"
	"github.com/shiqiyue/go-admin-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"path"
)

func (c *GenContext) graphqlModelName() string {
	return c.T.Name()
}

func (c *GenContext) ModelName() string {
	return c.T.Name()

}

func (c *GenContext) ModelSneakName() string {
	return strcase.ToSnake(c.ModelName())
}

func (c *GenContext) graphqlModelSneakName() string {
	return strcase.ToSnake(c.graphqlModelName())
}

func (c *GenContext) fullModelName() string {
	return c.T.PkgPath() + "." + c.ModelName()

}

func (c *GenContext) fullModelPath() string {
	return c.T.PkgPath()

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

func (c GenContext) goModelDirective(model string) *ast.Directive {
	return &ast.Directive{
		Name: "goModel",
		Arguments: []*ast.Argument{&ast.Argument{
			Name: "model",
			Value: &ast.Value{
				Raw:  model,
				Kind: ast.StringValue,
			},
		}},
	}
}

func (c GenContext) validateDirective(rules string, name string) *ast.Directive {
	return &ast.Directive{
		Name: "validate",
		Arguments: []*ast.Argument{&ast.Argument{
			Name: "rules",
			Value: &ast.Value{
				Raw:  rules,
				Kind: ast.StringValue,
			},
		}, &ast.Argument{
			Name: "name",
			Value: &ast.Value{
				Raw:  name,
				Kind: ast.StringValue,
			},
		}},
	}
}

func (c *GenContext) genGraphqlModel(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Name = c.graphqlModelName()
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

func (c *GenContext) GenModelSchema() error {
	schemaDocument := &ast.SchemaDocument{}
	c.genGraphqlModel(schemaDocument)
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(schemaDocument)
	filePath := path.Join(c.Cfg.GetModuleGraphqlDir(), c.ModelCfg.GetModelNameWithModuleToSnake(c.Cfg.ModuleName)+".graphql")
	err := util.WriteFile([]byte(buf.String()), filePath, false)
	if err != nil {
		return err
	}
	return nil
}
