package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostGres struct {
	Enabled  bool   `env:"ENABLED"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	UserName string `env:"USER_NAME"`
	Password string `env:"PASS_WORD"`
	Database string `env:"DATABASE"`
}
type Config struct {
	AppName string     `env:"APP_NAME"`
	Port    string     `env:"PORT"`
	DB      []PostGres `env:"DB"`
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
	}, nil
}
