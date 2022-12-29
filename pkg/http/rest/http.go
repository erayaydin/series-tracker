package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthController struct {
}

func HealthController() *healthController {
	return &healthController{}
}

func (h healthController) GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
