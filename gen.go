package go_admin_gen

import (
	"errors"
	"github.com/shiqiyue/go-admin-gen/core"
	"github.com/shiqiyue/go-admin-gen/util"
	"os"
	"path"
)

func Gen(options ...Option) error {
	// 创建配置
	cfg, err := newConfig(options...)
	if err != nil {
		return err
	}
	// 检查配置信息，设置默认配置信息
	err = checkConfig(&cfg)
	if err != nil {
		return err
	}
	// 创建目录
	err = mkdirs(&cfg)
	if err != nil {
		return err
	}
	for _, modelConfig := range cfg.Models {
		genCtx := core.Resolve(modelConfig.Model, modelConfig.Name)
		modelGraphql := genCtx.GenModelSchema()
		filePath := path.Join(cfg.GetModuleGraphqlDir(), modelConfig.GetModelNameToSnake()+".graphql")
		err := util.WriteFile([]byte(modelGraphql), filePath, false)
		if err != nil {
			return err
		}
	}

	return nil
}

// checkConfig 检查配置
func checkConfig(cfg *config) error {
	if cfg.ModuleName == "" {
		return errors.New("ModuleName can not be empty")
	}
	if len(cfg.Models) == 0 {
		return errors.New("Models can not be empty")
	}
	if cfg.ModuleDir == "" {
		return errors.New("ModuleDir can not be empty")
	}
	if cfg.ModulePackage == "" {
		return errors.New("ModulePackage can not be empty")
	}
	if cfg.ApiDir == "" {
		return errors.New("ApiDir can not be empty")
	}
	if cfg.ApiPackage == "" {
		return errors.New("ApiPackage can not be empty")
	}
	if cfg.DataloaderDir == "" {
		cfg.DataloaderDir = "dataloader"
	}
	if cfg.DtoDir == "" {
		cfg.DtoDir = "dto"
	}
	if cfg.ServiceDir == "" {
		cfg.ServiceDir = "service"
	}
	if cfg.ModuleGraphqlDir == "" {
		cfg.ModuleGraphqlDir = "graphql"
	}
	if cfg.ApiGraphqlDir == "" {
		cfg.ApiGraphqlDir = "schema" + string(os.PathSeparator) + cfg.ModuleName
	}
	return nil
}

// mkdirs 创建目录
func mkdirs(cfg *config) error {
	err := util.EnsureDirExist(cfg.GetDataloaderDir())
	if err != nil {
		return err
	}
	err = util.EnsureDirExist(cfg.GetDtoDir())
	if err != nil {
		return err
	}
	err = util.EnsureDirExist(cfg.GetServiceDir())
	if err != nil {
		return err
	}
	err = util.EnsureDirExist(cfg.GetModuleGraphqlDir())
	if err != nil {
		return err
	}
	err = util.EnsureDirExist(cfg.GetApiGraphqlDir())
	if err != nil {
		return err
	}

	return nil
}
