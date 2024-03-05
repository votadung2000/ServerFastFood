package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

func (s *sqlStorage) UpdateProduct(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelProduct.ProductUpdate,
) error {
	if err := s.db.Where(cond).
		Updates(dataUpdate).
		First(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
