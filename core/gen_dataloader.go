package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/templates"
	"path"
)

func (c *GenContext) dataloaderName() string {
	return c.ModelName() + "Loader"
}

func (c *GenContext) genDataloader() error {
	dataloaderModel := &templates.Model{
		Name:        c.dataloaderName(),
		ShortName:   "l",
		Description: fmt.Sprintf("%s-dataloader", c.Name),
		Fields:      make([]*templates.ModelField, 0),
		Methods:     make([]*templates.ModelMethod, 0),
		Remarks:     make([]string, 0),
	}

	dataloaderModel.Remarks = append(dataloaderModel.Remarks, fmt.Sprintf("go:generate dataloaden %sPkLoader int64 *%s", c.ModelName(), c.fullModelName()))

	dataloaderModel.Fields = append(dataloaderModel.Fields, &templates.ModelField{
		Name:        "Db",
		Description: "DB 实例",
		Type:        "gorm.DB",
		Ptr:         true,
		Tag:         "`inject:\"\"`",
	})
	dataloaderModel.Fields = append(dataloaderModel.Fields, &templates.ModelField{
		Name:        "pkLoader",
		Description: "主键Loader",
		Type:        fmt.Sprintf("dataloaders.LoadHelper[%sPkLoader]", c.ModelName()),
		Ptr:         false,
	})

	dataloaderModel.Methods = append(dataloaderModel.Methods, c.genDataloaderSetupMethod())
	dataloaderModel.Methods = append(dataloaderModel.Methods, c.genDataloaderPkLoader())

	dataloaderImports := append(defaultImports, c.Cfg.GetDtoFullPackage())
	dataloaderImports = append(defaultImports, path.Join(c.Cfg.PkgPackage, "gqlgens/dataloaders"))

	// 这里忽略输出model的错误，因为使用了泛型,goimports会报错
	err := c.writeModel([]*templates.Model{dataloaderModel}, c.Cfg.GetDataloaderPackage(), path.Join(c.Cfg.GetDataloaderDir(), fmt.Sprintf("%s.go", c.ModelSneakName())), dataloaderImports, true)
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) genDataloaderSetupMethod() *templates.ModelMethod {
	body := fmt.Sprintf(`
	l.GetPkLoader()
`)
	return &templates.ModelMethod{
		Name:        "SetUp",
		Description: "初始化Dataloader",
		Body:        body,
	}
}

func (c *GenContext) genDataloaderPkLoader() *templates.ModelMethod {
	body := fmt.Sprintf(`
	return l.pkLoader.Get(func() *%sPkLoader {
		return New%sPkLoader(%sPkLoaderConfig{
			Fetch: dataloaders.GenLoadIn(l.Db.Unscoped(), func(data *model.%s) int64 {
				return data.ID
			}),
			Wait:     10 * time.Millisecond,
			MaxBatch: 100,
		})
	})
`, c.ModelName(), c.ModelName(), c.ModelName(), c.ModelName())
	return &templates.ModelMethod{
		Name:        "GetPkLoader",
		Description: "获取主键Loader",
		Body:        body,
		Results: []*templates.ModelMethodResult{
			&templates.ModelMethodResult{
				Type: fmt.Sprintf("%sPkLoader", c.ModelName()),
				Ptr:  true,
			},
		},
	}
}
