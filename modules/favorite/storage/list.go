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

	if err := db.
		Select("id").
		Table(modelFavorite.Favorite{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("Product", func(dbPros *gorm.DB) *gorm.DB {

		if f := filter; f != nil {
			fStatus := f.StatusPr
			if fStatus != 0 {
				dbPros = dbPros.Where("status = ?", fStatus)
			} else {
				dbPros = dbPros.Where("status = ?", modelProduct.STATUS_ACTION)
			}

			fCategoryId := f.CategoryId
			if fCategoryId != 0 {
				dbPros = dbPros.Where("category_id = ?", fCategoryId)
			}
		}

		dbPros = dbPros.Preload("Image")

		return dbPros
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

	filteredResult := make([]modelFavorite.Favorite, 0, len(result))
	for _, favorite := range result {
		if favorite.Product != nil {
			filteredResult = append(filteredResult, favorite)
		}
	}

	if len(filteredResult) == 0 {
		return nil, nil
	}

	return filteredResult, nil
}
