package http

import (
	"auth-service/auth"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers) {
	authGroup.POST("/login", h.Login)
	authGroup.POST("/register", h.Register)
}
