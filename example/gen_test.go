package model

import (
	go_admin_gen "github.com/shiqiyue/go-admin-gen"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	os.RemoveAll("D:\\project\\go-admin-gen\\example\\modules\\")
	err := go_admin_gen.Gen(go_admin_gen.OptionModuleName("test"),
		go_admin_gen.OptionModels([]*go_admin_gen.ModelConfig{&go_admin_gen.ModelConfig{
			Model: &Contact{},
			Name:  "联系人",
		}}),
		go_admin_gen.OptionModuleDir("D:\\project\\go-admin-gen\\example\\modules\\"),
		go_admin_gen.OptionModulePackage("test"),
		go_admin_gen.OptionApiDir("D:\\project\\go-admin-gen\\example\\api\\"),
	)
	if err != nil {
		t.Error(err)
	}
}
