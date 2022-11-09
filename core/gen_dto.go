package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/dto"
	"github.com/shiqiyue/go-admin-gen/core/templates"
	"github.com/shiqiyue/go-admin-gen/util"
	"path"
)

func (c *GenContext) editReqDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "/" + c.editReqDtoName()
}

func (c *GenContext) editReqDtoName() string {
	return c.modelName() + "EditDto"
}

func (c *GenContext) addReqDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "/" + c.addReqDtoName()
}

func (c *GenContext) addReqDtoName() string {
	return c.modelName() + "EditDto"
}

func (c *GenContext) genReqDto() error {

	addReqDtoModel := &dto.Model{
		Name:        fmt.Sprintf("%sAddDto", c.modelName()),
		Description: fmt.Sprintf("添加%s-入参", c.Name),
		Fields:      make([]*dto.ModelField, 0),
	}
	editReqDtoModel := &dto.Model{
		Name:        fmt.Sprintf("%sEditDto", c.modelName()),
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

	err := c.outputModel(addReqDtoModel, "dto", path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("add_%s_dto.go", c.modelSneakName())))
	if err != nil {
		return err
	}
	err = c.outputModel(editReqDtoModel, "dto", path.Join(c.Cfg.GetDtoDir(), fmt.Sprintf("edit_%s_dto.go", c.modelSneakName())))
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) outputModel(m *dto.Model, pack string, filePath string) error {
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
