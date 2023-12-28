package bizProduct

import (
	"context"
	modelProduct "fastFood/modules/product/model"
)

type CreateProductStorage interface {
	CreateProduct(ctx context.Context, data *modelProduct.ProductCreate) error
}

type createProductBiz struct {
	store CreateProductStorage
}

func CreateProductBiz(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(ctx context.Context, data *modelProduct.ProductCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateProduct(ctx, data); err != nil {
		return err
	}

	return nil
}
