package storageOrderItem

import (
	"context"
	"fastFood/common"
	modelOrderItem "fastFood/modules/order_item/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateOrderItem(ctx context.Context, data *[]modelOrderItem.CreateOrderItem) error {
	if err := s.db.Create(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
