package repoOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"
	modelOrderItem "fastFood/modules/order_item/model"
)

type CreateOrderStorage interface {
	CreateOrder(ctx context.Context, data *modelOrder.CreateOrder) (int, error)
}

type CreateOrderItemStorage interface {
	CreateOrderItem(ctx context.Context, data *[]modelOrderItem.CreateOrderItem) error
}

type createOrderRepo struct {
	storeOrder     CreateOrderStorage
	storeOrderItem CreateOrderItemStorage
}

func NewCreateOrderRepo(
	storeOrder CreateOrderStorage,
	storeOrderItem CreateOrderItemStorage,
) *createOrderRepo {
	return &createOrderRepo{
		storeOrder:     storeOrder,
		storeOrderItem: storeOrderItem,
	}
}

func (repo *createOrderRepo) CreateOrder(ctx context.Context, data *modelOrder.OrderParams) error {
	order := modelOrder.CreateOrder{
		UserId:      data.UserId,
		TaxFees:     data.TaxFees,
		DeliveryFee: data.DeliveryFee,
		Total:       data.Total,
		CouponId:    data.CouponId,
	}

	if err := order.Validate(); err != nil {
		return err
	}

	id, err := repo.storeOrder.CreateOrder(ctx, &order)

	if err != nil {
		return common.ErrCannotCreateEntity(modelOrder.EntityName, err)
	}

	var orderItems = make([]modelOrderItem.CreateOrderItem, len(data.Products))

	for i, v := range data.Products {
		orderItems[i] = modelOrderItem.CreateOrderItem{
			OrderId:     id,
			ProductId:   v.Id,
			ProductName: v.Name,
			Quantity:    v.Quantity,
			Price:       float64(v.Price),
		}

		if err := orderItems[i].Validate(); err != nil {
			return err
		}
	}

	if err := repo.storeOrderItem.CreateOrderItem(ctx, &orderItems); err != nil {
		return common.ErrCannotCreateEntity(modelOrderItem.EntityName, err)
	}

	return nil
}
