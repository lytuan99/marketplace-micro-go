package server

import (
	"auth-service/auth/delivery-imp/http"
	"auth-service/config"
	db "auth-service/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	// redisClient
	// awsClient
	conf   config.Config
	store  db.Store
	logger zerolog.Logger
}

func NewServer(conf config.Config, store db.Store, logger zerolog.Logger) *Server {
	return &Server{store: store, logger: logger, conf: conf}
}

func (s *Server) Run(address string) {
	router := gin.Default()

	authHandlers := http.NewAuthHandlers(s.store, s.logger)
	http.MapAuthRoutes(router.Group("/auth"), authHandlers)

	router.Run(address)
}
