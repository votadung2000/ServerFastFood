package storageCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

func (s *sqlStorage) FindCategory(ctx context.Context, cond map[string]interface{}) (*modelCategory.Category, error) {
	var data modelCategory.Category
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
