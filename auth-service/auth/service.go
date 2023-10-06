package auth

import (
	db "auth-service/db/sqlc"
	"context"
	"time"
)

type Service interface {
	Login(ctx context.Context, username, password string) (accessToken string, expiresAt time.Time, err error)
	Register(context.Context, RegisterRequest) (*db.User, error)
	Logout() error
	GetMe(ctx context.Context, username string) (db.User, error)
}
