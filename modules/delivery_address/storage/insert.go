package storageDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateDeliveryAddress(ctx context.Context, id int, data *modelDeliveryAddress.CreateDeliveryAddress) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if data.Default == modelDeliveryAddress.DEFAULT {
			if err := tx.Table(modelDeliveryAddress.DeliveryAddress{}.TableName()).
				Where("user_id = ? and `default` = ?", id, modelDeliveryAddress.DEFAULT).
				Update("`default`", modelDeliveryAddress.NOT_DEFAULT).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return common.RecordNoFound
				}

				return common.ErrDB(err)
			}
		}

		if err := tx.Create(&data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return common.RecordNoFound
			}

			return common.ErrDB(err)
		}

		return nil
	})
}
