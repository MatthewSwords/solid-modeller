package main

import (
	"context"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/solid-modeller/common/config"
	"github.com/solid-modeller/common/helper"
	"github.com/solid-modeller/modeller/api"
	"github.com/tablenu/servitor/common/db"
	"github.com/tablenu/servitor/pad/migrate"

	"github.com/spf13/pflag"
)

var (
	configFile = ""
	addr       = "0.0.0.0:8080"
)

func init() {
	pflag.StringVarP(&configFile, "config", "c", stringEnv("API_CONFIG", configFile), "config file")
	pflag.Parse()
}

func main() {
	// Default level for this project is debug.  All log messages will show up by default
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	c, err := helper.LoadConfig(configFile)
	if err != nil {
		log.Error().Err(err).Msg("failed to load config: " + err.Error())
		os.Exit(2)
	}

	log.Info().Msg("Database Setup...")

	// err := initDB(c)
	if err != nil {
		log.Error().Err(err).Msg("db migrate failed: " + err.Error())
		os.Exit(1)
	}

	// Initialize Firebase
	// app := initFirebaseController(c)

	// This needs to be set before using stripe
	// stripe.Key = c.STRIPE.StripeKey

	r, err := api.New(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to configure handler")
		os.Exit(2)
	}

	log.Info().Str("addr", addr).Msg("listening")
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Error().Err(err).Msg("serve failed")
		os.Exit(1)
	}
}

// func initFirebaseController(c *config.Config) *firebase.App {
// 	firebaseApp := firebase_controller.InitializeWithServiceAccount(c)
// 	return firebaseApp
// }

func initDB(c *config.Config) (*db.Client, error) {
	var err error

	dbClient, err := db.New(c)

	if err != nil {
		return nil, err
	}
	log.Info().Msg("connecting to DB: " + c.DB.URL)
	err = dbClient.Health(context.Background(), c.DB.StartupTimeout)
	if err != nil {
		return nil, err
	}

	// while always migrate
	//if migrateDB || migrateOnly {
	log.Info().Msg("starting to migrate...")
	v, err := migrate.Migrate(dbClient, "")
	if err != nil {
		log.Error().Err(err).Msg("failed migrate db: " + err.Error())
		os.Exit(2)
	}
	if v > 0 {
		log.Info().Int64("version", v).Msg("DB updated")
	}

	//if migrateOnly {
	//	os.Exit(0)
	//}
	//}

	log.Info().Msg("finished migrate")
	return dbClient, nil
}

func stringEnv(value, fallback string) string {
	if v := os.Getenv(value); v != "" {
		return v
	}
	return fallback
}
