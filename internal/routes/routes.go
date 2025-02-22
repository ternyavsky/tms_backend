package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/middlewares"
)

func InitRoutes(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())
	api := r.Group("api/v1")
	{
		links := api.Group("links")
		InitLinksRoutes(links)
	}
}
