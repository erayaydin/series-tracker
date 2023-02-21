package server

import (
	"github.com/erayaydin/series-tracker/pkg/auth"
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

func (s *server) authRoutes(api *gin.RouterGroup) {
	var authService auth.Service
	err := s.cont.Invoke(func(as auth.Service) {
		authService = as
	})
	if err != nil {
		return
	}

	authRoutes := api.Group("/auth")
	{
		a := rest.AuthController(s.logger, authService)
		authRoutes.POST("/", a.Login)
		authRoutes.POST("/register", a.Register)
	}
}
