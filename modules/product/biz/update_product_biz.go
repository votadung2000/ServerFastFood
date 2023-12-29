package bizProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

type UpdateProductStorage interface {
	FindProduct(ctx context.Context, cond map[string]interface{}) (*modelProduct.Product, error)
	UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *modelProduct.ProductUpdate) error
}

type updateProductBiz struct {
	store UpdateProductStorage
}

func UpdateProductBiz(store UpdateProductStorage) *updateProductBiz {
	return &updateProductBiz{store: store}
}

func (biz *updateProductBiz) UpdateProduct(ctx context.Context, id int, dataUpdate *modelProduct.ProductUpdate) error {
	data, err := biz.store.FindProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(modelProduct.EntityName, err)
	}

	if data.Status != 0 && data.Status == modelProduct.STATUS_DELETED {
		return modelProduct.ErrDeleted
	}

	if err := biz.store.UpdateProduct(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(modelProduct.EntityName, err)
	}

	return nil
}
