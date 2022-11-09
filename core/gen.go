package core

import (
	go_admin_gen "github.com/shiqiyue/go-admin-gen"
	"reflect"
)

func Resolve(m interface{}, name string, cfg *go_admin_gen.Config) *GenContext {
	t := reflect.ValueOf(m).Elem().Type()
	context := &GenContext{
		T:    t,
		Name: name,
		Cfg:  cfg,
	}

	context.resolveType(t)
	return context
}
