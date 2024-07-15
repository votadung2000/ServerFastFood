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
	return s.db.Transaction(func(tx *gorm.DB) error {
		if *dataUpdate.Default == modelDeliveryAddress.DEFAULT {
			if userId, ok := cond["user_id"]; ok {
				if err := tx.Table(modelDeliveryAddress.DeliveryAddress{}.TableName()).
					Where("user_id = ? and `default` = ?", userId, modelDeliveryAddress.DEFAULT).
					Update("`default`", modelDeliveryAddress.NOT_DEFAULT).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						return common.RecordNoFound
					}

					return common.ErrDB(err)
				}
			}
		}

		if id, ok := cond["id"]; ok {
			if err := tx.Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return common.RecordNoFound
				}

				return common.ErrDB(err)
			}
		}

		return nil
	})
}
