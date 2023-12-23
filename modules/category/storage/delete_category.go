package storageCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

func (s *sqlStorage) DeleteCategory(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(modelCategory.Category{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": modelCategory.STATUS_DELETED,
		}).Error; err != nil {
		return err
	}

	return nil
}
