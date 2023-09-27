package main

import (
	"auth-service/server"
	"os"

	"github.com/rs/zerolog"
)

func main() {

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	server := server.NewServer(logger)
	server.Run(":8001")
}
