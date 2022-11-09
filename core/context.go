package core

import (
	"bytes"
	"github.com/shiqiyue/go-admin-gen/config"
	"github.com/shiqiyue/go-admin-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"path"
	"reflect"
	"strings"
)

type GenContext struct {
	T reflect.Type

	Name string

	Fields []FieldInfo

	Cfg *config.Config

	ModelCfg *config.ModelConfig
}

func Resolve(m interface{}, name string, cfg *config.Config, config *config.ModelConfig) *GenContext {
	t := reflect.ValueOf(m).Elem().Type()
	context := &GenContext{
		T:        t,
		Name:     name,
		Cfg:      cfg,
		ModelCfg: config,
	}

	context.resolveType(t)
	return context
}

func (c *GenContext) Gen() error {
	err := c.GenModelSchema()
	if err != nil {
		return err
	}
	err = c.GenApiSchema()
	if err != nil {
		return err
	}
	err = c.genReqDto()
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) GenModelSchema() error {
	schemaDocument := &ast.SchemaDocument{}
	c.genModel(schemaDocument)
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

func (c *GenContext) GenApiSchema() error {
	schemaDocument := &ast.SchemaDocument{}
	c.genAddReq(schemaDocument)
	c.genEditReq(schemaDocument)
	c.genRemoveReq(schemaDocument)
	c.genMuation(schemaDocument)
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(schemaDocument)
	filePath := path.Join(c.Cfg.GetApiGraphqlDir(), c.ModelCfg.GetModelNameWithModuleToSnake(c.Cfg.ModuleName)+".graphql")
	err := util.WriteFile([]byte(c.betterGraphqlFormat(buf.String())), filePath, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *GenContext) betterGraphqlFormat(graphqlStr string) string {
	r := strings.ReplaceAll(graphqlStr, "}", "}\n")
	return r
}

func (c *GenContext) resolveType(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		fieldTag := field.Tag
		ignore := strings.Contains(fieldTag.Get("gen"), "ignore")
		if ignore {
			continue
		}
		switch fieldType.Kind() {
		case reflect.Struct:
			if fieldType.Name() == "Time" {
				c.addFieldInfo(field.Name, fieldType, false, fieldTag)
				continue
			}
			if fieldType.Name() == "DeletedAt" {
				continue
			}
			c.resolveType(fieldType)
		case reflect.Ptr:
			c.addFieldInfo(field.Name, fieldType.Elem(), true, fieldTag)
		default:
			c.addFieldInfo(field.Name, fieldType, false, fieldTag)
		}
	}
}

func (c *GenContext) addFieldInfo(fieldName string, t reflect.Type, nullable bool, tag reflect.StructTag) {
	c.Fields = append(c.Fields, FieldInfo{
		Name:     fieldName,
		Type:     t.String(),
		Nullable: nullable,
		Tag:      tag.Get("gen"),
		GormTag:  tag.Get("gorm"),
	})
}
