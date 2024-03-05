package storageFavorite

import (
	"context"
	modelFavorite "fastFood/modules/favorite/model"
)

func (s *sqlStorage) DeleteFavorite(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(modelFavorite.Favorite{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": modelFavorite.STATUS_DELETED,
		}).Error; err != nil {
		return err
	}

	return nil
}
