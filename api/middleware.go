package api

import (
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin api for authorization
func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// some middleware

		ctx.Next()
	}
}
