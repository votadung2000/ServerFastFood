package bizFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"
)

type DeleteFAQStorage interface {
	FindFAQ(ctx context.Context, cond map[string]interface{}) (*modelFAQ.FAQ, error)
	DeleteFAQ(ctx context.Context, cond map[string]interface{}) error
}

type deleteFAQBiz struct {
	store DeleteFAQStorage
}

func NewDeleteFAQBiz(store DeleteFAQStorage) *deleteFAQBiz {
	return &deleteFAQBiz{store: store}
}

func (biz *deleteFAQBiz) DeleteFAQ(ctx context.Context, id int) error {
	_, err := biz.store.FindFAQ(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(modelFAQ.EntityName, err)
	}

	if err := biz.store.DeleteFAQ(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(modelFAQ.EntityName, err)
	}

	return nil
}
