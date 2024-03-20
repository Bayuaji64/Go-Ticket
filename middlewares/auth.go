package middlewares

import (
	"net/http"

	"example.com/rest-api/helper"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token format not valid"})
		return
	}

	token := authHeader[len(bearerPrefix):]

	userId, err := helper.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
