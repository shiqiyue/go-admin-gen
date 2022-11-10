package dto

import (
	"time"

	model "github.com/shiqiyue/go-admin-gen/example"
)

// ContactPageFilter 过滤联系人-入参
type ContactPageFilter struct {

	// CreatedAtMin 创建时间-最小值
	CreatedAtMin *time.Time `qssql:"created_at >= ?"`

	// CreatedAtMax 创建时间-最大值
	CreatedAtMax *time.Time `qssql:"created_at <= ?"`

	// UpdatedAtMin 更新时间-最小值
	UpdatedAtMin *time.Time `qssql:"updated_at >= ?"`

	// UpdatedAtMax 更新时间-最大值
	UpdatedAtMax *time.Time `qssql:"updated_at <= ?"`

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
