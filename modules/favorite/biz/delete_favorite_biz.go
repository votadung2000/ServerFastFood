package bizFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"
)

type DeleteFavoriteStorage interface {
	FindFavorite(ctx context.Context, cond map[string]interface{}) (*modelFavorite.Favorite, error)
	DeleteFavorite(ctx context.Context, cond map[string]interface{}) error
}

type deleteFavoriteBiz struct {
	store DeleteFavoriteStorage
}

func DeleteFavoriteBiz(store DeleteFavoriteStorage) *deleteFavoriteBiz {
	return &deleteFavoriteBiz{store: store}
}

func (biz *deleteFavoriteBiz) DeleteFavorite(ctx context.Context, id int) error {
	data, err := biz.store.FindFavorite(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(modelFavorite.EntityName, err)
	}

	if data.Status != 0 && data.Status == modelFavorite.STATUS_DELETED {
		return modelFavorite.ErrDeleted
	}

	if err := biz.store.DeleteFavorite(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
