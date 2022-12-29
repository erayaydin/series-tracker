package server

import (
	"github.com/erayaydin/series-tracker/pkg/http/rest"
	"github.com/gin-gonic/gin"
)

func (s *server) RegisterRoutes() {
	apiV1 := s.router.Group("api/v1")

	s.healthRoutes(apiV1)
}

func (s *server) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := rest.HealthController()
		healthRoutes.GET("/", h.GetStatus)
	}
}
