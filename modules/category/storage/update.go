package storageCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateCategory(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelCategory.CategoryUpdate,
) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
