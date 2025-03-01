package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/controllers"
	"github.com/ternyavsky/tms_backend/internal/repo"
	"github.com/ternyavsky/tms_backend/internal/services"
	"github.com/ternyavsky/tms_backend/pkg/db"
)

func InitSearchRoutes(router *gin.RouterGroup) {
	linkRepository := repo.NewLinkRepository(db.DB)
	searchService := services.NewSearchService(linkRepository)
	searchController := controllers.NewSearchController(*searchService)
	router.GET("/:channelUrl", searchController.SearchChannel)
	router.GET("/status/:linkId", searchController.GetSearchStatus)
}
