package storageDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateCategory(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelDeliveryAddress.DeliveryAddressUpdate,
) error {
	if err := s.db.Where(cond).Updates(dataUpdate).First(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
