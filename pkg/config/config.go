package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var envConfig *config

func Config() *config {
	return envConfig
}

func InitConfig() {
	envConfig = &config{}
	if err := envconfig.Process("", envConfig); err != nil {
		log.Fatalf("Failed to load the env configuration: %v", err)
	}
}

type config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
	}
	Database struct {
		Host     string `envconfig:"DB_HOST" required:"true"`
		Port     int    `envconfig:"DB_PORT" default:"27017"`
		Username string `envconfig:"DB_USERNAME" required:"true"`
		Password string `envconfig:"DB_PASSWORD" required:"true"`
		Name     string `envconfig:"DB_NAME" default:"it_store"`
		Products struct {
			Collection string `envconfig:"DB_PRODUCTS_COLLECTION" default:"products"`
		}
	}
}
