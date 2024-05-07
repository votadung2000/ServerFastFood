package modelVerification

import (
	"fastFood/common"
	"strings"
)

const (
	STATUS_ACTION  = 1
	STATUS_DELETED = -1

	TYPE_EMAIL = 1
)

const (
	EntityName = "Verification"
)

var (
	ErrUserIdIsBlank  = "User id cannot be blank"
	ErrOTPCodeIsBlank = "OTP code cannot be blank"
	ErrTypeIsBlank    = "Type cannot be blank"
	ErrTokenIsBlank   = "Token cannot be blank"
)

type ParamsVerification struct {
	Email string `json:"email"`
}

type Verification struct {
	common.SQLModel
	UserId  int    `json:"user_id" gorm:"column:user_id;"`
	OTPCode string `json:"otp_code" gorm:"column:otp_code;"`
	Type    int    `json:"type" gorm:"column:type;"`
	Status  int    `json:"status" gorm:"column:status;"`
	Token   string `json:"token" gorm:"column:token;"`
}

func (Verification) TableName() string {
	return "verifications"
}

type VerificationCreate struct {
	common.SQLModel
	UserId  int    `json:"user_id" gorm:"column:user_id;"`
	OTPCode string `json:"otp_code" gorm:"column:otp_code;"`
	Type    int    `json:"type" gorm:"column:type;"`
	Token   string `json:"token" gorm:"column:token;"`
}

func (VerificationCreate) TableName() string {
	return Verification{}.TableName()
}

func (i *VerificationCreate) Validate() error {
	i.OTPCode = strings.TrimSpace(i.OTPCode)
	i.Token = strings.TrimSpace(i.Token)

	if i.UserId == 0 {
		return ErrValidateRequest(ErrUserIdIsBlank, "ERR_USER_ID_IS_BLANK")
	}

	if i.OTPCode == "" {
		return ErrValidateRequest(ErrOTPCodeIsBlank, "ERR_OTP_CODE_IS_BLANK")
	}

	if i.Token == "" {
		return ErrValidateRequest(ErrTokenIsBlank, "ERR_TOKEN_IS_BLANK")
	}

	return nil
}
