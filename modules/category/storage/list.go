package storageCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"
	modelProduct "fastFood/modules/product/model"

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

	if err := db.
		Select("id").
		Table(modelCategory.Category{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("Image")
	db = db.Preload("Products", func(dbPros *gorm.DB) *gorm.DB {
		dbPros = dbPros.Preload("Image")
		dbPros = dbPros.Where("status = ?", modelProduct.STATUS_ACTION)
		dbPros = dbPros.Where("featured = ?", modelProduct.FEATURED_OUTSTANDING)
		return dbPros.Order("id asc")
	})

	if err := db.Select("*").
		// Order("id desc").
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
