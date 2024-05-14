package storageVerification

import (
	"context"
	"fastFood/common"
	modelVerification "fastFood/modules/verification/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindVerification(ctx context.Context, cond map[string]interface{}) (*modelVerification.Verification, error) {
	var data modelVerification.Verification

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
