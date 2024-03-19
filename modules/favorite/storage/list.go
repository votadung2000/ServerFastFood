package storageFavorite

import (
	"context"
	"fastFood/common"
	modelFavorite "fastFood/modules/favorite/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) ListFavorite(
	ctx context.Context,
	cond map[string]interface{},
	filter *modelFavorite.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]modelFavorite.Favorite, error) {
	var result []modelFavorite.Favorite

	db := s.db.Where(cond)

	// if f := filter; f != nil {

	// }

	if err := db.
		Select("id").
		Table(modelFavorite.Favorite{}.TableName()).
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
