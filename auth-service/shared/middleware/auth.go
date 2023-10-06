package middleware

import (
	"auth-service/config"
	"auth-service/shared/response"
	"auth-service/shared/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationPayloadKey = "tokenPayload"
)

func Auth(conf config.Config) gin.HandlerFunc {
	secretKey := conf.TokenSymmetricKey

	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			response.ResponseError(ctx, nil, http.StatusUnauthorized, "")
			ctx.Abort()
			return
		}

		fields := strings.Fields(auth)
		if len(fields) < 2 || fields[0] != "Bearer" {
			response.ResponseError(ctx, nil, http.StatusUnauthorized, "")
			ctx.Abort()
			return
		}

		payload, err := utils.VerifyToken(fields[1], secretKey)
		if err != nil {
			response.ResponseError(ctx, nil, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
