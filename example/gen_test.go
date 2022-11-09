package model

import (
	go_admin_gen "github.com/shiqiyue/go-admin-gen"
	"github.com/shiqiyue/go-admin-gen/config"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	os.RemoveAll("D:\\project\\go-admin-gen\\example\\modules\\")
	err := go_admin_gen.Gen(config.OptionModuleName("test"),
		config.OptionModels([]*config.ModelConfig{&config.ModelConfig{
			Model: &Contact{},
			Name:  "联系人",
		}}),
		config.OptionModuleDir("D:\\project\\go-admin-gen\\example\\modules\\"),
		config.OptionModulePackage("test"),
		config.OptionApiDir("D:\\project\\go-admin-gen\\example\\api\\"),
	)
	if err != nil {
		t.Error(err)
	}
}
