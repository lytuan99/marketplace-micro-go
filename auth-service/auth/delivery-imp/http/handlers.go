package http

import (
	"auth-service/auth"
	"auth-service/auth/service-imp"
	db "auth-service/db/sqlc"
	"auth-service/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog"
)

type authHandlers struct {
	service auth.Service
	logger  zerolog.Logger
}

var _ auth.Handlers = (*authHandlers)(nil)

func NewAuthHandlers(store db.Store, logger zerolog.Logger) *authHandlers {
	service := service.NewAuthService(store, logger)
	return &authHandlers{
		service: service,
		logger:  logger,
	}
}

type userResponse struct {
	Username    string           `json:"username"`
	PhoneNumber string           `json:"phone_number"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

// Login implements auth.Handlers.
func (h *authHandlers) Login(ctx *gin.Context) {
	h.service.Login("Lytuan", "1234")
}

// Register implements auth.Handlers.
func (h *authHandlers) Register(ctx *gin.Context) {
	var req auth.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseError(ctx, err, http.StatusBadRequest, "")
		return
	}

	user, err := h.service.Register(ctx, req)
	if err != nil {
		response.ResponseError(ctx, err, 0, "")
		return
	}

	userRes := userResponse{
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response.ResponseData{
		Message: "Register successfully",
		Data:    userRes,
	})
}
