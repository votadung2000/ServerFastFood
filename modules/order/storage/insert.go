package storageOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateOrder(ctx context.Context, data *modelOrder.CreateOrder) error {
	if err := s.db.Create(data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
