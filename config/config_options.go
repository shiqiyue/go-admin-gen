package config

// Code generated by github.com/launchdarkly/go-options.  DO NOT EDIT.

import "fmt"

import "github.com/google/go-cmp/cmp"

type ApplyOptionFunc func(c *Config) error

func (f ApplyOptionFunc) apply(c *Config) error {
	return f(c)
}

func NewConfig(options ...Option) (Config, error) {
	var c Config
	err := applyConfigOptions(&c, options...)
	return c, err
}

func applyConfigOptions(c *Config, options ...Option) error {
	for _, o := range options {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

type Option interface {
	apply(*Config) error
}

type optionModuleNameImpl struct {
	o string
}

func (o optionModuleNameImpl) apply(c *Config) error {
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

func (o optionModelsImpl) apply(c *Config) error {
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

type optionPkgPackageImpl struct {
	o string
}

func (o optionPkgPackageImpl) apply(c *Config) error {
	c.PkgPackage = o.o
	return nil
}

func (o optionPkgPackageImpl) Equal(v optionPkgPackageImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionPkgPackageImpl) String() string {
	name := "OptionPkgPackage"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionPkgPackage pkg包名
func OptionPkgPackage(o string) Option {
	return optionPkgPackageImpl{
		o: o,
	}
}

type optionModuleDirImpl struct {
	o string
}

func (o optionModuleDirImpl) apply(c *Config) error {
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

func (o optionModulePackageImpl) apply(c *Config) error {
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

func (o optionDataloaderDirImpl) apply(c *Config) error {
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

func (o optionDtoDirImpl) apply(c *Config) error {
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

func (o optionServiceDirImpl) apply(c *Config) error {
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

func (o optionModuleGraphqlDirImpl) apply(c *Config) error {
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

func (o optionApiDirImpl) apply(c *Config) error {
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

func (o optionApiGraphqlDirImpl) apply(c *Config) error {
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

type optionVueSrcDirImpl struct {
	o string
}

func (o optionVueSrcDirImpl) apply(c *Config) error {
	c.VueSrcDir = o.o
	return nil
}

func (o optionVueSrcDirImpl) Equal(v optionVueSrcDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionVueSrcDirImpl) String() string {
	name := "OptionVueSrcDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionVueSrcDir Vue Src 目录
func OptionVueSrcDir(o string) Option {
	return optionVueSrcDirImpl{
		o: o,
	}
}

type optionVueViewDirImpl struct {
	o string
}

func (o optionVueViewDirImpl) apply(c *Config) error {
	c.VueViewDir = o.o
	return nil
}

func (o optionVueViewDirImpl) Equal(v optionVueViewDirImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o optionVueViewDirImpl) String() string {
	name := "OptionVueViewDir"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

// OptionVueViewDir Vue view相对路径
func OptionVueViewDir(o string) Option {
	return optionVueViewDirImpl{
		o: o,
	}
}
