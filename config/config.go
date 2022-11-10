package config

import (
	"github.com/iancoleman/strcase"
	"path"
	"reflect"
	"strings"
)

//go:generate go-options Config
type Config struct {
	// 模块名称
	ModuleName string
	// 要执行生成任务的model
	Models []*ModelConfig

	// pkg包名
	PkgPackage string
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

	// Api graphql相对路径
	ApiGraphqlDir string
}

type ModelConfig struct {
	Model interface{}

	Name string
}

func (c *ModelConfig) GetModelName() string {
	t := reflect.ValueOf(c.Model).Elem().Type()
	return t.Name()
}

func (c *ModelConfig) GetModelNameToLowerCamel() string {
	return strcase.ToLowerCamel(c.GetModelName())
}

func (c *ModelConfig) GetModelNameToSnake() string {
	return strcase.ToSnake(c.GetModelName())
}

func (c *ModelConfig) GetModelNameWithModuleToSnake(moduleName string) string {
	modelName := c.GetModelName()
	return strings.ToLower(moduleName) + "_" + strcase.ToSnake(modelName)
}

func (c Config) GetDataloaderDir() string {
	return path.Join(c.ModuleDir, c.ModuleName, c.DataloaderDir)
}

func (c Config) GetDataloaderFullPackage() string {
	return path.Join(c.ModulePackage, c.ModuleName, c.DataloaderDir)
}

func (c Config) GetDataloaderPackage() string {
	return path.Base(c.GetDataloaderFullPackage())
}

func (c Config) GetDtoDir() string {
	return path.Join(c.ModuleDir, c.ModuleName, c.DtoDir)
}

func (c Config) GetDtoFullPackage() string {
	return path.Join(c.ModulePackage, c.ModuleName, c.DtoDir)
}

func (c Config) GetDtoPackage() string {
	return path.Base(c.GetDtoFullPackage())
}

func (c Config) GetServiceDir() string {
	return path.Join(c.ModuleDir, c.ModuleName, c.ServiceDir)
}

func (c Config) GetServiceFullPackage() string {
	return path.Join(c.ModulePackage, c.ModuleName, c.ServiceDir)
}

func (c Config) GetServicePackage() string {
	return path.Base(c.GetServiceFullPackage())
}

func (c Config) GetModuleGraphqlDir() string {
	return path.Join(c.ModuleDir, c.ModuleName, c.ModuleGraphqlDir)
}

func (c Config) GetModuleGraphqlFullPackage() string {
	return path.Join(c.ModulePackage, c.ModuleName, c.ModuleGraphqlDir)
}

func (c Config) GetModuleGraphqlPackage() string {
	return path.Base(c.GetModuleGraphqlFullPackage())
}

func (c Config) GetApiGraphqlDir() string {
	return path.Join(c.ApiDir, c.ApiGraphqlDir)
}
