package core

import (
	"fmt"
	"github.com/shiqiyue/go-admin-gen/core/dto"
	"path"
)

func (c *GenContext) dataloaderName() string {
	return c.ModelName() + "Loader"
}

func (c *GenContext) genDataloader() error {
	dataloaderModel := &dto.Model{
		Name:        c.dataloaderName(),
		ShortName:   "l",
		Description: fmt.Sprintf("%s-dataloader", c.Name),
		Fields:      make([]*dto.ModelField, 0),
		Methods:     make([]*dto.ModelMethod, 0),
		Remarks:     make([]string, 0),
	}

	dataloaderModel.Remarks = append(dataloaderModel.Remarks, fmt.Sprintf("go:generate go run github.com/shiqiyue/dataloaden %sPkLoader int64 *%s", c.ModelName(), c.fullModelName()))

	dataloaderModel.Fields = append(dataloaderModel.Fields, &dto.ModelField{
		Name:        "Db",
		Description: "DB 实例",
		Type:        "gorm.DB",
		Ptr:         true,
		Tag:         "`inject:\"\"`",
	})
	dataloaderModel.Fields = append(dataloaderModel.Fields, &dto.ModelField{
		Name:        "pkLoader",
		Description: "主键Loader",
		Type:        "dataloaders.LoadHelper[AlarmPkLoader]",
		Ptr:         false,
	})

	dataloaderModel.Methods = append(dataloaderModel.Methods, c.genDataloaderSetupMethod())
	dataloaderModel.Methods = append(dataloaderModel.Methods, c.genDataloaderPkLoader())

	dataloaderImports := append(defaultImports, c.Cfg.GetDtoFullPackage())
	dataloaderImports = append(defaultImports, path.Join(c.Cfg.PkgPackage, "gqlgens/dataloaders"))

	err := c.writeModel([]*dto.Model{dataloaderModel}, c.Cfg.GetDataloaderPackage(), path.Join(c.Cfg.GetDataloaderDir(), fmt.Sprintf("%s.go", c.ModelSneakName())), dataloaderImports)
	if err != nil {
		return err
	}

	return nil
}

func (c *GenContext) genDataloaderSetupMethod() *dto.ModelMethod {
	body := fmt.Sprintf(`
	l.GetPkLoader()
`)
	return &dto.ModelMethod{
		Name:        "SetUp",
		Description: "初始化Dataloader",
		Body:        body,
	}
}

func (c *GenContext) genDataloaderPkLoader() *dto.ModelMethod {
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
	return &dto.ModelMethod{
		Name:        "GetPkLoader",
		Description: "获取主键Loader",
		Body:        body,
		Results: []*dto.ModelMethodResult{
			&dto.ModelMethodResult{
				Type: fmt.Sprintf("%sPkLoader", c.ModelName()),
				Ptr:  true,
			},
		},
	}
}
