package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

}
