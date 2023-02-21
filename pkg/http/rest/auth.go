package rest

import (
	"github.com/erayaydin/series-tracker/pkg/auth"
	"github.com/erayaydin/series-tracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	log logger.LogInfoFormat
	svc auth.Service
}

func AuthController(log logger.LogInfoFormat, svc auth.Service) *authController {
	return &authController{log, svc}
}

func (a authController) Login(ctx *gin.Context) {
	var req auth.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		a.log.Errorf("login request failed with reason: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok := a.svc.Login(req.Username, req.Password); !ok {
		ctx.Status(http.StatusForbidden)
		return
	}
	ctx.Status(http.StatusOK)
}

func (a authController) Register(ctx *gin.Context) {
	var req auth.Register
	if err := ctx.ShouldBindJSON(&req); err != nil {
		a.log.Errorf("register request failed with reason: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok := a.svc.Register(req.Username, req.Password, req.Password, req.Name); !ok {
		ctx.Status(http.StatusForbidden)
		return
	}
	ctx.Status(http.StatusOK)
}
