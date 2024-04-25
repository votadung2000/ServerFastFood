package middleware

import (
	"context"
	"errors"
	"fastFood/common"
	"fastFood/components/tokenProvider"
	modelUser "fastFood/modules/user/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrDeletedOrBanned = errors.New("user has been deleted or banned")
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.CustomError(
		err,
		"Wrong authentication header",
		"ERR_WRONG_AUTH_HEADER",
	)
}

type AuthStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(authStorage AuthStorage, tokenProvider tokenProvider.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := authStorage.FindUser(ctx.Request.Context(), map[string]interface{}{"id": payload.UserId()})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if user.Status == 0 {
			ctx.JSON(http.StatusBadRequest, common.ErrNoPermission(ErrDeletedOrBanned))
			return
		}

		ctx.Set(common.CurrentUser, user)
		ctx.Next()
	}
}
