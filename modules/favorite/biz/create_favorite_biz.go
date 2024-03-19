package bizFavorite

import (
	"context"
	modelFavorite "fastFood/modules/favorite/model"
)

type CreateFavoriteStorage interface {
	CreateFavorite(ctx context.Context, data *modelFavorite.FavoriteCreate) error
}

type createFavoriteBiz struct {
	store CreateFavoriteStorage
}

func CreateFavoriteBiz(store CreateFavoriteStorage) *createFavoriteBiz {
	return &createFavoriteBiz{store: store}
}

func (biz *createFavoriteBiz) CreateFavorite(ctx context.Context, data *modelFavorite.FavoriteCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateFavorite(ctx, data); err != nil {
		return err
	}

	return nil
}
