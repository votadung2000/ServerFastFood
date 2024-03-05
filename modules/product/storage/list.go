package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListProduct(
	ctx context.Context,
	filter *modelProduct.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]modelProduct.Product, error) {
	var result []modelProduct.Product

	db := s.db

	if f := filter; f != nil {
		fStatus := f.Status
		if fStatus != 0 {
			db = db.Where("status = ?", fStatus)
		} else {
			db = db.Where("status = ?", modelProduct.STATUS_ACTION)
		}

		fName := f.Name
		if fName != "" {
			db.Where("name LIKE ?", "%"+fName+"%")
		}

		fCategoryId := f.CategoryId
		if fCategoryId != 0 {
			db = db.Where("category_id = ?", fCategoryId)
		}
	}

	if err := db.Select("id").Table(modelProduct.Product{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Select("*").Order("id desc").
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return result, nil
}