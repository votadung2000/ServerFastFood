package storageOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindOrderWithPreload(ctx context.Context, cond map[string]interface{}) (*modelOrder.Order, error) {
	var data modelOrder.Order

	db := s.db.Where(cond)

	db = db.Preload("OrderItems", func(dbOrderItem *gorm.DB) *gorm.DB {
		dbOrderItem = dbOrderItem.Preload("Product", func(dbProduct *gorm.DB) *gorm.DB {
			return dbProduct.Preload("Image")
		})
		return dbOrderItem.Order("id asc")
	})

	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
