package storageUpload

import (
	"context"
	"fastFood/common"

	"gorm.io/gorm"
)

func (s *sqlStorage) InsertImage(ctx context.Context, data *common.Image) error {
	if err := s.db.Create(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}
		return common.ErrDB(err)
	}
	return nil
}
