package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/controllers"
	"github.com/ternyavsky/tms_backend/internal/repo"
	"github.com/ternyavsky/tms_backend/internal/services"
	"github.com/ternyavsky/tms_backend/pkg/db"
)

func InitSessionRoutes(router *gin.RouterGroup) {
	sessionRepository := repo.NewSessionRepository(db.DB)
	sessionService := services.NewSessionService(sessionRepository)
	sessionController := controllers.NewSessionController(*sessionService)
	router.GET("/", sessionController.GetAllSessions)
	router.POST("/", sessionController.CreateSession)
}
