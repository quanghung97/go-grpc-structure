package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	DB       int32  `default:"0"`
}
type Config struct {
	AppName string
	Port    string
	DB      []PostGres
	Redis   Redis
}

func (c *Config) NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	return &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		DB: []PostGres{
			{
				Host:     os.Getenv("HOST_POSTGRES"),
				Port:     os.Getenv("PORT_POSRGRES"),
				UserName: os.Getenv("USER_NAME_POSTGRES"),
				Password: os.Getenv("PASS_WORD_POSTGRES"),
				Database: os.Getenv("DATA_BASE"),
			},
		},
		Redis: Redis{
			Addr: os.Getenv("Addr"),
		},
	}, nil
}
