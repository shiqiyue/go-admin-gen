package go_admin_gen

// Code generated by github.com/launchdarkly/go-options.  DO NOT EDIT.

import "fmt"

import "github.com/google/go-cmp/cmp"

type ApplyOptionFunc func(c *config) error

func (f ApplyOptionFunc) apply(c *config) error {
	return f(c)
}

func newConfig(options ...Option) (config, error) {
	var c config
	err := applyConfigOptions(&c, options...)
	return c, err
}

func applyConfigOptions(c *config, options ...Option) error {
	for _, o := range options {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

type Option interface {
	apply(*config) error
}

type optionModuleNameImpl struct {
	o string
}

func (o optionModuleNameImpl) apply(c *config) error {
	c.ModuleName = o.o
	return nil
}

func (o optionModuleNameImpl) Equal(v optionModuleNameImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionModuleNameImpl) String() string {
	name := "OptionModuleName"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionModuleName 模块名称
func OptionModuleName(o string) Option {
	return optionModuleNameImpl{
		o: o,
	}
}

type optionModelsImpl struct {
	o []*ModelConfig
}

func (o optionModelsImpl) apply(c *config) error {
	c.Models = o.o
	return nil
}

func (o optionModelsImpl) Equal(v optionModelsImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionModelsImpl) String() string {
	name := "OptionModels"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionModels 要执行生成任务的model
func OptionModels(o []*ModelConfig) Option {
	return optionModelsImpl{
		o: o,
	}
}

type optionModuleDirImpl struct {
	o string
}

func (o optionModuleDirImpl) apply(c *config) error {
	c.ModuleDir = o.o
	return nil
}

func (o optionModuleDirImpl) Equal(v optionModuleDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionModuleDirImpl) String() string {
	name := "OptionModuleDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionModuleDir 模块目录
func OptionModuleDir(o string) Option {
	return optionModuleDirImpl{
		o: o,
	}
}

type optionModulePackageImpl struct {
	o string
}

func (o optionModulePackageImpl) apply(c *config) error {
	c.ModulePackage = o.o
	return nil
}

func (o optionModulePackageImpl) Equal(v optionModulePackageImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionModulePackageImpl) String() string {
	name := "OptionModulePackage"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionModulePackage 模块包名
func OptionModulePackage(o string) Option {
	return optionModulePackageImpl{
		o: o,
	}
}

type optionDataloaderDirImpl struct {
	o string
}

func (o optionDataloaderDirImpl) apply(c *config) error {
	c.DataloaderDir = o.o
	return nil
}

func (o optionDataloaderDirImpl) Equal(v optionDataloaderDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionDataloaderDirImpl) String() string {
	name := "OptionDataloaderDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionDataloaderDir dataloader相对路径
func OptionDataloaderDir(o string) Option {
	return optionDataloaderDirImpl{
		o: o,
	}
}

type optionDtoDirImpl struct {
	o string
}

func (o optionDtoDirImpl) apply(c *config) error {
	c.DtoDir = o.o
	return nil
}

func (o optionDtoDirImpl) Equal(v optionDtoDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionDtoDirImpl) String() string {
	name := "OptionDtoDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionDtoDir dto 相对路径
func OptionDtoDir(o string) Option {
	return optionDtoDirImpl{
		o: o,
	}
}

type optionServiceDirImpl struct {
	o string
}

func (o optionServiceDirImpl) apply(c *config) error {
	c.ServiceDir = o.o
	return nil
}

func (o optionServiceDirImpl) Equal(v optionServiceDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionServiceDirImpl) String() string {
	name := "OptionServiceDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionServiceDir service 相对路径
func OptionServiceDir(o string) Option {
	return optionServiceDirImpl{
		o: o,
	}
}

type optionModuleGraphqlDirImpl struct {
	o string
}

func (o optionModuleGraphqlDirImpl) apply(c *config) error {
	c.ModuleGraphqlDir = o.o
	return nil
}

func (o optionModuleGraphqlDirImpl) Equal(v optionModuleGraphqlDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionModuleGraphqlDirImpl) String() string {
	name := "OptionModuleGraphqlDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionModuleGraphqlDir 模块的graphql 相对路径
func OptionModuleGraphqlDir(o string) Option {
	return optionModuleGraphqlDirImpl{
		o: o,
	}
}

type optionApiDirImpl struct {
	o string
}

func (o optionApiDirImpl) apply(c *config) error {
	c.ApiDir = o.o
	return nil
}

func (o optionApiDirImpl) Equal(v optionApiDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionApiDirImpl) String() string {
	name := "OptionApiDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionApiDir Api 目录
func OptionApiDir(o string) Option {
	return optionApiDirImpl{
		o: o,
	}
}

type optionApiGraphqlDirImpl struct {
	o string
}

func (o optionApiGraphqlDirImpl) apply(c *config) error {
	c.ApiGraphqlDir = o.o
	return nil
}

func (o optionApiGraphqlDirImpl) Equal(v optionApiGraphqlDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionApiGraphqlDirImpl) String() string {
	name := "OptionApiGraphqlDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionApiGraphqlDir Api graphql相对路径
func OptionApiGraphqlDir(o string) Option {
	return optionApiGraphqlDirImpl{
		o: o,
	}
}
