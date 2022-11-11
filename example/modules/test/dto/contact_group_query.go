package dto

import (
	model "github.com/shiqiyue/go-admin-gen/example"
)

// ContactGroupPageFilter 过滤联系人分组-入参
type ContactGroupPageFilter struct {
	// Name 名称
	Name *string `qssql:"name like ?" qsformat:"%%%s%%"`
	// EmailEnable 是否启用邮箱通知
	EmailEnable *bool `qssql:"email_enable = ?"`
	// DingTalkEnable 是否启用钉钉推送
	DingTalkEnable *bool `qssql:"ding_talk_enable = ?"`
	// WebhookEnable 是否启用webhook
	WebhookEnable *bool `qssql:"webhook_enable = ?"`
}

// ContactGroupQuery 查询联系人分组-入参
type ContactGroupQuery struct {
	// PageNum 第几页
	PageNum *int
	// PageSize 每页几条记录
	PageSize *int
	// Filter 过滤条件
	Filter *ContactGroupPageFilter
	// Reverse 排序方向; true:asc, false:desc
	Reverse *bool
	// SortKey 排序字段
	SortKey *model.ContactGroupDBSchemaField
}
