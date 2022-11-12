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
	
	service := controllers.Service{
		AdminPass: config.AdminPass,
		Domain: config.Domain,
		BaseUrl: "https://" + config.Domain + config.ServerPort,
	}

	router := gin.Default()

	admin := router.Group("/admin")
	admin.Use(service.ValidateAdmin)

	admin.GET("/health", service.GetHealth)

	router.RunTLS(config.Port, config.CertPath, config.KeyPath)
}
