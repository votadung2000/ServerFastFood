package bizProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

type ListProductStorage interface {
	ListProduct(
		ctx context.Context,
		filter *modelProduct.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]modelProduct.Product, error)
}

type listProductBiz struct {
	store ListProductStorage
}

func NewListProductBiz(store ListProductStorage) *listProductBiz {
	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProduct(
	ctx context.Context,
	filter *modelProduct.Filter,
	paging *common.Paging,
) ([]modelProduct.Product, error) {
	data, err := biz.store.ListProduct(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(modelProduct.EntityName, err)
	}

	return data, nil
}
