package core

import (
	"bytes"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/shiqiyue/go-admin-gen/util"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"path"
	"strings"
)

func (c *GenContext) graphqlAddReqName() string {
	return c.graphqlModelName() + "AddInput"
}

func (c *GenContext) graphqlEditReqName() string {
	return c.graphqlModelName() + "EditInput"
}

func (c *GenContext) graphqlPageInputName() string {
	return fmt.Sprintf("%sPageInput", c.graphqlModelName())
}

func (c *GenContext) graphqlPageResultName() string {
	return fmt.Sprintf("%sPageResult", c.graphqlModelName())
}

func (c *GenContext) graphqlPageFilterName() string {
	return fmt.Sprintf("%sPageFilter", c.graphqlModelName())
}

func (c GenContext) graphqlSortKeyEnumName() string {
	return fmt.Sprintf("%sSortKeys", c.graphqlModelName())
}

func (c *GenContext) graphqlRemoveReqName() string {
	return c.graphqlModelName() + "RemovesInput"
}

func (c *GenContext) GenGraphqlApiSchema() error {
	schemaDocument := &ast.SchemaDocument{}
	c.genGraphqlAddReq(schemaDocument)
	c.genGraphqlEditReq(schemaDocument)
	c.genGraphqlRemoveReq(schemaDocument)
	c.genGraphqlSortKey(schemaDocument)
	c.genGraphqlPageFilter(schemaDocument)
	c.genGraphqlPageInput(schemaDocument)
	c.genGraphqlPageResult(schemaDocument)
	c.genGraphqlMuation(schemaDocument)
	c.genGraphqlQuery(schemaDocument)
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

func (c *GenContext) genGraphqlMuation(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{
		Kind: ast.Object,
		Name: "Mutation",
	}
	def.Fields = make([]*ast.FieldDefinition, 0)
	// 添加
	addArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "data",
		Type: ast.NonNullNamedType(c.graphqlAddReqName(), nil),
	}}
	addMutation := &ast.FieldDefinition{
		Description: "添加" + c.Name,
		Name:        "add" + c.graphqlModelName(),
		Arguments:   addArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, addMutation)
	// 修改
	editArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "data",
		Type: ast.NonNullNamedType(c.graphqlEditReqName(), nil),
	}}
	editMutation := &ast.FieldDefinition{
		Description: "修改" + c.Name,
		Name:        "edit" + c.graphqlModelName(),
		Arguments:   editArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, editMutation)
	// 删除
	removeArguments := []*ast.ArgumentDefinition{&ast.ArgumentDefinition{
		Name: "data",
		Type: ast.NonNullNamedType(c.graphqlRemoveReqName(), nil),
	}}
	removeMutation := &ast.FieldDefinition{
		Description: "删除" + c.Name,
		Name:        fmt.Sprintf("remove%ss", c.graphqlModelName()),
		Arguments:   removeArguments,
		Type:        NewNotNullType(SCALAR_BOOLEAN),
	}
	def.Fields = append(def.Fields, removeMutation)

	SchemaDocument.Extensions = append(SchemaDocument.Extensions, def)
}

func (c *GenContext) genGraphqlQuery(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{
		Kind: ast.Object,
		Name: "Query",
	}
	def.Fields = make([]*ast.FieldDefinition, 0)

	// 主键查询
	pkArguments := []*ast.ArgumentDefinition{
		NewNotNullArgument(SCALAR_INT64, "id", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: c.Name,
		Name:        c.graphqlModelSneakName(),
		Arguments:   pkArguments,
		Type:        NewType(c.graphqlModelName()),
	})

	// 分页查询
	listQueryArguments := []*ast.ArgumentDefinition{
		NewNotNullArgument(c.graphqlPageInputName(), "data", ""),
	}
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Description: fmt.Sprintf("%s分页", c.Name),
		Name:        c.graphqlModelSneakName() + "s",
		Arguments:   listQueryArguments,
		Type:        ast.NonNullNamedType(c.graphqlPageResultName(), nil),
		Directives:  nil,
		Position:    nil,
	})

	SchemaDocument.Extensions = append(SchemaDocument.Extensions, def)

}

func (c *GenContext) genGraphqlAddReq(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.graphqlAddReqName()
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
			Description: field.Description(),
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}

