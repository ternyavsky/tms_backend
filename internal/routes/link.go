package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/controllers"
	"github.com/ternyavsky/tms_backend/internal/repo"
	"github.com/ternyavsky/tms_backend/internal/services"
	"github.com/ternyavsky/tms_backend/pkg/db"
)

func InitLinksRoutes(router *gin.RouterGroup) {
	linkRepository := repo.NewLinkRepository(db.DB)
	linkService := services.NewLinkService(linkRepository)
	linkController := controllers.NewLinkController(*linkService)
	router.GET("/", linkController.GetAllLinks)
	router.POST("/", linkController.CreateLink)
}
