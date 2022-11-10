package service

import (
	"gorm.io/gorm"

	model "github.com/shiqiyue/go-admin-gen/example"

	"context"

	"test/pkg/ferror"

	"test/pkg/gorms"
)

// ContactGroupSrv 联系人分组-服务
type ContactGroupSrv struct {

	// DB DB实例
	DB *gorm.DB `inject:""`
}

// Add 添加联系人分组
func (s *ContactGroupSrv) Add(ctx context.Context, entity model.ContactGroup) (model.ContactGroup, error) {

	db := gorms.GetDb(ctx, s.Db)
	// 添加
	err := entity.Create(db)
	if err != nil {
		return nil, ferror.Wrap("添加联系人分组异常", err)
	}
	return entity, nil

}

// Edit 修改联系人分组
func (s *ContactGroupSrv) Edit(ctx context.Context, entity model.ContactGroup) error {

	db := gorms.GetDb(ctx, s.Db)
	// 修改
	err := db.Model(&model.ContactGroup{}).Updates(entity)
	if err != nil {
		return ferror.Wrap("修改联系人分组异常", err)
	}
	return nil

}
