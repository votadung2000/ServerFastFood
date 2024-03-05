package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

func (s *sqlStorage) CreateProduct(ctx context.Context, data *modelProduct.ProductCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
