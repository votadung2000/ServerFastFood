package bizCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, data *modelCategory.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *modelCategory.CategoryCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}
