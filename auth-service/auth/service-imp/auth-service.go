package service

import (
	"auth-service/auth"
	db "auth-service/db/sqlc"
	"auth-service/shared/utils"
	"context"

	"github.com/rs/zerolog"
)

type authService struct {
	store  db.Store
	logger zerolog.Logger
}

var _ auth.Service = (*authService)(nil)

func NewAuthService(store db.Store, logger zerolog.Logger) auth.Service {
	return &authService{
		store:  store,
		logger: logger,
	}
}

// Login implements auth.Service.
func (a *authService) Login(username string, password string) {
	a.logger.Info().Msg("hello")
}

// Register implements auth.Service.
func (h *authService) Register(ctx context.Context, req auth.RegisterRequest) (*db.User, error) {
	h.logger.Info().Msgf("Register payload %+v", req)
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	params := db.CreateUserParams{
		Username:    req.Username,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
	}

	user, err := h.store.CreateUser(ctx, params)

	return &user, err
}

// GetMe implements auth.Service.
func (*authService) GetMe() {
	panic("unimplemented")
}

// Logout implements auth.Service.
func (*authService) Logout() {
	panic("unimplemented")
}

// VerifyToken implements auth.Service.
func (*authService) VerifyToken(accessToken string) {
	panic("unimplemented")
}
