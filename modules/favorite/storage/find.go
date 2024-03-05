package storageFavorite

import (
	"context"
	modelFavorite "fastFood/modules/favorite/model"
)

func (s *sqlStorage) FindFavorite(ctx context.Context, cond map[string]interface{}) (*modelFavorite.Favorite, error) {
	var data modelFavorite.Favorite

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
