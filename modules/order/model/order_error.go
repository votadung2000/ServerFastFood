package modelOrder

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}
