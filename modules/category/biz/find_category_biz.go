package bizCategory

import (
	"context"

	"fastFood/common"
	modelCategory "fastFood/modules/category/model"
)

type FindCategoryStorage interface {
	FindCategory(ctx context.Context, cond map[string]interface{}) (*modelCategory.Category, error)
}

type findCategoryBiz struct {
	store FindCategoryStorage
}

func NewFindCategoryBiz(store FindCategoryStorage) *findCategoryBiz {
	return &findCategoryBiz{store: store}
}

func (biz *findCategoryBiz) FindCategory(ctx context.Context, id int) (*modelCategory.Category, error) {
	data, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(modelCategory.EntityName, err)
	}

	return data, nil
}
