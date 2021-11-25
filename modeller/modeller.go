package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/solid-modeller/MatthewSwords/common/helper"
	"github.com/solid-modeller/MatthewSwords/modeller/api"

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

func stringEnv(value, fallback string) string {
	if v := os.Getenv(value); v != "" {
		return v
	}
	return fallback
}
