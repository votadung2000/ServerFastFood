package storageFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) DeleteFavorite(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(modelFavorite.Favorite{}.TableName()).
		Where(cond).
		Delete(nil).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
