package storageUser

import (
	"context"
	"fastFood/common"
	modelUser "fastFood/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateUser(
	ctx context.Context,
	cond map[string]interface{},
	dataUpdate *modelUser.UserUpdate,
) error {
	if err := s.db.Where(cond).
		Updates(dataUpdate).
		First(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
