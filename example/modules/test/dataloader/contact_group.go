package dataloader

import (
	"encoding/json"

	"gorm.io/datatypes"

	"gorm.io/gorm"

	"time"

	"test/pkg/gqlgens/dataloaders"
)

// ContactGroupLoader 联系人分组-dataloader
//go:generate go run github.com/shiqiyue/dataloaden ContactGroupPkLoader int64 *github.com/shiqiyue/go-admin-gen/example.ContactGroup

type ContactGroupLoader struct {
	// Db DB 实例
	Db *gorm.DB `inject:""`
	// pkLoader 主键Loader
	pkLoader dataloaders.LoadHelper[ContactGroupPkLoader]
}

// SetUp 初始化Dataloader
func (l *ContactGroupLoader) SetUp() {

	l.GetPkLoader()

}

// GetPkLoader 获取主键Loader
func (l *ContactGroupLoader) GetPkLoader() *ContactGroupPkLoader {

	return l.pkLoader.Get(func() *ContactGroupPkLoader {
		return NewContactGroupPkLoader(ContactGroupPkLoaderConfig{
			Fetch: dataloaders.GenLoadIn(l.Db.Unscoped(), func(data *model.ContactGroup) int64 {
				return data.ID
			}),
			Wait:     10 * time.Millisecond,
			MaxBatch: 100,
		})
	})

}
