package storageDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListDeliveryAddress(
	ctx context.Context,
	filter *modelDeliveryAddress.Filter,
	paging *common.Paging,
	cond map[string]interface{},
	moreKeys ...string,
) ([]modelDeliveryAddress.DeliveryAddress, error) {
	var result []modelDeliveryAddress.DeliveryAddress

	db := s.db.Where(cond)

	if f := filter; f != nil {
		fStatus := f.Status

		if fStatus != 0 {
			db = db.Where("status IN ?", fStatus)
		}
	}

	if err := db.
		Select("id").
		Table(modelDeliveryAddress.DeliveryAddress{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Select("*").
		Order("id desc").
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return result, nil
}
