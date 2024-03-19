package bizFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"
)

type ListFavoriteStorage interface {
	ListFavorite(
		ctx context.Context,
		cond map[string]interface{},
		filter *modelFavorite.Filter,
		paging *common.Paging,
	) ([]modelFavorite.Favorite, error)
}

type listFavoriteBiz struct {
	store ListFavoriteStorage
}

func NewListFavoriteBiz(store ListFavoriteStorage) *listFavoriteBiz {
	return &listFavoriteBiz{store: store}
}

func (biz *listFavoriteBiz) ListFavorite(
	ctx context.Context,
	id int,
	filter *modelFavorite.Filter,
	paging *common.Paging,
) ([]modelFavorite.Favorite, error) {
	data, err := biz.store.ListFavorite(
		ctx,
		map[string]interface{}{"user_id": id},
		filter,
		paging,
	)

	if err != nil {
		return nil, err
	}

	return data, nil
}
