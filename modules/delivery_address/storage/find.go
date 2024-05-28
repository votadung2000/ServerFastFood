package storageDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindDeliveryAddress(ctx context.Context, cond map[string]interface{}) (*modelDeliveryAddress.DeliveryAddress, error) {
	var data modelDeliveryAddress.DeliveryAddress
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
