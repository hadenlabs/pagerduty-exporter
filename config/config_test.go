package config

import (
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	var err error
	if err = InitializeViper(); err != nil {
		log.Debug(err)
	}
	if err = godotenv.Load("./fixtures/test.env"); err != nil {
		log.Debug("not found env")
	}
}

func TestEnvConfig(t *testing.T) {
	assert.Equal(t, viper.GetString("APP_ENV"), "test", viper.GetString("APP_ENV"))
}
