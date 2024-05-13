package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindProductWithJoins(ctx context.Context, productId, userId int) (*modelProduct.Product, error) {
	var data modelProduct.Product

	if err := s.db.
		Table(modelProduct.Product{}.TableName()).
		Select("products.*, IF(favorites.id IS NOT NULL, TRUE, FALSE) AS is_favorite").
		Joins("LEFT JOIN favorites ON products.id = favorites.product_id AND favorites.user_id = ?", userId).
		Preload("Image").
		Where("products.id = ?", productId).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
