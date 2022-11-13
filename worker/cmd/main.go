package main

import (
	"log"

	"github.com/DanArmor/MovieDB_bots/pkg/config"
	"github.com/DanArmor/MovieDB_bots/pkg/controllers"
	"github.com/DanArmor/MovieDB_bots/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config parse! ", err)
	}

	config.AdminPass = utils.HashPassword(config.AdminPass)
	config.ServerURL = "https://" + config.Domain
	
	service := controllers.Service{
		Config: config,
	}

	router := gin.Default()

	bots := router.Group("/bots")
	bots.Use(service.ValidateAdmin)
	bots.GET("/health", service.GetHealth)
	bots.GET("/healthSQL", service.GetHealthSQL)
	bots.GET("/run_server", service.RunServer)
	bots.GET("/stop_server", service.StopServer)
	bots.GET("/run_backup", service.RunBackup)
	bots.Static("/backups", "backups")

	router.RunTLS("127.0.0.1" + config.Port, config.CertPath, config.KeyPath)
}
