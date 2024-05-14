package storageCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateCategory(ctx context.Context, data *modelCategory.CategoryCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
