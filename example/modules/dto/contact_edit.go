package dto

// ContactEditDto 修改联系人-入参
type ContactEditDto struct {
	// Id ID
	Id int64

	// Name 名称
	Name string

	// Email 邮箱
	Email string

	// Phone 手机号码
	Phone string
}
