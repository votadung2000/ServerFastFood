package modelOrder

import (
	"time"
)

const (
	STATUS_WAITING    = 1
	STATUS_PROCESSED  = 2
	STATUS_DELIVERING = 3
	STATUS_DELIVERED  = 4
	STATUS_COMPLETED  = 5
	STATUS_CANCELED   = -1
)

const (
	EntityName = "Order"
)

var (
	ErrUserIsBlank  = "user cannot be blank"
	ErrTotalIsBlank = "total cannot be blank"
)

type Order struct {
	Id          int        `json:"id" gorm:"column:id;"`
	UserId      int        `json:"user_id" gorm:"column:user_id;"`
	Status      int        `json:"status" gorm:"column:status;"`
	TaxFees     float64    `json:"tax_fees" gorm:"column:tax_fees;"`
	DeliveryFee float64    `json:"delivery_fee" gorm:"column:delivery_fee;"`
	Total       float64    `json:"total" gorm:"column:total;"`
	CouponId    int        `json:"coupon_id" gorm:"column:coupon_id;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;"`
	CanceledAt  *time.Time `json:"canceled_at" gorm:"column:canceled_at;"`
	CompletedAt *time.Time `json:"completed_at" gorm:"column:completed_at;"`
	DeliveryAt  *time.Time `json:"delivery_at" gorm:"column:delivery_at;"`
}

func (Order) TableName() string {
	return "orders"
}

type CreateOrder struct {
	UserId      int        `json:"user_id" gorm:"column:user_id;"`
	TaxFees     float64    `json:"tax_fees" gorm:"column:tax_fees;"`
	DeliveryFee float64    `json:"delivery_fee" gorm:"column:delivery_fee;"`
	Total       float64    `json:"total" gorm:"column:total;"`
	CouponId    int        `json:"coupon_id" gorm:"column:coupon_id;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (CreateOrder) TableName() string {
	return Order{}.TableName()
}

func (o *CreateOrder) Validate() error {
	if o.UserId == 0 {
		return ErrValidateRequest(ErrUserIsBlank, "ERR_USER_IS_BLANK")
	}
	if o.Total == 0 {
		return ErrValidateRequest(ErrTotalIsBlank, "ERR_TOTAL_IS_BLANK")
	}

	return nil
}
