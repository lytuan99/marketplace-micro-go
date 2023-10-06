package http

import (
	"auth-service/auth"
	"auth-service/config"
	"auth-service/shared/middleware"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers, conf config.Config) {
	authGroup.POST("/login", h.Login)
	authGroup.POST("/register", h.Register)

	authGroup.GET("/me", middleware.Auth(conf), h.GetMe)
	authGroup.GET("/verify-token", middleware.Auth(conf), h.VerifyAccessToken)
}
