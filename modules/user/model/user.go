package modelUser

import (
	"fastFood/common"
	"strings"
)

const (
	STATUS_ACTION  = 1
	STATUS_BLOCK   = -1
	STATUS_DELETED = -2

	ROLE_USER  = 1
	ROLE_STAFF = 2
	ROLE_ADMIN = 3
)

const (
	EntityName = "Product"
)

var (
	ErrDeleted                  = "the user has been deleted"
	ErrUserNameIsBlank          = "User name cannot be blank"
	ErrPasswordIsBlank          = "Password cannot be blank"
	ErrNameIsBlank              = "Last name cannot be blank"
	ErrPhoneNumberIsBlank       = "Phone number cannot be blank"
	ErrInvalidPhoneNumberFormat = "Invalid phone number format"
	ErrEmailIsBlank             = "Email cannot be blank"
	ErrInvalidEmailFormat       = "Invalid email format"
)

type User struct {
	common.SQLModel
	Name        string `json:"name" gorm:"column:name;"`
	UserName    string `json:"user_name" gorm:"column:user_name;"`
	Password    string `json:"-" gorm:"column:password;"`
	Salt        string `json:"-" gorm:"column:salt"`
	Status      int    `json:"status" gorm:"column:status;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Email       string `json:"email" gorm:"column:email"`
	Address     string `json:"address" gorm:"column:address"`
	Role        int    `json:"role" gorm:"column:role"`
	AvatarId    int    `json:"-" gorm:"column:avatar_id"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

type UserCreate struct {
	Name        string `json:"name" gorm:"column:name;"`
	UserName    string `json:"user_name" gorm:"column:user_name;"`
	Password    string `json:"password" gorm:"column:password;"`
	Salt        string `json:"-" gorm:"column:salt"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Role        int    `json:"role" gorm:"column:role"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (i *UserCreate) Validate() error {
	i.UserName = strings.TrimSpace(i.UserName)
	i.Password = strings.TrimSpace(i.Password)
	i.Name = strings.TrimSpace(i.Name)
	i.PhoneNumber = strings.TrimSpace(i.PhoneNumber)

	if i.UserName == "" {
		return ErrValidateRequest(ErrUserNameIsBlank, "ERR_USER_NAME_IS_BLANK")
	}

	if i.Password == "" {
		return ErrValidateRequest(ErrPasswordIsBlank, "ERR_PASSWORD_IS_BLANK")
	}

	if i.Name == "" {
		return ErrValidateRequest(ErrNameIsBlank, "ERR_AME_IS_BLANK")
	}

	if i.PhoneNumber == "" {
		return ErrValidateRequest(ErrPhoneNumberIsBlank, "ERR_PHONE_NUMBER_IS_BLANK")
	}

	if len(i.PhoneNumber) != 10 {
		return ErrValidateRequest(ErrInvalidPhoneNumberFormat, "ERR_INVALID_PHONE_NUMBER_FORMAT")
	}

	return nil
}

type UserUpdate struct {
	Name        *string `json:"name,omitempty" gorm:"column:name;"`
	PhoneNumber *string `json:"phone_number,omitempty" gorm:"column:phone_number"`
	Email       *string `json:"email,omitempty" gorm:"column:email"`
	Address     *string `json:"address,omitempty" gorm:"column:address"`
	AvatarId    *int    `json:"avatar_id,omitempty" gorm:"column:avatar_id"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (i *UserUpdate) Validate() error {
	if i.Name != nil && strings.TrimSpace(*i.Name) == "" {
		return ErrValidateRequest(ErrNameIsBlank, "ERR_AME_IS_BLANK")
	}

	if i.PhoneNumber != nil && strings.TrimSpace(*i.PhoneNumber) == "" {
		return ErrValidateRequest(ErrPhoneNumberIsBlank, "ERR_PHONE_NUMBER_IS_BLANK")
	}

	if i.PhoneNumber != nil && len(strings.TrimSpace(*i.PhoneNumber)) != 10 {
		return ErrValidateRequest(ErrInvalidPhoneNumberFormat, "ERR_INVALID_PHONE_NUMBER_FORMAT")
	}

	if i.Email != nil && strings.TrimSpace(*i.Email) == "" {
		return ErrValidateRequest(ErrEmailIsBlank, "ERR_EMAIL_IS_BLANK")
	}

	partsEmail := strings.Split(strings.TrimSpace(*i.Email), "@")
	if len(partsEmail) != 2 {
		return ErrValidateRequest(ErrInvalidEmailFormat, "ERR_INVALID_EMAIL_FORMAT")
	}

	return nil
}

type Login struct {
	UserName string `json:"user_name" form:"user_name" gorm:"column:user_name"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (Login) TableName() string {
	return User{}.TableName()
}
