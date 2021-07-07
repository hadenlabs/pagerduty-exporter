package config

const (
	configFile           = ".env"
	applicationName      = "PAGERDUTYEXPORTER"
	applicationEnvPrefix = "PAGERDUTYEXPORTER"
)

// Configuration struct for config.
type Configuration struct {
	AppEnv string `env:"PAGERDUTYEXPORTER_APP_ENV" envDefault:"test"`
}
