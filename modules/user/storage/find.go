package storageUser

import (
	"context"
	"fastFood/common"
	modelUser "fastFood/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error) {
	var data modelUser.User

	if err := s.db.Preload("Image").Where(cond).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
