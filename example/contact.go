package model

import (
	"gorm.io/gorm"
	"time"
)

//go:generate goqueryset -in $GOFILE
// gen:qs
// Contact 联系人
type Contact struct {
	ID        int64          `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null; size:255; comment:名称"`
	Email     string         `gorm:"not null; size:255; comment:邮箱"`
	Phone     string         `gorm:"not null; size: 255; comment:手机号码"`
}

func (p *Contact) TableName() string {
	return "nc_contact"
}
