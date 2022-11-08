package core

import (
	"reflect"
)

func Resolve(m interface{}, name string) *GenContext {
	t := reflect.ValueOf(m).Elem().Type()
	context := &GenContext{
		T:    t,
		Name: name,
	}

	context.resolveType(t)
	return context
}
