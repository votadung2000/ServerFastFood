package storageFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateFavorite(ctx context.Context, data *modelFavorite.FavoriteCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
