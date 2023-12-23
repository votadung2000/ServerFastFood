package storageCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

func (s *sqlStorage) CreateCategory(ctx context.Context, data *modelCategory.CategoryCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
