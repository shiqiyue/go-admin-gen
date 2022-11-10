package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/dto"
	"github.com/shiqiyue/go-admin-gen/core/templates"
	"github.com/shiqiyue/go-admin-gen/util"
	"path"
)

var defaultImports = []string{
	"encoding/json",
	"gorm.io/datatypes",
	"gorm.io/gorm",
	"time",
}

func (c *GenContext) editDtoFullName() string {
	return c.Cfg.GetDtoFullPackage() + "." + c.editDtoName()
}

func (c *GenContext) editDtoName() string {
	return c.graphqlModelName() + "EditDto"
}

func (c *GenContext) addDtoFullName() string {
	return c.Cfg.GetDtoFullPackage() + "." + c.addDtoName()
}

func (c *GenContext) addDtoName() string {
	return c.graphqlModelName() + "AddDto"
}

func (c *GenContext) queryDtoFullName() string {
	return c.Cfg.GetDtoFullPackage() + "." + c.queryDtoName()
}

func (c *GenContext) queryDtoName() string {
	return c.graphqlModelName() + "Query"
}

func (c *GenContext) filterDtoFullName() string {
	return c.Cfg.GetDtoFullPackage() + "." + c.filterDtoName()
}

func (c *GenContext) filterDtoName() string {
	return c.graphqlModelName() + "PageFilter"
}

func (c *GenContext) genDTO() error {

	addDtoModel := &dto.Model{
		Name:        c.addDtoName(),
		Description: fmt.Sprintf("添加%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}
	editDtoModel := &dto.Model{
		Name:        c.editDtoName(),
		Description: fmt.Sprintf("修改%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}
	filterDtoModel := &dto.Model{
		Name:        c.filterDtoName(),
		Description: fmt.Sprintf("过滤%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}
	queryDtoModel := &dto.Model{
		Name:        c.queryDtoName(),
		Description: fmt.Sprintf("查询%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}

	for _, field := range c.Fields {
		if field.IsAdd() {
			addDtoModel.Fields = append(addDtoModel.Fields, &dto.ModelField{
				Name:        field.GoFieldName(),
				Description: field.Description(),
				Type:        field.GoFieldType(),
				Ptr:         field.GoFieldPtr(),
			})
		}
		if field.IsEdit() {
			editDtoModel.Fields = append(editDtoModel.Fields, &dto.ModelField{
				Name:        field.GoFieldName(),
				Description: field.Description(),
				Type:        field.GoFieldType(),
				Ptr:         field.GoFieldPtr(),
			})
		}
		if field.IsFilter() {
			if !field.IsFilter() {
				continue
			}
			goType := field.Type
			if goType == "time.Time" {
				filterDtoModel.Fields = append(filterDtoModel.Fields, &dto.ModelField{
					Name:        field.GoFieldName() + "Min",
					Description: field.Description() + "-最小值",
					Type:        field.GoFieldType(),
					Ptr:         true,
					Tag:         fmt.Sprintf("`qssql:\"%s >= ?\"`", field.DBFieldName()),
				})
				filterDtoModel.Fields = append(filterDtoModel.Fields, &dto.ModelField{
					Name:        field.GoFieldName() + "Max",
					Description: field.Description() + "-最大值",
					Type:        field.GoFieldType(),
					Ptr:         true,
					Tag:         fmt.Sprintf("`qssql:\"%s <= ?\"`", field.DBFieldName()),
				})
			}
			if goType == "string" {
				filterDtoModel.Fields = append(filterDtoModel.Fields, &dto.ModelField{
					Name:        field.GoFieldName(),
					Description: field.Description(),
					Type:        field.GoFieldType(),
					Ptr:         true,
					Tag:         fmt.Sprintf("`qssql:\"%s like ?\" qsformat:\"%s\"`", field.DBFieldName(), "%%%s%%"),
				})
			}
			if goType == "bool" {
				filterDtoModel.Fields = append(filterDtoModel.Fields, &dto.ModelField{
					Name:        field.GoFieldName(),
					Description: field.Description(),
					Type:        field.GoFieldType(),
					Ptr:         true,
					Tag:         fmt.Sprintf("`qssql:\"%s = ?\"`", field.DBFieldName()),
				})
			}
			if goType == "int32" || goType == "int" || goType == "int64" {
				filterDtoModel.Fields = append(filterDtoModel.Fields, &dto.ModelField{
					Name:        field.GoFieldName() + "s",
					Description: field.Description(),
					Type:        "[]" + field.GoFieldType(),
					Ptr:         false,
					Tag:         fmt.Sprintf("`qssql:\"%s in ?\"`", field.DBFieldName()),
				})
			}

		}
	}
	queryDtoModel.Fields = append(queryDtoModel.Fields, &dto.ModelField{
		Name:        "PageNum",
		Description: "第几页",
		Type:        "int",
		Ptr:         true,
		Tag:         "",
	})
	queryDtoModel.Fields = append(queryDtoModel.Fields, &dto.ModelField{
		Name:        "PageSize",
		Description: "每页几条记录",
		Type:        "int",
		Ptr:         true,
		Tag:         "",
	})
	queryDtoModel.Fields = append(queryDtoModel.Fields, &dto.ModelField{
		Name:        "Filter",
		Description: "过滤条件",
		Type:        c.filterDtoName(),
		Ptr:         true,
		Tag:         "",
	})
	queryDtoModel.Fields = append(queryDtoModel.Fields, &dto.ModelField{
		Name:        "Reverse",
		Description: "排序方向; true:asc, false:desc",
		Type:        "bool",
		Ptr:         true,
		Tag:         "",
	})
	queryDtoModel.Fields = append(queryDtoModel.Fields, &dto.ModelField{
		Name:        "SortKey",
		Description: "排序字段",
		Type:        fmt.Sprintf("model.%sDBSchemaField", c.ModelName()),
		Ptr:         true,
		Tag:         "",
	})

	err := c.writeModel(addDtoModel, c.Cfg.GetDtoPackage(), path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("%s_add.go", c.graphqlModelSneakName())), defaultImports)
	if err != nil {
		return err
	}
	err = c.writeModel(editDtoModel, c.Cfg.GetDtoPackage(), path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("%s_edit.go", c.graphqlModelSneakName())), defaultImports)
	if err != nil {
		return err
	}
	err = c.writeModel(filterDtoModel, c.Cfg.GetDtoPackage(), path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("%s_filter.go", c.graphqlModelSneakName())), defaultImports)
	if err != nil {
		return err
	}
	queryImports := append(defaultImports, c.fullModelPath())
	err = c.writeModel(queryDtoModel, c.Cfg.GetDtoPackage(), path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("%s_query.go", c.graphqlModelSneakName())), queryImports)
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) writeModel(m *dto.Model, pack string, filePath string, inputs []string) error {
	templateData := make(map[string]interface{}, 0)
	templateData["PACKAGE"] = pack
	templateData["MODEL"] = m
	templateData["INPUTS"] = inputs

	r, err := util.DoTemplate(templates.MODEL, "test.go", templateData)
	if err != nil {
		return err
	}
	content := string(r)
	err = util.WriteFile([]byte(content), filePath, false)
	if err != nil {
		return err
	}
	if path.Ext(filePath) == ".go" {
		return util.RunInteractive(fmt.Sprintf("goimports -w %s", filePath))
	}
	return nil
}