func (c *GenContext) genGraphqlEditReq(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.graphqlEditReqName()
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

func (c *GenContext) genGraphqlSortKey(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.Enum
	def.Name = c.graphqlSortKeyEnumName()
	def.Description = c.Name + " 排序"
	def.EnumValues = make([]*ast.EnumValueDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsSortKey() {
			continue
		}
		def.EnumValues = append(def.EnumValues, &ast.EnumValueDefinition{
			Description: field.Description(),
			Name:        strings.ToUpper(strcase.ToSnake(field.Name)),
		})
	}
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)

}

func (c *GenContext) genGraphqlPageFilter(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.graphqlPageFilterName()
	def.Description = c.Name + "分页过滤参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	for _, field := range c.Fields {
		if !field.IsFilter() {
			continue
		}
		scalar := field.Scalar()
		if scalar == SCALAR_TIME {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description() + "-最小值",
				Name:        field.GqlFieldName() + "Min",
				Type:        ast.NamedType(scalar, nil),
			})
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description() + "-最大值",
				Name:        field.GqlFieldName() + "Max",
				Type:        ast.NamedType(scalar, nil),
			})
		}
		if scalar == SCALAR_STRING {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description(),
				Name:        field.GqlFieldName(),
				Type:        ast.NamedType(scalar, nil),
			})
		}
		if scalar == SCALAR_BOOLEAN {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description(),
				Name:        field.GqlFieldName(),
				Type:        ast.NamedType(scalar, nil),
			})
		}
		if scalar == SCALAR_INT32 || scalar == SCALAR_INT || scalar == SCALAR_INT64 {
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Description: field.Description(),
				Name:        field.GqlFieldName() + "s",
				Type:        ast.ListType(&ast.Type{NamedType: scalar, NonNull: false}, nil),
			})
		}
	}

	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}

func (c *GenContext) genGraphqlPageInput(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.graphqlPageInputName()
	def.Directives = []*ast.Directive{c.goModelDirective(c.queryDtoFullName())}
	def.Description = c.Name + "分页参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "pageNum",
		Description:  "页数",
		Type:         ast.NonNullNamedType(SCALAR_INT, nil),
		DefaultValue: &ast.Value{Kind: ast.IntValue, Raw: "1"},
		Directives:   []*ast.Directive{c.validateDirective("min=1,max=1000", "页数")},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "pageSize",
		Description:  "分页大小",
		Type:         ast.NonNullNamedType(SCALAR_INT, nil),
		DefaultValue: &ast.Value{Kind: ast.IntValue, Raw: "10"},
		Directives:   []*ast.Directive{c.validateDirective("min=1,max=200", "分页大小")},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:        "filter",
		Description: "过滤条件",
		Type:        ast.NamedType(c.graphqlPageFilterName(), nil),
		Directives:  []*ast.Directive{c.goModelDirective(c.filterDtoFullName())},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "reverse",
		Description:  "排序方向；true:asc;false:desc",
		Type:         ast.NamedType(SCALAR_BOOLEAN, nil),
		DefaultValue: &ast.Value{Kind: ast.BooleanValue, Raw: "false"},
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:         "sortKey",
		Description:  "排序字段",
		Type:         ast.NamedType(c.graphqlSortKeyEnumName(), nil),
		DefaultValue: &ast.Value{Kind: ast.EnumValue, Raw: "ID"},
	})
	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}

func (c *GenContext) genGraphqlPageResult(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Name = c.graphqlPageResultName()
	def.Description = c.Name + "分页-结果"
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:        "records",
		Description: "记录",
		Type:        ast.NonNullListType(NewType(c.graphqlModelName()), nil),
	})
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:        "total",
		Description: "总数",
		Type:        ast.NonNullNamedType(SCALAR_INT64, nil),
	})

	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}

func (c *GenContext) genGraphqlRemoveReq(SchemaDocument *ast.SchemaDocument) {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Name = c.graphqlRemoveReqName()
	def.Description = "删除" + c.Name + "参数"
	def.Fields = make([]*ast.FieldDefinition, 0)
	def.Fields = append(def.Fields, &ast.FieldDefinition{
		Name:        "ids",
		Type:        ast.NonNullListType(&ast.Type{NamedType: SCALAR_INT64, NonNull: true}, nil),
		Description: "ID列表",
	})

	SchemaDocument.Definitions = append(SchemaDocument.Definitions, def)
}
