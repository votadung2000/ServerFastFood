package storageProduct

import (
	"context"
	modelProduct "fastFood/modules/product/model"
)

func (s *sqlStorage) FindProduct(ctx context.Context, cond map[string]interface{}) (*modelProduct.Product, error) {
	var data modelProduct.Product

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
