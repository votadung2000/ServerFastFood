package middleware

import (
	"context"
	"fastFood/components/tokenProvider"
	modelUser "fastFood/modules/user/model"

	"github.com/gin-gonic/gin"
)

type AuthStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}

func RequireAuth(authStorage AuthStorage, tokenProvider tokenProvider.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
