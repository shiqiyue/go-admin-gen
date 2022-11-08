package core

import (
	"bytes"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"reflect"
	"strings"
)

type GenContext struct {
	T reflect.Type

	Name string

	Fields []FieldInfo
}

func (c *GenContext) GenModelSchema() string {
	schemaDocument := &ast.SchemaDocument{}
	c.genModel(schemaDocument)
	//c.genAddReq()
	//c.genEditReq()
	//c.genColumnEnum()
	//c.genBoolExp()
	//c.genOrderBy()
	//
	//c.genMuation()
	//c.genQuery()
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(schemaDocument)
	return buf.String()
}

func (c *GenContext) GenApiSchema() string {

	return ""
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
