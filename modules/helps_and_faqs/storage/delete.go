package storageFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) DeleteFAQ(
	ctx context.Context,
	cond map[string]interface{},
) error {
	if err := s.db.Table(modelFAQ.FAQ{}.TableName()).
		Where(cond).
		Delete(nil).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNoFound
		}

		return common.ErrDB(err)
	}

	return nil
}
