package service

import (
	"encoding/json"

	"gorm.io/datatypes"

	"gorm.io/gorm"

	"time"

	"github.com/shiqiyue/go-admin-gen/example"

	"context"

	"github.com/shiqiyue/go-admin-gen/example/modules/test/dto"

	"test/pkg/ferror"

	"test/pkg/gorms"

	"test/pkg/qstool"
)

// ContactSrv 联系人-服务
type ContactSrv struct {
	// Db DB实例
	Db *gorm.DB `inject:""`
}

// Add 添加联系人
func (s *ContactSrv) Add(ctx context.Context, entity *model.Contact) (*model.Contact, error) {

	db := gorms.GetDb(ctx, s.Db)
	// 添加
	err := entity.Create(db)
	if err != nil {
		return nil, ferror.Wrap("添加联系人异常", err)
	}
	return entity, nil

}

// Edit 修改联系人
func (s *ContactSrv) Edit(ctx context.Context, entity *model.Contact) error {

	db := gorms.GetDb(ctx, s.Db)
	// 修改
	err := db.Model(&model.Contact{ID: entity.ID}).Updates(entity).Error
	if err != nil {
		return ferror.Wrap("修改联系人异常", err)
	}
	return nil

}

// RemoveByIds 删除联系人
func (s *ContactSrv) RemoveByIds(ctx context.Context, ids []int64) error {

	db := gorms.GetDb(ctx, s.Db)

	err := model.NewContactQuerySet(db).IDIn(ids...).Delete()
	if err != nil {
		return ferror.Wrap("删除联系人异常", err)
	}
	return nil

}

// GetById 通过ID获取联系人
func (s *ContactSrv) GetById(ctx context.Context, id int64) (*model.Contact, error) {

	db := gorms.GetDb(ctx, s.Db)

	r := &model.Contact{}
	err := model.NewContactQuerySet(db).IDEq(id).One(r)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, ferror.Wrap("获取联系人异常", err)
	}
	return r, nil

}

// List 联系人列表查询
func (s *ContactSrv) List(ctx context.Context, query dto.ContactQuery) ([]*model.Contact, int64, error) {

	db := gorms.GetDb(ctx, s.Db)
	qs := model.NewContactQuerySet(db)
	qs = qs.With(qstool.ParseWhere(query.Filter))
	total, err := qs.Count()
	if err != nil {
		return nil, 0, ferror.Wrap("获取联系人数量异常", err)
	}
	rs := make([]*model.Contact, 0)
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
		return nil, 0, ferror.Wrap("获取联系人列表异常", err)
	}
	return rs, total, nil

}

// All 获取所有联系人
func (s *ContactSrv) All(ctx context.Context) ([]*model.Contact, error) {

	db := gorms.GetDb(ctx, s.Db)

	rs := make([]*model.Contact, 0)
	err := model.NewContactQuerySet(db).OrderDescByID().All(&rs)
	if err != nil {
		return nil, ferror.Wrap("获取联系人异常", err)
	}
	return rs, nil

}
