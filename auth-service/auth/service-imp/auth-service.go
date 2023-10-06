package service

import (
	"auth-service/auth"
	"auth-service/config"
	db "auth-service/db/sqlc"
	"auth-service/shared/utils"
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
)

type authService struct {
	conf   config.Config
	store  db.Store
	logger zerolog.Logger
}

var _ auth.Service = (*authService)(nil)

func NewAuthService(conf config.Config, store db.Store, logger zerolog.Logger) auth.Service {
	return &authService{
		conf:   conf,
		store:  store,
		logger: logger,
	}
}

// Login implements auth.Service.
func (a *authService) Login(ctx context.Context, username string, password string) (string, time.Time, error) {
	user, err := a.store.GetUserByUsername(ctx, username)
	if err != nil {
		return "", time.Time{}, err
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		return "", time.Time{}, errors.New("password does not match")
	}

	accessToken, expiresAt, err := utils.GenerateToken(username, a.conf.AccessTokenDuration, a.conf.TokenSymmetricKey)

	return accessToken, expiresAt, err
}

// Register implements auth.Service.
func (a *authService) Register(ctx context.Context, req auth.RegisterRequest) (*db.User, error) {
	a.logger.Info().Msgf("Register payload %+v", req)
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	params := db.CreateUserParams{
		Username:    req.Username,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
	}

	user, err := a.store.CreateUser(ctx, params)

	return &user, err
}

// GetMe implements auth.Service.
func (a *authService) GetMe(ctx context.Context, username string) (db.User, error) {
	user, err := a.store.GetUserByUsername(ctx, username)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Logout implements auth.Service.
func (*authService) Logout() error {
	panic("unimplemented")
}
