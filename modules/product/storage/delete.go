package storageProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) DeleteProduct(
	ctx context.Context,
	cond map[string]interface{},
) error {
	if err := s.db.Table(modelProduct.Product{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": modelProduct.STATUS_DELETED,
		}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
