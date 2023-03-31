package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"grcp"`
	Port    int32  `default:"9090"`
	DB      struct {
		Use      string `default:"postgres"`
		Postgres []struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"5432"`
			UserName string `default:"postgres"`
			Password string `default:"123"`
			Database string `default:"demo-bamboo"`
		}
	}
	Redis struct {
		Addr     string `default:"localhost:6379"`
		Password string `default:""`
		DB       int32  `default:"0"`
	}
	Contacts struct {
		Name  string `default:"javier Lecca"`
		Email string `default:"leccajavier@gmail.com"`
	}
}

func (c *Config) NewConfig() (*Config, error) {
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
