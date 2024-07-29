package storageFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListFAQ(
	ctx context.Context,
	filter *modelFAQ.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]modelFAQ.FAQ, error) {
	var result []modelFAQ.FAQ

	db := s.db

	if err := db.
		Select("id").
		Table(modelFAQ.FAQ{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if f := filter; f != nil {
		fStatus := f.Status
		if fStatus != 0 {
			db = db.Where("products.status = ?", fStatus)
		} else {
			db = db.Where("products.status = ?", modelProduct.STATUS_ACTION)
		}
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
