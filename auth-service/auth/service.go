package auth

import (
	db "auth-service/db/sqlc"
	"context"
)

type Service interface {
	Login(username, password string)
	Register(context.Context, RegisterRequest) (*db.User, error)
	Logout()
	VerifyToken(accessToken string)
	GetMe()
}
