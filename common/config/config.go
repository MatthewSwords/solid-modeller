/*
	Config is the definition of all our internal config toml files.
*/
package config

import (
	"io"

	"github.com/pelletier/go-toml"
	"github.com/rs/cors"
)

type Config struct {
	API struct {
		ShowGQLPlayground bool
	}
	CORS cors.Options
}

func Decode(c *Config, d io.Reader) error {
	return toml.NewDecoder(d).Decode(c)
}
