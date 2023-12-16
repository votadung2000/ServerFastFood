package bizCategory

import (
	"context"

	modelCategory "fastFood/modules/category/model"
)

type FindCategoryStorage interface {
	FindCategory(ctx context.Context, cond map[string]interface{}) (*modelCategory.Category, error)
}

type findCategoryBiz struct {
	store FindCategoryStorage
}

func FindCategoryBiz(store FindCategoryStorage) *findCategoryBiz {
	return &findCategoryBiz{store: store}
}

func (biz *findCategoryBiz) FindCategory(ctx context.Context, id int) (*modelCategory.Category, error) {
	data, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}
