package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.WithError(err).Error("failed to start the application")
		os.Exit(1)
	}
}

func run() error {
	logrus.Info("loading main config")
	conf := &Config{}
	err := Load(conf)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to load main config")
	}

	// Set loglevel according to config value
	logrus.WithField("loglevel", conf.LogLevel).Info("setting logrus loglevel")
	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.Fatalf("unable to parse loglevel %v", err)
	}
	logrus.SetLevel(logLevel)

	return nil
}
