package storageCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListCategory(
	ctx context.Context,
	filter *modelCategory.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]modelCategory.Category, error) {
	var result []modelCategory.Category

	db := s.db

	if f := filter; f != nil {
		fStatus := f.Status
		if fStatus != 0 {
			db = db.Where("status = ?", fStatus)
		} else {
			db = db.Where("status = ?", modelCategory.STATUS_ACTION)
		}
	}

	if err := db.Table(modelCategory.Category{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Order("id desc").
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
