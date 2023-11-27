package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		clientToken := context.Request.Header.Get("Authorization")

		if clientToken == "" {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "No Authorization header provided"})
		}

		fmt.Println("clientToken", clientToken)
	}
}
