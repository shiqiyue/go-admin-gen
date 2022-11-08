package go_admin_gen

import (
	"github.com/iancoleman/strcase"
	"path"
	"reflect"
)

//go:generate go-options config
type config struct {
	// 模块名称
	ModuleName string
	// 要执行生成任务的model
	Models []*ModelConfig

	// 模块目录
	ModuleDir string
	// 模块包名
	ModulePackage string
	// dataloader相对路径
	DataloaderDir string
	// dto 相对路径
	DtoDir string
	// service 相对路径
	ServiceDir string
	// 模块的graphql 相对路径
	ModuleGraphqlDir string

	// Api 目录
	ApiDir string
	// Api包名
	ApiPackage string
	// Api graphql相对路径
	ApiGraphqlDir string
}

type ModelConfig struct {
	Model interface{}

	Name string
}

func (c *ModelConfig) GetModelName() string {
	return reflect.TypeOf(c).Name()
}

func (c *ModelConfig) GetModelNameToLowerCamel() string {
	return strcase.ToLowerCamel(c.GetModelName())
}

func (c *ModelConfig) GetModelNameToSnake() string {
	return strcase.ToSnake(c.GetModelName())
}

func (c config) GetDataloaderDir() string {
	return path.Join(c.ModuleDir, c.DataloaderDir)
}

func (c config) GetDataloaderPackage() string {
	return path.Join(c.ModulePackage, c.DataloaderDir)
}

func (c config) GetDtoDir() string {
	return path.Join(c.ModuleDir, c.DtoDir)
}

func (c config) GetDtoPackage() string {
	return path.Join(c.ModulePackage, c.DtoDir)
}

func (c config) GetServiceDir() string {
	return path.Join(c.ModuleDir, c.ServiceDir)
}

func (c config) GetServicePackage() string {
	return path.Join(c.ModulePackage, c.ServiceDir)
}

func (c config) GetModuleGraphqlDir() string {
	return path.Join(c.ModuleDir, c.ModuleGraphqlDir)
}

func (c config) GetModuleGraphqlPackage() string {
	return path.Join(c.ModulePackage, c.ModuleGraphqlDir)
}

func (c config) GetApiGraphqlDir() string {
	return path.Join(c.ModuleDir, c.ApiGraphqlDir)
}

func (c config) GetApiGraphqlPackage() string {
	return path.Join(c.ModulePackage, c.ApiGraphqlDir)
}
