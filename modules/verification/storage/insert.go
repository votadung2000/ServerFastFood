package storageVerification

import (
	"context"
	"fastFood/common"
	modelVerification "fastFood/modules/verification/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) InsertVerification(ctx context.Context, data *modelVerification.VerificationCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
