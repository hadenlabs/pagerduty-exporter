package config

import (
	"strings"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func readConfig() error {
	config := Configuration{}
	viper.SetEnvPrefix(applicationEnvPrefix)
	viper.SetConfigFile(configFile)
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		return err
	}

	if err := godotenv.Load(); err != nil {
		log.Infof("unable to load .env file: %s %s", applicationName, err)
	}

	if err := env.Parse(&config); err != nil {
		log.WithError(err).Info("parse env value error")
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		return err
	}
	return nil
}

// ReadConfig read config of viper, execute config and also return error.
func ReadConfig() error {
	return readConfig()
}
