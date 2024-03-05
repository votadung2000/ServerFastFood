package bizCategory

import (
	"context"
	"fastFood/common"
	modelCategory "fastFood/modules/category/model"
)

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, data *modelCategory.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func CreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *modelCategory.CategoryCreate) error {
	data.Name = common.Sanitize(data.Name)

	if data.Name == "" {
		return modelCategory.ErrNameIsBlank
	}

	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}
