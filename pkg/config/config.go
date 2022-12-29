package config

import "github.com/jinzhu/configor"

type Config struct {
	Port   string `default:"8080"`
	Logger struct {
		Use         string `default:"zaplogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"debug"`
		FileName    string `default:"application.log"`
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "pkg/config/config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
