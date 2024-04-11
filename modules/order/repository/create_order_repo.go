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
	var order modelOrder.CreateOrder

	order.UserId = data.UserId
	order.TaxFees = data.TaxFees
	order.DeliveryFee = data.DeliveryFee
	order.Total = data.Total
	order.CouponId = data.CouponId

	if err := order.Validate(); err != nil {
		return err
	}

	id, err := repo.storeOrder.CreateOrder(ctx, &order)

	if err != nil {
		return common.ErrCannotCreateEntity(modelOrder.EntityName, err)
	}

	var orderItems = make([]modelOrderItem.CreateOrderItem, len(data.Products))

	for i, v := range data.Products {
		orderItems[i].OrderId = id
		orderItems[i].ProductId = v.Id
		orderItems[i].ProductName = v.Name
		orderItems[i].Quantity = v.Quantity
		orderItems[i].Price = float64(v.Price)

		if err := orderItems[i].Validate(); err != nil {
			return err
		}
	}

	if err := repo.storeOrderItem.CreateOrderItem(ctx, &orderItems); err != nil {
		return common.ErrCannotCreateEntity(modelOrderItem.EntityName, err)
	}

	return nil
}
