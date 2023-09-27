package http

import (
	"auth-service/auth"
	"auth-service/auth/service-implement"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type authHandlers struct {
	service auth.Service
	logger  zerolog.Logger
}

var _ auth.Handlers = (*authHandlers)(nil)

func NewAuthHandlers(logger zerolog.Logger) *authHandlers {
	service := service.NewAuthService(logger)
	return &authHandlers{
		service: service,
		logger:  logger,
	}
}

// Login implements auth.Handlers.
func (h *authHandlers) Login(ctx *gin.Context) {
	h.service.Login("Lytuan", "1234")
}
