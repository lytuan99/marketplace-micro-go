package main

import (
	"auth-service/config"
	db "auth-service/db/sqlc"
	"auth-service/server"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	conf, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal().Msg("Can not load config")
	}

	if conf.Environment == "development" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewSqlStore(connPool)

	server := server.NewServer(conf, store, logger)
	server.Run(conf.HttpServerAddress)
}
