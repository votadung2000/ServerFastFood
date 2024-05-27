package modelDeliveryAddress

import "fastFood/common"

const (
	STATUS_ACTION  = 1
	STATUS_DELETED = -2

	TYPE_HOME   = 1
	TYPE_OFFICE = 2
	TYPE_OTHER  = 3
)

var (
	ErrUserIdIsBlank        = "the user id cannot be blank"
	ErrRecipientNameIsBlank = "the recipient name cannot be blank"
	ErrPhoneNumberIsBlank   = "the phone number cannot be blank"
	ErrStreetAddressIsBlank = "the street address cannot be blank"
	ErrCountryIsBlank       = "the country cannot be blank"
	ErrCityIsBlank          = "the city cannot be blank"
	ErrPostalCodeIsBlank    = "the postal code cannot be blank"
)

type DeliveryAddress struct {
	common.SQLModel
	UserId        int    `json:"user_id" gorm:"column:user_id;"`
	Status        int    `json:"status" gorm:"column:status;"`
	Type          int    `json:"type" gorm:"column:type;"`
	RecipientName string `json:"recipient_name" gorm:"column:recipient_name;"`
	PhoneNumber   string `json:"phone_number" gorm:"column:phone_number"`
	StreetAddress string `json:"street_address" gorm:"column:street_address;"`
	Country       string `json:"country" gorm:"column:country;"`
	City          string `json:"city" gorm:"column:city;"`
	PostalCode    string `json:"postal_code" gorm:"column:postal_code;"`
	Description   string `json:"description" gorm:"column:description;"`
}

func (DeliveryAddress) TableName() string {
	return "delivery_address"
}

type CreateDeliveryAddress struct {
	UserId        int    `json:"user_id" gorm:"column:user_id;"`
	RecipientName string `json:"recipient_name" gorm:"column:recipient_name;"`
	PhoneNumber   string `json:"phone_number" gorm:"column:phone_number"`
	StreetAddress string `json:"street_address" gorm:"column:street_address;"`
	Type          int    `json:"type" gorm:"column:type;"`
	Country       string `json:"country" gorm:"column:country;"`
	City          string `json:"city" gorm:"column:city;"`
	PostalCode    string `json:"postal_code" gorm:"column:postal_code;"`
	Description   string `json:"description" gorm:"column:description;"`
}

func (CreateDeliveryAddress) TableName() string {
	return DeliveryAddress{}.TableName()
}

func (c *CreateDeliveryAddress) Validate() error {
	if c.UserId == 0 {
		return ErrValidateRequest(ErrUserIdIsBlank, "ERR_USER_ID_IS_BLANK")
	}

	if c.RecipientName == "" {
		return ErrValidateRequest(ErrRecipientNameIsBlank, "ERR_RECIPIENT_NAME_IS_BLANK")
	}

	if c.PhoneNumber == "" {
		return ErrValidateRequest(ErrPhoneNumberIsBlank, "ERR_PHONE_NUMBER_IS_BLANK")
	}

	if c.StreetAddress == "" {
		return ErrValidateRequest(ErrStreetAddressIsBlank, "ERR_STREET_ADDRESS_IS_BLANK")
	}

	if c.Country == "" {
		return ErrValidateRequest(ErrCountryIsBlank, "ERR_COUNTRY_IS_BLANK")
	}

	if c.City == "" {
		return ErrValidateRequest(ErrCityIsBlank, "ERR_CITY_IS_BLANK")
	}

	if c.PostalCode == "" {
		return ErrValidateRequest(ErrPostalCodeIsBlank, "ERR_POSTAL_CODE_IS_BLANK")
	}

	return nil
}
