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
	if data.UserId == 0 {
		return modelFavorite.ErrUserIsBlank
	}

	if data.ProductId == 0 {
		return modelFavorite.ErrProductIsBlank
	}

	if err := biz.store.CreateFavorite(ctx, data); err != nil {
		return err
	}

	return nil
}
