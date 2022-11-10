package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/dto"
	"path"
)

func (c *GenContext) serviceName() string {
	return c.ModelName() + "Srv"
}

func (c *GenContext) genService() error {
	serviceModel := &dto.Model{
		Name:        c.serviceName(),
		ShortName:   "s",
		Description: fmt.Sprintf("%s-服务", c.Name),
		Fields:      make([]*dto.ModelField, 0),
		Methods:     make([]*dto.ModelMethod, 0),
	}
	serviceModel.Fields = append(serviceModel.Fields, &dto.ModelField{
		Name:        "DB",
		Description: "DB实例",
		Type:        "gorm.DB",
		Ptr:         true,
		Tag:         "`inject:\"\"`",
	})

	serviceImports := append(defaultImports, c.fullModelPath())
	serviceImports = append(serviceImports, "context")
	serviceImports = append(serviceImports, c.Cfg.GetDtoFullPackage())
	serviceImports = append(serviceImports, path.Join(c.Cfg.PkgPackage, "ferror"))
	serviceImports = append(serviceImports, path.Join(c.Cfg.PkgPackage, "gorms"))
	serviceImports = append(serviceImports, path.Join(c.Cfg.PkgPackage, "qstool"))

	serviceModel.Methods = append(serviceModel.Methods, c.genServiceAddMethod())
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceEditMethod())
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceGetByIdMethod())
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceListMethod())

	err := c.writeModel([]*dto.Model{serviceModel}, c.Cfg.GetServicePackage(), path.Join(c.Cfg.GetServiceDir(), fmt.Sprintf("%s.go", c.ModelSneakName())), serviceImports)
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) genServiceAddMethod() *dto.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)
	// 添加
	err := entity.Create(db)
	if err != nil {
		return nil, ferror.Wrap("添加%s异常", err)
	}
	return entity, nil
`, c.Name)
	return &dto.ModelMethod{
		Name:        "Add",
		Description: "添加" + c.Name,
		Body:        body,
		Args: []*dto.ModelMethodArg{
			&dto.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&dto.ModelMethodArg{
				Name: "entity",
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
		},
		Results: []*dto.ModelMethodResult{
			&dto.ModelMethodResult{
				Name: "",
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
			&dto.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceEditMethod() *dto.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)
	// 修改
	err := db.Model(&model.%s{}).Updates(entity)
	if err != nil {
		return ferror.Wrap("修改%s异常", err)
	}
	return nil
`, c.ModelName(), c.Name)
	return &dto.ModelMethod{
		Name:        "Edit",
		Description: "修改" + c.Name,
		Body:        body,
		Args: []*dto.ModelMethodArg{
			&dto.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&dto.ModelMethodArg{
				Name: "entity",
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
		},
		Results: []*dto.ModelMethodResult{
			&dto.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceGetByIdMethod() *dto.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)

	r := &model.%s{}
	err := model.New%sQuerySet(db).IDEq(id).One(r)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, ferror.Wrap("获取%s异常", err)
	}
	return r, nil
`, c.ModelName(), c.ModelName(), c.Name)
	return &dto.ModelMethod{
		Name:        "GetById",
		Description: "通过ID获取" + c.Name,
		Body:        body,
		Args: []*dto.ModelMethodArg{
			&dto.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&dto.ModelMethodArg{
				Name: "id",
				Type: "int64",
				Ptr:  false,
			},
		},
		Results: []*dto.ModelMethodResult{
			&dto.ModelMethodResult{
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
			&dto.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceListMethod() *dto.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)
	qs := model.New%sQuerySet(db)
	qs = qs.With(qstool.ParseWhere(query.Filter))
	total, err := qs.Count()
	if err != nil {
		return nil, 0, ferror.Wrap("获取%s数量异常", err)
	}
	rs := make([]*model.%s, 0)
	if query.PageNum != nil && query.PageSize != nil {
		qs = qs.With(qstool.ParsePage(*query.PageNum, *query.PageSize))
	}
	if query.SortKey != nil && query.Reverse != nil {
		qs = qs.With(qstool.ParseOrder(string(*query.SortKey), *query.Reverse))
	} else {
		qs = qs.OrderDescByID()
	}
	err = qs.All(&rs)
	if err != nil {
		return nil, 0, ferror.Wrap("获取%s列表异常", err)
	}
	return rs, total, nil
`, c.ModelName(), c.Name, c.ModelName(), c.Name)
	return &dto.ModelMethod{
		Name:        "List",
		Description: fmt.Sprintf("%s列表查询", c.Name),
		Body:        body,
		Args: []*dto.ModelMethodArg{
			&dto.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&dto.ModelMethodArg{
				Name: "query",
				Type: fmt.Sprintf("dto." + c.queryDtoName()),
				Ptr:  false,
			},
		},
		Results: []*dto.ModelMethodResult{
			&dto.ModelMethodResult{
				Type: fmt.Sprintf("[]*model.%s", c.ModelName()),
				Ptr:  false,
			},
			&dto.ModelMethodResult{
				Type: "int64",
				Ptr:  false,
			},
			&dto.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}
