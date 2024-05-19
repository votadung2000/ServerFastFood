package bizOrder

import (
	"context"
	modelOrder "fastFood/modules/order/model"
)

type UpdateOrderStorage interface {
	FindOrder(
		ctx context.Context,
		cond map[string]interface{},
	) (*modelOrder.Order, error)
	UpdateOrder(
		ctx context.Context,
		cond map[string]interface{},
		dataUpdate *modelOrder.UpdateOrder,
	) error
}

type updateOrderBiz struct {
	store UpdateOrderStorage
}

func NewUpdateOrderBiz(store UpdateOrderStorage) *updateOrderBiz {
	return &updateOrderBiz{store: store}
}

func (biz *updateOrderBiz) UpdateOrder(
	ctx context.Context,
	id int,
	data *modelOrder.UpdateOrder,
) error {

	order, err := biz.store.FindOrder(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if order.Status != 0 && data.Status == intToPtr(modelOrder.STATUS_CANCELED) {
		return modelOrder.ErrOrderHasBeenCanceled()
	}

	if err := biz.store.UpdateOrder(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}

func intToPtr(i int) *int {
	return &i
}
