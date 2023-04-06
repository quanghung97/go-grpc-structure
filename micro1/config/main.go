package config

import (
	"log"

	"github.com/jinzhu/configor"
)

type PostGres struct {
	Enabled  bool
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

type Redis struct {
	Addr     string `default:"localhost:6379"`
	Password string `default:""`
	DB       int64  `default:"0"`
}
type Config struct {
	AppName string `default:"grcp"`
	Port    int32  `default:"8000"`
	DB      struct {
		Use      string `default:"postgres"`
		Postgres []PostGres
	}
	Redis Redis
}

func (c *Config) NewConfig() (*Config, error) {
	err := configor.Load(c, "config.yml")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	return c, nil
}
