package bizCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"
)

type ListCategoryStorage interface {
	ListCategory(
		ctx context.Context,
		filter *modelCategory.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]modelCategory.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStorage
}

func NewListCategoryBiz(store ListCategoryStorage) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(
	ctx context.Context,
	filter *modelCategory.Filter,
	paging *common.Paging,
) ([]modelCategory.Category, error) {
	data, err := biz.store.ListCategory(ctx, filter, paging, "Image")
	if err != nil {
		return nil, err
	}
	return data, nil
}
