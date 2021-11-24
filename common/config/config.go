/*
	Config is the definition of all our internal config toml files.
*/
package config

import (
	"io"

	"github.com/pelletier/go-toml"
)

type Config struct {
	API struct {
		ShowGQLPlayground bool
	}
}

func Decode(c *Config, d io.Reader) error {
	return toml.NewDecoder(d).Decode(c)
}
