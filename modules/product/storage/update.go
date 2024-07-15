package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateProduct(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelProduct.ProductUpdate,
) error {
	if err := s.db.Where(cond).
		Updates(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStorage) IncreaseProductSold(ctx context.Context, id, quantity int) error {
	if err := s.db.Table(modelProduct.Product{}.TableName()).
		Where("id = ?", id).
		Update("sold", gorm.Expr("sold + ?", quantity)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
