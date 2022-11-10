package dto

import (
	"time"
)

// ContactGroupPageFilter 过滤联系人分组-入参
type ContactGroupPageFilter struct {
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

	// EmailEnable 是否启用邮箱通知
	EmailEnable *bool

	// DingTalkEnable 是否启用钉钉推送
	DingTalkEnable *bool

	// WebhookEnable 是否启用webhook
	WebhookEnable *bool
}
