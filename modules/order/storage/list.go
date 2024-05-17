package storageOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListOrder(
	ctx context.Context,
	filter *modelOrder.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]modelOrder.Order, error) {
	var result []modelOrder.Order

	db := s.db

	if f := filter; f != nil {
		fUpcoming := f.IsUpcoming
		fHistory := f.IsHistory
		fStatus := f.Status

		if fUpcoming && !(fStatus != 0) && !fHistory {
			statuses := []int{
				modelOrder.STATUS_WAITING,
				modelOrder.STATUS_PROCESSED,
				modelOrder.STATUS_DELIVERING,
				modelOrder.STATUS_DELIVERED,
			}
			db = db.Where("status IN ?", statuses)
		}

		if fHistory && !(fStatus != 0) && !fUpcoming {
			statuses := []int{modelOrder.STATUS_COMPLETED, modelOrder.STATUS_CANCELED}
			db = db.Where("status IN ?", statuses)
		}

		if fStatus != 0 && !fUpcoming && !fHistory {
			db = db.Where("status IN ?", fStatus)
		}
	}

	if err := db.
		Select("id").
		Table(modelOrder.Order{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("OrderItems", func(dbOrderItem *gorm.DB) *gorm.DB {
		dbOrderItem = dbOrderItem.Preload("Product", func(dbProduct *gorm.DB) *gorm.DB {
			return dbProduct.Preload("Image")
		})
		return dbOrderItem.Order("id asc")
	})

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
