package dto

import (
	"gorm.io/datatypes"

	"gorm.io/datatypes"
)

// ContactGroupEditDto 修改联系人分组-入参
type ContactGroupEditDto struct {
	// Id ID
	Id int64

	// Name 名称
	Name string

	// EmailEnable 是否启用邮箱通知
	EmailEnable bool

	// EmailConfig 邮箱配置
	EmailConfig datatypes.JSON

	// DingTalkEnable 是否启用钉钉推送
	DingTalkEnable bool

	// DingTalkConfig 钉钉推送配置
	DingTalkConfig datatypes.JSON

	// WebhookEnable 是否启用webhook
	WebhookEnable bool

	// WebhookConfig webhook配置
	WebhookConfig datatypes.JSON
}
