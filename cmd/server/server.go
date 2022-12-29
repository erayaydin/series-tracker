package server

import (
	"fmt"
	"github.com/erayaydin/series-tracker/pkg/config"
	"github.com/erayaydin/series-tracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type server struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *server {
	return &server{
		router: e,
		cont:   c,
		logger: l,
	}
}

func (s *server) Start() error {
	var cfg *config.Config
	if err := s.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return s.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
