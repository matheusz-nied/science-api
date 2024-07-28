package api

import (
	"nied-science/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, apodService service.APODService) {
	apod := r.Group("/apod")
	{
		apod.GET("/", GetAPODs)
	}
}
