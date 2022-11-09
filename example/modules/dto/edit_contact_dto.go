package dto

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// ContactEditDto 修改联系人-入参
type ContactEditDto struct {
	// Id
	Id int64

	// Name 名称
	Name string

	// Email 邮箱
	Email string

	// Phone 手机号码
	Phone string
}
