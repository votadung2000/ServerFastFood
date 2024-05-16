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
		fStatus := f.Status
		if fStatus != 0 {
			db = db.Where("status = ?", fStatus)
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
