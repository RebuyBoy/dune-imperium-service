package configs

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	MongoURI string
	AppName  string
	Port     string
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() (*Config, error) {
	var initErr error
	once.Do(func() {
		cfg = &Config{}

		cfg.MongoURI, initErr = getRequiredEnv("MONGODB_URI")
		if initErr != nil {
			return
		}

		cfg.AppName, initErr = getRequiredEnv("APP_NAME")
		if initErr != nil {
			return
		}

		cfg.Port, initErr = getRequiredEnv("PORT")
		if initErr != nil {
			return
		}

	})
	if initErr != nil {
		return nil, initErr
	}
	return cfg, nil
}

func getRequiredEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", key)
	}
	return value, nil
}
