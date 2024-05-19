package storageOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateOrder(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelOrder.UpdateOrder,
) error {
	if err := s.db.Where(cond).Updates(dataUpdate).First(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
