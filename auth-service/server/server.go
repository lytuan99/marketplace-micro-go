package server

import (
	"auth-service/auth/delivery-implement/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	// redisClient
	// awsClient
	logger zerolog.Logger
}

func NewServer(logger zerolog.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) Run(address string) {
	router := gin.Default()

	authHandlers := http.NewAuthHandlers(s.logger)
	http.MapAuthRoutes(router.Group("/auth"), authHandlers)

	router.Run(address)
}
