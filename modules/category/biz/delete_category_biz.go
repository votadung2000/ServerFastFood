package bizCategory

import (
	"context"
	modelCategory "fastFood/modules/category/model"
)

type DeleteCategoryStorage interface {
	FindCategory(ctx context.Context, cond map[string]interface{}) (*modelCategory.Category, error)
	DeleteCategory(ctx context.Context, cond map[string]interface{}) error
}

type deleteCategoryBiz struct {
	store DeleteCategoryStorage
}

func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategory(ctx context.Context, id int) error {
	data, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status != 0 && data.Status == modelCategory.STATUS_DELETED {
		return modelCategory.ErrCategoryHasBeenDeleted()
	}

	if err := biz.store.DeleteCategory(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
