package storageDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateDeliveryAddress(ctx context.Context, data *modelDeliveryAddress.CreateDeliveryAddress) error {
	if err := s.db.Create(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
