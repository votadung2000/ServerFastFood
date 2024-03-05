package bizProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

type FindProductStorage interface {
	FindProduct(ctx context.Context, cond map[string]interface{}) (*modelProduct.Product, error)
}

type findProductBiz struct {
	store FindProductStorage
}

func NewFindProductBiz(store FindProductStorage) *findProductBiz {
	return &findProductBiz{store: store}
}

func (biz *findProductBiz) FindProduct(ctx context.Context, id int) (*modelProduct.Product, error) {
	data, err := biz.store.FindProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(modelProduct.EntityName, err)
	}

	return data, nil
}
