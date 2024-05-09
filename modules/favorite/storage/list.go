package storageFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListFavorite(
	ctx context.Context,
	cond map[string]interface{},
	filter *modelFavorite.Filter,
	paging *common.Paging,
) ([]modelFavorite.Favorite, error) {
	var result []modelFavorite.Favorite

	db := s.db.Where(cond)

	db = db.Joins("LEFT JOIN products ON products.id = favorites.product_id")

	if f := filter; f != nil {
		fStatus := f.StatusPr
		if fStatus != 0 {
			db = db.Where("products.status = ?", fStatus)
		} else {
			db = db.Where("products.status = ?", modelProduct.STATUS_ACTION)
		}

		fCategoryId := f.CategoryId
		if fCategoryId != 0 {
			db = db.Where("products.category_id = ?", fCategoryId)
		}
	}

	// db = db.Preload("Product", func(dbPros *gorm.DB) *gorm.DB {

	// 	if f := filter; f != nil {
	// 		fStatus := f.StatusPr
	// 		if fStatus != 0 {
	// 			dbPros = dbPros.Where("status = ?", fStatus)
	// 		} else {
	// 			dbPros = dbPros.Where("status = ?", modelProduct.STATUS_ACTION)
	// 		}

	// 		fCategoryId := f.CategoryId
	// 		if fCategoryId != 0 {
	// 			dbPros = dbPros.Where("category_id = ?", fCategoryId)
	// 		}
	// 	}

	// 	dbPros = dbPros.Preload("Image")

	// 	return dbPros
	// })

	if err := db.
		Select("favorites.id").
		Table(modelFavorite.Favorite{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Select("*").
		Order("favorites.id desc").
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	for i := range result {
		// var product modelProduct.PreloadProduct
		// if err := db.Preload("Product").First(&product, result[i].ProductId).Error; err != nil {
		// 	return nil, common.ErrDB(err)
		// }
		// result[i].Product = &product
		var product modelProduct.PreloadProduct

		if err := db.Preload("Image").Where(cond).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, common.RecordNoFound
			}
			return nil, common.ErrDB(err)
		}

		result[i].Product = &product
	}

	return result, nil
}
