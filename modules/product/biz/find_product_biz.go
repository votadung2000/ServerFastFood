package bizProduct

import (
	"context"
	modelProduct "fastFood/modules/product/model"
)

type FindProductStorage interface {
	FindProduct(ctx context.Context, cond map[string]interface{}) (*modelProduct.Product, error)
}

type findProductBiz struct {
	store FindProductStorage
}

func FindProductBiz(store FindProductStorage) *findProductBiz {
	return &findProductBiz{store: store}
}

func (biz *findProductBiz) FindProduct(ctx context.Context, id int) (*modelProduct.Product, error) {
	data, err := biz.store.FindProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}
