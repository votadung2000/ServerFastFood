package storageFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) FindFAQ(ctx context.Context, cond map[string]interface{}) (*modelFAQ.FAQ, error) {
	var data modelFAQ.FAQ

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
