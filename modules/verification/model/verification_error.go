package modelVerification

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrCannotCreateEntity(err error) *common.AppError {
	return common.CustomError(
		err,
		"Cannot create verification",
		"ERR_CANNOT_CREATE_VERIFICATION",
	)
}
