package model

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

//go:generate goqueryset -in $GOFILE
// gen:qs
// ContactGroup 联系人组
type ContactGroup struct {
	ID             int64          `gorm:"primarykey; comment:ID"`
	CreatedAt      time.Time      `gorm:"not null; comment:创建时间"`
	UpdatedAt      time.Time      `gorm:"not null; comment:更新时间"`
	DeletedAt      gorm.DeletedAt `gorm:"index; comment:删除时间"`
	Name           string         `gorm:"not null; size:255; comment:名称"`
	EmailEnable    bool           `gorm:"not null; comment:是否启用邮箱通知"`
	EmailConfig    datatypes.JSON `gorm:"comment:邮箱配置"`
	DingTalkEnable bool           `gorm:"not null; comment:是否启用钉钉推送"`
	DingTalkConfig datatypes.JSON `gorm:"comment:钉钉推送配置"`
	WebhookEnable  bool           `gorm:"not null; comment:是否启用webhook"`
	WebhookConfig  datatypes.JSON `gorm:"comment:webhook配置"`
	Type           int            `gorm:"comment:类型"`
}

func (p *ContactGroup) TableName() string {
	return "nc_contact_group"
}

func (c *ContactGroup) GetEmailConfig() (*EmailConfig, error) {
	if c.EmailConfig != nil {
		r := &EmailConfig{}
		err := json.Unmarshal(c.EmailConfig, r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return nil, nil
}

func (c *ContactGroup) GetDingTalkConfig() (*DingTalkConfig, error) {
	if c.DingTalkConfig != nil {
		r := &DingTalkConfig{}
		err := json.Unmarshal(c.DingTalkConfig, r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return nil, nil
}

func (c *ContactGroup) GetWebhookConfig() (*WebhookConfig, error) {
	if c.WebhookConfig != nil {
		r := &WebhookConfig{}
		err := json.Unmarshal(c.WebhookConfig, r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return nil, nil
}

// EmailConfig 邮箱配置信息
type EmailConfig struct {
	// smtp邮件服务器地址
	SmtpServerAddress string `json:"smtpServerAddress"`
	// 邮箱账号
	Account string `json:"account"`
	// 邮箱密码
	Password string `json:"password"`
	// 是否开启tls
	Tls bool `json:"tls"`
	// 发件人信息
	Sender string `json:"sender"`
}

// DingTalkConfig 钉钉推送配置
type DingTalkConfig struct {
	// access token
	AccessToken string `json:"accessToken"`
	// 加签密钥
	SignSecret string `json:"signSecret"`
}

// WebhookConfig webhook配置
type WebhookConfig struct {
	Items []WebhookItemConfig `json:"items"`
}

// WebhookItemConfig webhook配置项
type WebhookItemConfig struct {
	// 链接
	Url string `json:"url"`
	// 头部
	Headers map[string]interface{} `json:"headers"`
}
