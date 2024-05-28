package storageDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) DeleteDeliveryAddress(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(modelDeliveryAddress.DeliveryAddress{}.TableName()).
		Where(cond).
		Delete(nil).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
