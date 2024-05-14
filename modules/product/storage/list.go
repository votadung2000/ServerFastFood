package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListProduct(
	ctx context.Context,
	userId int,
	filter *modelProduct.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]modelProduct.Product, error) {
	var result []modelProduct.Product

	db := s.db

	if err := db.
		Select("id").
		Table(modelProduct.Product{}.TableName()).
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

		fName := f.Name
		if fName != "" {
			db = db.Where("name LIKE ?", "%"+fName+"%")
		}

		fCategoryId := f.CategoryId
		if fCategoryId != 0 {
			db = db.Where("products.category_id = ?", fCategoryId)
		}
	}

	if err := db.
		Select("products.*, IF(favorites.id IS NOT NULL, TRUE, FALSE) AS is_favorite").
		Joins("LEFT JOIN favorites ON products.id = favorites.product_id AND favorites.user_id = ?", userId).
		Preload("Image").
		Order("products.id desc").
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}
		return nil, common.ErrDB(err)
	}

	return result, nil
}
