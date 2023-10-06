package utils

import (
	"auth-service/shared/response"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestVerifyJwtToken(t *testing.T) {
	username := RandomString(6)
	duration := time.Minute
	secretKey := RandomString(32)

	accessToken, expiresAt, err := GenerateToken(username, duration, secretKey)
	require.NoError(t, err)

	require.NotEmpty(t, accessToken)
	require.WithinDuration(t, expiresAt, time.Now(), time.Minute)

	payload, err := VerifyToken(accessToken, secretKey)

	require.NoError(t, err)

	require.NotEmpty(t, payload)
	require.Equal(t, username, payload.Username)
}

func TestExpiredJwtToken(t *testing.T) {
	username := RandomString(6)
	duration := -time.Millisecond
	secretKey := RandomString(32)

	accessToken, _, err := GenerateToken(username, duration, secretKey)
	require.NoError(t, err)

	require.NotEmpty(t, accessToken)

	payload, err := VerifyToken(accessToken, secretKey)

	require.Error(t, err)
	require.EqualError(t, err, response.ErrExpiredToken.Error())
	require.Empty(t, payload)
}
