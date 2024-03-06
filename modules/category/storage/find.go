package storageCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindCategory(ctx context.Context, cond map[string]interface{}) (*modelCategory.Category, error) {
	var data modelCategory.Category
	if err := s.db.Preload("Image").Where(cond).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
