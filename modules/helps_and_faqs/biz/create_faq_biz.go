package bizFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"
)

type CreateFAQStorage interface {
	CreateFAQ(ctx context.Context, data *modelFAQ.FAQCreate) error
}

type createFAQBiz struct {
	store CreateFAQStorage
}

func NewCreateFAQBiz(store CreateFAQStorage) *createFAQBiz {
	return &createFAQBiz{store: store}
}

func (biz *createFAQBiz) CreateFAQ(ctx context.Context, data *modelFAQ.FAQCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateFAQ(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(modelFAQ.EntityName, err)
	}

	return nil
}
