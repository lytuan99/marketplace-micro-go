package response

import "errors"

// Postgresql error code
const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)
