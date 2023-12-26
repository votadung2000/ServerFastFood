package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func FullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func ErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func ErrorUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func CustomError(root error, msg, key string) *AppError {
	if root != nil {
		return ErrorResponse(root, msg, root.Error(), key)
	}

	return ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return FullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrInternalRequest(err error) *AppError {
	return ErrorResponse(
		err,
		"invalid request",
		err.Error(),
		"ERR_INTERNAL_REQUEST",
	)
}

func ErrInternal(err error) *AppError {
	return FullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong in the server",
		err.Error(),
		"ERR_INTERNAL",
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ERRCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERRCannotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ERRCannotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ERRCannotGet%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ERRCannotCreate%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("ERR%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("ERR%sAlreadyExists", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("ERR%sNotFound", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return CustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

var RecordNoFound = errors.New("record not found")
