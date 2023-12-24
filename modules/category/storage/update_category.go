package storageCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

func (s *sqlStorage) UpdateCategory(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelCategory.CategoryUpdate,
) error {
	if err := s.db.Where(cond).Updates(dataUpdate).First(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
