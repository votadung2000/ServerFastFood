package storageUpload

import (
	"context"
	"fastFood/common"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindImage(ctx context.Context, cond map[string]interface{}) (*common.Image, error) {
	var data common.Image

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
