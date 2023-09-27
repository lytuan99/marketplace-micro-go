package service

import (
	"auth-service/auth"

	"github.com/rs/zerolog"
)

type authService struct {
	logger zerolog.Logger
}

var _ auth.Service = (*authService)(nil)

func NewAuthService(logger zerolog.Logger) auth.Service {
	return &authService{
		logger: logger,
	}
}

// Login implements auth.Service.
func (a *authService) Login(username string, password string) {
	a.logger.Info().Msg("hello")
}

// GetMe implements auth.Service.
func (*authService) GetMe() {
	panic("unimplemented")
}

// Logout implements auth.Service.
func (*authService) Logout() {
	panic("unimplemented")
}

// Register implements auth.Service.
func (*authService) Register() {
	panic("unimplemented")
}

// VerifyToken implements auth.Service.
func (*authService) VerifyToken(accessToken string) {
	panic("unimplemented")
}
