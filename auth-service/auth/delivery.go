package auth

import "github.com/gin-gonic/gin"

type Handlers interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetMe(ctx *gin.Context)
	VerifyAccessToken(ctx *gin.Context)
}
