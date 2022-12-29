package di

import (
	"github.com/erayaydin/series-tracker/pkg/config"
	"github.com/erayaydin/series-tracker/pkg/logger"
	"go.uber.org/dig"
)

var container = dig.New()

func Build() (*dig.Container, error) {

	// Config
	err := container.Provide(config.NewConfig)
	if err != nil {
		return nil, err
	}

	// Logger
	err = container.Provide(logger.NewLogger)
	if err != nil {
		return nil, err
	}

	return container, nil
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
