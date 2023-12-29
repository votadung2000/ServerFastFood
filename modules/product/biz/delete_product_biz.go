package bizProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

type DeleteProductStorage interface {
	FindProduct(ctx context.Context, cond map[string]interface{}) (*modelProduct.Product, error)
	DeleteProduct(ctx context.Context, cond map[string]interface{}) error
}

type deleteProductBiz struct {
	store DeleteProductStorage
}

func DeleteProductBiz(store DeleteProductStorage) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

func (biz *deleteProductBiz) DeleteProduct(ctx context.Context, id int) error {
	data, err := biz.store.FindProduct(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(modelProduct.EntityName, err)
	}

	if data.Status != 0 && data.Status == modelProduct.STATUS_DELETED {
		return modelProduct.ErrDeleted
	}

	if err := biz.store.DeleteProduct(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(modelProduct.EntityName, err)
	}

	return nil
}
