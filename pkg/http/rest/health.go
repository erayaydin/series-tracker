package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthController struct {
}

func HealthController() *healthController {
	return &healthController{}
}

func (h healthController) GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
