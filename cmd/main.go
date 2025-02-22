package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ternyavsky/tms_backend/cmd/docs"
	"github.com/ternyavsky/tms_backend/config"
	"github.com/ternyavsky/tms_backend/internal/routes"

	"github.com/ternyavsky/tms_backend/pkg/db"
)

func setup() {
	config.LoadConfig()
	err := db.InitDatabse()
	if err != nil {
		log.Fatalf("Failed init db, error: %v", err)
	}
}

// @title Telegram Metric System Backend
// @version 1.0
// @description API для Telegram Metric System
// @BasePath /
func main() {
	setup()
	if config.GlobalConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	ginDaemon := gin.Default()
	routes.InitRoutes(ginDaemon)
	ginDaemon.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := ginDaemon.Run(config.GlobalConfig.ServerHost + ":" + config.GlobalConfig.ServerPort); err != nil {
		log.Fatalf("Server shutdown with error: %v ", err)
	}
}
