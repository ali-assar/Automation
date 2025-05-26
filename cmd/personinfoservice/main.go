package main

import (
	"backend/internal/api/personinfoservice/admin"
	"backend/internal/api/personinfoservice/credentials"
	"backend/internal/api/personinfoservice/person"
	"backend/internal/config"
	"backend/internal/logger"
	"backend/pkg/services/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.LogLevel, cfg.LogFormat)
	adminCtrl := admin.NewController()
	personCtrl := person.NewController()
	credCtrl := credentials.NewController()
	if err := router.InitRouter(cfg.AppHost, cfg.AppPort, func(r *gin.Engine) {
		router.RegisterPersonInfoRoutes(r, adminCtrl, personCtrl, credCtrl)
	}); err != nil {
		logger.Error("Failed to start PersonInfoService:", err)
		panic(err)
	}
}
