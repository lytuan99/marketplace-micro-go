// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Username    string           `json:"username"`
	Password    string           `json:"password"`
	PhoneNumber string           `json:"phone_number"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}
