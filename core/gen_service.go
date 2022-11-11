package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/templates"
	"path"
)

func (c *GenContext) serviceName() string {
	return c.ModelName() + "Srv"
}

func (c *GenContext) genService() error {
	serviceModel := &templates.Model{
		Name:        c.serviceName(),
		ShortName:   "s",
		Description: fmt.Sprintf("%s-服务", c.Name),
		Fields:      make([]*templates.ModelField, 0),
		Methods:     make([]*templates.ModelMethod, 0),
	}
	serviceModel.Fields = append(serviceModel.Fields, &templates.ModelField{
		Name:        "Db",
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
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceRemovesMethod())
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceGetByIdMethod())
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceListMethod())
	serviceModel.Methods = append(serviceModel.Methods, c.genServiceAllMethod())

	err := c.writeModel([]*templates.Model{serviceModel}, c.Cfg.GetServicePackage(), path.Join(c.Cfg.GetServiceDir(), fmt.Sprintf("%s.go", c.ModelSneakName())), serviceImports, true)
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) genServiceAddMethod() *templates.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)
	// 添加
	err := entity.Create(db)
	if err != nil {
		return nil, ferror.Wrap("添加%s异常", err)
	}
	return entity, nil
`, c.Name)
	return &templates.ModelMethod{
		Name:        "Add",
		Description: "添加" + c.Name,
		Body:        body,
		Args: []*templates.ModelMethodArg{
			&templates.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&templates.ModelMethodArg{
				Name: "entity",
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
		},
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Name: "",
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
			&templates.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceEditMethod() *templates.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)
	// 修改
	err := db.Model(&model.%s{}).Updates(entity).Error
	if err != nil {
		return ferror.Wrap("修改%s异常", err)
	}
	return nil
`, c.ModelName(), c.Name)
	return &templates.ModelMethod{
		Name:        "Edit",
		Description: "修改" + c.Name,
		Body:        body,
		Args: []*templates.ModelMethodArg{
			&templates.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&templates.ModelMethodArg{
				Name: "entity",
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
		},
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceRemovesMethod() *templates.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)

	err := model.New%sQuerySet(db).IDIn(ids...).Delete()
	if err != nil {
		return ferror.Wrap("删除%s异常", err)
	}
	return nil
`, c.ModelName(), c.Name)
	return &templates.ModelMethod{
		Name:        "RemoveByIds",
		Description: "删除" + c.Name,
		Body:        body,
		Args: []*templates.ModelMethodArg{
			&templates.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&templates.ModelMethodArg{
				Name: "ids",
				Type: "[]int64",
				Ptr:  false,
			},
		},
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceGetByIdMethod() *templates.ModelMethod {
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
	return &templates.ModelMethod{
		Name:        "GetById",
		Description: "通过ID获取" + c.Name,
		Body:        body,
		Args: []*templates.ModelMethodArg{
			&templates.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&templates.ModelMethodArg{
				Name: "id",
				Type: "int64",
				Ptr:  false,
			},
		},
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Type: "model." + c.ModelName(),
				Ptr:  true,
			},
			&templates.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceListMethod() *templates.ModelMethod {
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
	return &templates.ModelMethod{
		Name:        "List",
		Description: fmt.Sprintf("%s列表查询", c.Name),
		Body:        body,
		Args: []*templates.ModelMethodArg{
			&templates.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
			&templates.ModelMethodArg{
				Name: "query",
				Type: fmt.Sprintf("dto." + c.queryDtoName()),
				Ptr:  false,
			},
		},
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Type: fmt.Sprintf("[]*model.%s", c.ModelName()),
				Ptr:  false,
			},
			&templates.ModelMethodResult{
				Type: "int64",
				Ptr:  false,
			},
			&templates.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}

func (c *GenContext) genServiceAllMethod() *templates.ModelMethod {
	body := fmt.Sprintf(`
	db := gorms.GetDb(ctx, s.Db)

	rs := make([]*model.%s, 0)
	err := model.New%sQuerySet(db).OrderDescByID().All(&rs)
	if err != nil {
		return nil, ferror.Wrap("获取%s异常", err)
	}
	return rs, nil
`, c.ModelName(), c.ModelName(), c.Name)
	return &templates.ModelMethod{
		Name:        "All",
		Description: fmt.Sprintf("获取所有%s", c.Name),
		Body:        body,
		Args: []*templates.ModelMethodArg{
			&templates.ModelMethodArg{
				Name: "ctx",
				Type: "context.Context",
				Ptr:  false,
			},
		},
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Type: fmt.Sprintf("[]*model.%s", c.ModelName()),
				Ptr:  false,
			},
			&templates.ModelMethodResult{
				Name: "",
				Type: "error",
				Ptr:  false,
			},
		},
	}

}
