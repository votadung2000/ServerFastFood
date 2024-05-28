package bizOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"
)

type FindOrderStorage interface {
	FindOrderWithPreload(ctx context.Context, cond map[string]interface{}) (*modelOrder.Order, error)
}

type findOrderBiz struct {
	store FindOrderStorage
}

func NewFindOrderBiz(store FindOrderStorage) *findOrderBiz {
	return &findOrderBiz{store: store}
}

func (biz *findOrderBiz) FindOrder(ctx context.Context, orderId, userId int) (*modelOrder.Order, error) {
	data, err := biz.store.FindOrderWithPreload(ctx, map[string]interface{}{"id": orderId, "user_id": userId})
	if err != nil {
		return nil, common.ErrCannotGetEntity(modelOrder.EntityName, err)
	}

	return data, nil
}
