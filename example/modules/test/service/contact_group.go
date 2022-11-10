package service

import (
	"gorm.io/gorm"

	model "github.com/shiqiyue/go-admin-gen/example"

	"context"

	"github.com/shiqiyue/go-admin-gen/example/modules/test/dto"

	"test/pkg/ferror"

	"test/pkg/gorms"

	"test/pkg/qstool"
)

// ContactGroupSrv 联系人分组-服务
type ContactGroupSrv struct {

	// DB DB实例
	DB *gorm.DB `inject:""`
}

// Add 添加联系人分组
func (s *ContactGroupSrv) Add(ctx context.Context, entity *model.ContactGroup) (*model.ContactGroup, error) {

	db := gorms.GetDb(ctx, s.Db)
	// 添加
	err := entity.Create(db)
	if err != nil {
		return nil, ferror.Wrap("添加联系人分组异常", err)
	}
	return entity, nil

}

// Edit 修改联系人分组
func (s *ContactGroupSrv) Edit(ctx context.Context, entity *model.ContactGroup) error {

	db := gorms.GetDb(ctx, s.Db)
	// 修改
	err := db.Model(&model.ContactGroup{}).Updates(entity)
	if err != nil {
		return ferror.Wrap("修改联系人分组异常", err)
	}
	return nil

}

// GetById 通过ID获取联系人分组
func (s *ContactGroupSrv) GetById(ctx context.Context, id int64) (*model.ContactGroup, error) {

	db := gorms.GetDb(ctx, s.Db)

	r := &model.ContactGroup{}
	err := model.NewContactGroupQuerySet(db).IDEq(id).One(r)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, ferror.Wrap("获取联系人分组异常", err)
	}
	return r, nil

}

// List 联系人分组列表查询
func (s *ContactGroupSrv) List(ctx context.Context, query dto.ContactGroupQuery) ([]model.ContactGroup, int64, error) {

	db := gorms.GetDb(ctx, s.Db)
	qs := model.NewContactGroupQuerySet(db)
	qs = qs.With(qstool.ParseWhere(query.Filter))
	total, err := qs.Count()
	if err != nil {
		return nil, 0, ferror.Wrap("获取联系人分组数量异常", err)
	}
	rs := make([]*model.ContactGroup, 0)
	if query.PageNum != nil && query.PageSize != nil {
		qs = qs.With(qstool.ParsePage(*query.PageNum, *query.PageSize))
	}
	if query.SortKey != nil && query.Reverse != nil {
		qs = qs.With(qstool.ParseOrder(string(*query.SortKey), *query.Reverse))
	} else {
		qs = qs.OrderDescByID()
	}
	err = qs.All(&rs)
	if err != nil {
		return nil, 0, ferror.Wrap("获取联系人分组列表异常", err)
	}
	return rs, total, nil

}
