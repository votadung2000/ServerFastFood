package storageOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateOrder(ctx context.Context, data *modelOrder.CreateOrder) (int, error) {
	if err := s.db.Create(data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, common.RecordNoFound
		}

		return 0, common.ErrDB(err)
	}

	return data.Id, nil
}
