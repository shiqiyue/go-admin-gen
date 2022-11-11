package model

import (
	go_admin_gen "github.com/shiqiyue/go-admin-gen"
	"github.com/shiqiyue/go-admin-gen/config"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	os.RemoveAll("D:\\project\\go-admin-gen\\example\\modules\\")
	os.RemoveAll("D:\\project\\go-admin-gen\\example\\api\\")

	err := go_admin_gen.Gen(
		config.OptionModuleName("test"),
		config.OptionPkgPackage("test/pkg"),
		config.OptionModels([]*config.ModelConfig{&config.ModelConfig{
			Model:         &Contact{},
			Name:          "联系人",
			DisableApiGen: true,
		}, &config.ModelConfig{
			Model: &ContactGroup{},
			Name:  "联系人分组",
		}}),
		config.OptionModuleDir("D:\\project\\go-admin-gen\\example\\modules\\"),
		config.OptionModulePackage("github.com/shiqiyue/go-admin-gen/example/modules"),
		config.OptionApiDir("D:\\project\\go-admin-gen\\example\\api\\"),
	)
	if err != nil {
		t.Error(err)
	}
}
