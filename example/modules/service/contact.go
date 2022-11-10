package service

import (
	"gorm.io/gorm"

	model "github.com/shiqiyue/go-admin-gen/example"

	"context"

	"test/pkg/ferror"

	"test/pkg/gorms"
)

// ContactSrv 联系人-服务
type ContactSrv struct {

	// DB DB实例
	DB *gorm.DB `inject:""`
}

// Add 添加联系人
func (s *ContactSrv) Add(ctx context.Context, entity model.Contact) (model.Contact, error) {

	db := gorms.GetDb(ctx, s.Db)
	// 添加
	err := entity.Create(db)
	if err != nil {
		return nil, ferror.Wrap("添加联系人异常", err)
	}
	return entity, nil

}

// Edit 修改联系人
func (s *ContactSrv) Edit(ctx context.Context, entity model.Contact) error {

	db := gorms.GetDb(ctx, s.Db)
	// 修改
	err := db.Model(&model.Contact{}).Updates(entity)
	if err != nil {
		return ferror.Wrap("修改联系人异常", err)
	}
	return nil

}
