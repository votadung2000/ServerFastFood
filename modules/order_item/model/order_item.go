package modelOrderItem

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
	"fmt"
	"strings"
)

const (
	STATUS_ACTION  = 1
	STATUS_DELETED = -1
)

const (
	EntityName = "OrderItem"
)

var (
	ErrOrderIdIsBlank     = "order id cannot be blank"
	ErrProductIdIsBlank   = "product id cannot be blank"
	ErrProductNameIsBlank = "product name cannot be blank"
	ErrQuantityIsBlank    = "quantity cannot be blank"
	ErrPriceIsBlank       = "price cannot be blank"
)

type OrderItem struct {
	common.SQLModel
	OrderId     int                          `json:"-" gorm:"column:order_id;"`
	ProductId   int                          `json:"-" gorm:"column:product_id;"`
	ProductName string                       `json:"-" gorm:"column:product_name;"`
	Status      int                          `json:"status" gorm:"column:status;"`
	Quantity    int                          `json:"quantity" gorm:"column:quantity;"`
	Price       float64                      `json:"price" gorm:"column:price;"`
	Product     *modelProduct.PreloadProduct `json:"product" gorm:"foreignKey:ProductId;"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

type CreateOrderItem struct {
	OrderId     int     `json:"order_id" gorm:"column:order_id;"`
	ProductId   int     `json:"product_id" gorm:"column:product_id;"`
	ProductName string  `json:"product_name" gorm:"column:product_name;"`
	Quantity    int     `json:"quantity" gorm:"column:quantity;"`
	Price       float64 `json:"price" gorm:"column:price;"`
}

func (CreateOrderItem) TableName() string {
	return OrderItem{}.TableName()
}

func (o *CreateOrderItem) SetProductName(name string) {
	o.ProductName = strings.TrimSpace(name)
}

func (o *CreateOrderItem) Validate() error {
	o.ProductName = strings.TrimSpace(o.ProductName)

	if o.OrderId == 0 {
		return ErrValidateRequest(ErrOrderIdIsBlank, "ERR_USER_ID_IS_BLANK")
	}
	if o.ProductId == 0 {
		return ErrValidateRequest(ErrProductIdIsBlank, "ERR_PRODUCT_ID_IS_BLANK")
	}
	if o.ProductName == "" {
		return ErrValidateRequest(ErrProductNameIsBlank, "ERR_PRODUCT_NAME_IS_BLANK")
	}
	if o.Quantity == 0 {
		return ErrValidateRequest(ErrQuantityIsBlank, "ERR_QUANTITY_IS_BLANK")
	}
	if o.Price == 0 {
		return ErrValidateRequest(ErrPriceIsBlank, "ERR_PRICE_IS_BLANK")
	}

	return nil
}

// get interface of DB ->
func (o *OrderItem) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var orderItem OrderItem
	if err := json.Unmarshal(bytes, &orderItem); err != nil {
		return err
	}

	*o = orderItem

	return nil
}

// struct -> DB
func (o *OrderItem) Value() (driver.Value, error) {
	if o == nil {
		return nil, nil
	}

	return json.Marshal(o)
}
