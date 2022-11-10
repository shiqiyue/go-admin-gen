package dto

import (
	model "github.com/shiqiyue/go-admin-gen/example"
)

// ContactQuery 查询联系人-入参
type ContactQuery struct {
	// PageNum 第几页
	PageNum *int

	// PageSize 每页几条记录
	PageSize *int

	// Filter 过滤条件
	Filter *ContactPageFilter

	// Reverse 排序方向; true:asc, false:desc
	Reverse *bool

	// SortKey 排序字段
	SortKey *model.ContactDBSchemaField
}
