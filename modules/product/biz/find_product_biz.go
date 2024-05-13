package bizProduct

import (
	"context"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

type FindProductWithJoinsStorage interface {
	FindProductWithJoins(ctx context.Context, productId, userId int) (*modelProduct.Product, error)
}

type findProductWithJoinsBiz struct {
	store FindProductWithJoinsStorage
}

func NewFindProductWithJoinsBiz(store FindProductWithJoinsStorage) *findProductWithJoinsBiz {
	return &findProductWithJoinsBiz{store: store}
}

func (biz *findProductWithJoinsBiz) FindProductWithJoins(ctx context.Context, productId, userId int) (*modelProduct.Product, error) {
	data, err := biz.store.FindProductWithJoins(ctx, productId, userId)
	if err != nil {
		return nil, common.ErrCannotGetEntity(modelProduct.EntityName, err)
	}

	return data, nil
}
