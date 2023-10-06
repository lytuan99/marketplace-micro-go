package utils

import (
	"auth-service/shared/response"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type JwtPayload struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// This is simple generate token function, in fact..it's will be more complicated
// TODO: need to pass payload param (can be an interface) instead of using only username
func GenerateToken(username string, duration time.Duration, secret string) (accessToken string, expiresAt time.Time, err error) {
	issuedAt := time.Now()
	expiresAt = issuedAt.Add(duration)

	payload := JwtPayload{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	accessToken, err = jwtToken.SignedString([]byte(secret))

	return accessToken, expiresAt, err
}

func VerifyToken(token string, secretKey string) (*JwtPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, response.ErrInvalidToken
		}

		return []byte(secretKey), nil
	}

	payload := &JwtPayload{}

	_, err := jwt.ParseWithClaims(token, payload, keyFunc)
	if err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			return nil, response.ErrExpiredToken
		}
		log.Default().Println("Error when parse with claim: ", err)
		return nil, err
	}

	return payload, nil
}
