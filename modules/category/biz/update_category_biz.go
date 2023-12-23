package bizCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

type UpdateCategoryStorage interface {
	FindCategory(
		ctx context.Context,
		cond map[string]interface{},
	) (*modelCategory.Category, error)
	UpdateCategory(
		ctx context.Context,
		cond map[string]interface{},
		dataUpdate *modelCategory.CategoryUpdate,
	) error
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func UpdateCategoryBiz(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategory(
	ctx context.Context,
	id int,
	dataUpdate *modelCategory.CategoryUpdate,
) error {
	data, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != 0 && data.Status == modelCategory.STATUS_BLOCK {
		return modelCategory.ErrBlocked
	}

	if err := biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}