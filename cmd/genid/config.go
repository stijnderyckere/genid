package main

import (
	"github.com/koding/multiconfig"
)

// Config is the main configuration interface. Other configuration
// should be part of this one
type Config struct {
	LogLevel string `default:"debug"`
}

func Load(cfg interface{}) error {
	loader := multiconfig.New()
	if err := loader.Load(cfg); err != nil {
		return err
	}

	if err := loader.Validate(cfg); err != nil {
		return err
	}

	return nil
}
