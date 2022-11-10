package dto

import (
	"time"
)

// ContactGroupPageFilter 过滤联系人分组-入参
type ContactGroupPageFilter struct {
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

	// EmailEnable 是否启用邮箱通知
	EmailEnable *bool `qssql:"email_enable = ?"`

	// DingTalkEnable 是否启用钉钉推送
	DingTalkEnable *bool `qssql:"ding_talk_enable = ?"`

	// WebhookEnable 是否启用webhook
	WebhookEnable *bool `qssql:"webhook_enable = ?"`
}
