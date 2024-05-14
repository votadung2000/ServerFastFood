package bizFavorite

import (
	"context"
	modelFavorite "fastFood/modules/favorite/model"
)

type CDFavoriteStorage interface {
	CreateFavorite(ctx context.Context, data *modelFavorite.FavoriteCreate) error
	FindFavorite(ctx context.Context, cond map[string]interface{}) (*modelFavorite.Favorite, error)
	DeleteFavorite(ctx context.Context, cond map[string]interface{}) error
}

type cdFavoriteBiz struct {
	store CDFavoriteStorage
}

func CDFavoriteBiz(store CDFavoriteStorage) *cdFavoriteBiz {
	return &cdFavoriteBiz{store: store}
}

func (biz *cdFavoriteBiz) CDFavorite(ctx context.Context, data *modelFavorite.FavoriteCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	favorite, err := biz.store.FindFavorite(ctx, map[string]interface{}{"product_id": data.ProductId, "user_id": data.UserId})

	if err != nil {
		return err
	}

	if favorite == nil {
		if err := biz.store.CreateFavorite(ctx, data); err != nil {
			return err
		}
	}

	if favorite != nil {
		if err := biz.store.DeleteFavorite(ctx, map[string]interface{}{"product_id": data.ProductId, "user_id": data.UserId}); err != nil {
			return err
		}
	}

	return nil
}
