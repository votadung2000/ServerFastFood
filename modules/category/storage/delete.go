package storageCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) DeleteCategory(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(modelCategory.Category{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": modelCategory.STATUS_DELETED,
		}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
