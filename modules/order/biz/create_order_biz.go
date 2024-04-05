package bizOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"
)

type CreateOrderStorage interface {
	CreateOrder(ctx context.Context, data *modelOrder.CreateOrder) error
}

type createOrderBiz struct {
	store CreateOrderStorage
}

func NewCreateOrder(store CreateOrderStorage) *createOrderBiz {
	return &createOrderBiz{store: store}
}

func (biz *createOrderBiz) CreateOrder(ctx context.Context, data *modelOrder.CreateOrder) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateOrder(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(modelOrder.EntityName, err)
	}

	return nil
}
