package main

import (
	"os"

	"github.com/oalexander6/web-app-template/config"
	"github.com/oalexander6/web-app-template/httpserver"
	"github.com/oalexander6/web-app-template/logger"
	"github.com/oalexander6/web-app-template/models"
	"github.com/oalexander6/web-app-template/store/postgres"
	"github.com/rs/zerolog"
)

func main() {
	logger.Init(zerolog.DebugLevel, os.Stdout)

	c := config.New()
	if err := c.Validate(); err != nil {
		logger.Log.Fatal().Msgf("Invalid configuration: %s", err.Error())
	}

	logger.Log.Info().Interface("config", c).Msg("Config initialized")

	var store models.Store

	switch c.StoreType {
	case config.STORE_TYPE_POSTGRES:
		store = postgres.New(c.PostgresOpts)
	default:
		logger.Log.Fatal().Msgf("Invalid store type: %s", c.StoreType)
	}

	defer store.Close()

	app := httpserver.New(c, store)

	app.Run()
}
