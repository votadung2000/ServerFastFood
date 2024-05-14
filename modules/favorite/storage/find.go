package storageFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindFavorite(ctx context.Context, cond map[string]interface{}) (*modelFavorite.Favorite, error) {
	var data modelFavorite.Favorite

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
