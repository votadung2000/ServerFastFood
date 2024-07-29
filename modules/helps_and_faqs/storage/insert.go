package storageFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateFAQ(ctx context.Context, data *modelFAQ.FAQCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
