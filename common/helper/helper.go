/*
	Has basic universal functions that can be useful though the application.
*/
package helper

import (
	"os"

	"github.com/solid-modeller/common/config"
)

// Contains is a basic method to see if an array contains a specified string
func Contains(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

// LoadConfig will load a toml config
func LoadConfig(configFilePath string) (*config.Config, error) {
	var c config.Config
	f, err := os.Open(configFilePath)
	if err != nil {
		return &config.Config{}, err
	}

	err = config.Decode(&c, f)
	if err != nil {
		return &config.Config{}, err
	}

	return &c, nil
}
