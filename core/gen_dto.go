package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/dto"
	"github.com/shiqiyue/go-admin-gen/core/templates"
	"github.com/shiqiyue/go-admin-gen/util"
	"path"
)

func (c *GenContext) editDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "." + c.editDtoName()
}

func (c *GenContext) editDtoName() string {
	return c.graphqlModelName() + "EditDto"
}

func (c *GenContext) addDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "." + c.addDtoName()
}

func (c *GenContext) addDtoName() string {
	return c.graphqlModelName() + "EditDto"
}

func (c *GenContext) queryDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "." + c.queryDtoName()
}

func (c *GenContext) queryDtoName() string {
	return c.graphqlModelName() + "Query"
}

func (c *GenContext) filterDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "." + c.filterDtoName()
}

func (c *GenContext) filterDtoName() string {
	return c.graphqlModelName() + "PageFilter"
}

func (c *GenContext) genDTO() error {

	addReqDtoModel := &dto.Model{
		Name:        fmt.Sprintf("%sAddDto", c.graphqlModelName()),
		Description: fmt.Sprintf("添加%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}
	editReqDtoModel := &dto.Model{
		Name:        fmt.Sprintf("%sEditDto", c.graphqlModelName()),
		Description: fmt.Sprintf("修改%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}
	for _, field := range c.Fields {
		if field.IsAdd() {
			addReqDtoModel.Fields = append(addReqDtoModel.Fields, &dto.ModelField{
				Name:        field.GoFieldName(),
				Description: field.Description(),
				Type:        field.GoFieldType(),
				Ptr:         field.GoFieldPtr(),
			})
		}
		if field.IsEdit() {
			editReqDtoModel.Fields = append(editReqDtoModel.Fields, &dto.ModelField{
				Name:        field.GoFieldName(),
				Description: field.Description(),
				Type:        field.GoFieldType(),
				Ptr:         field.GoFieldPtr(),
			})
		}
	}

	err := c.writeModel(addReqDtoModel, "dto", path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("add_%s_dto.go", c.graphqlModelSneakName())))
	if err != nil {
		return err
	}
	err = c.writeModel(editReqDtoModel, "dto", path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("edit_%s_dto.go", c.graphqlModelSneakName())))
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) writeModel(m *dto.Model, pack string, filePath string) error {
	templateData := make(map[string]interface{}, 0)
	templateData["PACKAGE"] = pack
	templateData["MODEL"] = m
	r, err := util.DoTemplate(templates.MODEL, "test.go", templateData)
	if err != nil {
		return err
	}
	content := string(r)
	err = util.WriteFile([]byte(content), filePath, false)
	if err != nil {
		return err
	}
	return nil
}
