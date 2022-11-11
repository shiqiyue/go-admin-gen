package dataloader

import (
	"encoding/json"
	model "github.com/shiqiyue/go-admin-gen/example"

	"gorm.io/datatypes"

	"gorm.io/gorm"

	"time"

	"test/pkg/gqlgens/dataloaders"
)

// ContactLoader 联系人-dataloader

//go:generate go run github.com/shiqiyue/dataloaden ContactPkLoader int64 *github.com/shiqiyue/go-admin-gen/example.Contact

type ContactLoader struct {

	// Db DB 实例
	Db *gorm.DB `inject:""`

	// pkLoader 主键Loader
	pkLoader dataloaders.LoadHelper[ContactPkLoader]
}

// SetUp 初始化Dataloader
func (l *ContactLoader) SetUp() {

	l.GetPkLoader()

}

// GetPkLoader 获取主键Loader
func (l *ContactLoader) GetPkLoader() *ContactPkLoader {

	return l.pkLoader.Get(func() *ContactPkLoader {
		return NewContactPkLoader(ContactPkLoaderConfig{
			Fetch: dataloaders.GenLoadIn(l.Db.Unscoped(), func(data *model.Contact) int64 {
				return data.ID
			}),
			Wait:     10 * time.Millisecond,
			MaxBatch: 100,
		})
	})

}
