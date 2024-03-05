package storageFavorite

import (
	"context"
	modelFavorite "fastFood/modules/favorite/model"
)

func (s *sqlStorage) CreateFavorite(ctx context.Context, data *modelFavorite.FavoriteCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
