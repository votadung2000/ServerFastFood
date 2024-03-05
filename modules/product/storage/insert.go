package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateProduct(ctx context.Context, data *modelProduct.ProductCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
