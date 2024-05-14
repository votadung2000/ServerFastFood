package storageUser

import (
	"context"
	"fastFood/common"
	modelUser "fastFood/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) InsertUser(ctx context.Context, data *modelUser.UserCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
