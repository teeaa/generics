package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Service struct {
		Name    string `env-default:"generics-service" env:"SERVICE_NAME"`
		Address string `env-default:"127.0.0.1" env:"SERVICE_ADDRESS"`
		Port    int    `env-default:"8080" env:"SERVICE_PORT"`
	} `yaml:"service"`
	Database struct {
		Host     string `env-default:"127.0.0.1" env:"DATABASE_HOST"`
		Port     int    `env-default:"5432" env:"DATABASE_PORT"`
		User     string `env-default:"postgres" env:"DATABASE_USER"`
		Password string `env-default:"postgres" env:"DATABASE_PASSWORD"`
		Name     string `env-default:"generics" env:"DATABASE_NAME"`
	} `yaml:"database"`
}

func New() (*Config, error) {
	var conf Config

	err := cleanenv.ReadConfig("config.yml", &conf)
	if err != nil {
		err = cleanenv.ReadEnv(&conf)
		if err != nil {
			return nil, err
		}
	}

	return &conf, nil
}
