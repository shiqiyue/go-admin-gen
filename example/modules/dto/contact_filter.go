package dto

import (
	"time"
)

// ContactPageFilter 过滤联系人-入参
type ContactPageFilter struct {
	// CreatedAtMin 创建时间-最小值
	CreatedAtMin *time.Time

	// CreatedAtMax 创建时间-最大值
	CreatedAtMax *time.Time

	// UpdatedAtMin 更新时间-最小值
	UpdatedAtMin *time.Time

	// UpdatedAtMax 更新时间-最大值
	UpdatedAtMax *time.Time

	// Name 名称
	Name *string

	// Email 邮箱
	Email *string

	// Phone 手机号码
	Phone *string
}
