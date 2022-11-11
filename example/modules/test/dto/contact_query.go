package dto

import (
	model "github.com/shiqiyue/go-admin-gen/example"
)

// ContactPageFilter 过滤联系人-入参
type ContactPageFilter struct {
	// Name 名称
	Name *string `qssql:"name like ?" qsformat:"%%%s%%"`
	// Email 邮箱
	Email *string `qssql:"email like ?" qsformat:"%%%s%%"`
	// Phone 手机号码
	Phone *string `qssql:"phone like ?" qsformat:"%%%s%%"`
}

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
