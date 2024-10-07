package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(viper.GetInt32("LOG_LEVEL")))
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true, // Enable colors
		FullTimestamp: true, // Display full timestamp
	})

	return log
}
