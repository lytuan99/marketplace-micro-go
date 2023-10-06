package http

import (
	"auth-service/auth"
	"auth-service/auth/service-imp"
	"auth-service/config"
	db "auth-service/db/sqlc"
	"auth-service/shared/middleware"
	"auth-service/shared/response"
	"auth-service/shared/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog"
)

type authHandlers struct {
	service auth.Service
	logger  zerolog.Logger
}

var _ auth.Handlers = (*authHandlers)(nil)

func NewAuthHandlers(conf config.Config, store db.Store, logger zerolog.Logger) *authHandlers {
	service := service.NewAuthService(conf, store, logger)
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

type loginResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// Login implements auth.Handlers.
func (h *authHandlers) Login(ctx *gin.Context) {
	var req auth.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseError(ctx, err, http.StatusBadRequest, "")
		return
	}

	accessToken, expiresAt, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		response.ResponseError(ctx, err, http.StatusBadRequest, "")
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseData{
		Message: "Login successfully",
		Data: loginResponse{
			AccessToken: accessToken,
			ExpiresAt:   expiresAt,
		},
	})
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

func (h *authHandlers) GetMe(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*utils.JwtPayload)
	user, err := h.service.GetMe(ctx, authPayload.Username)
	if err != nil {
		response.ResponseError(ctx, err, 0, "")
		return
	}

	userRes := userResponse{
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, userRes)

}

func (h *authHandlers) VerifyAccessToken(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*utils.JwtPayload)

	ctx.JSON(http.StatusOK, authPayload)
}
