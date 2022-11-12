package controllers

import (
	"net/http"

	"github.com/DanArmor/MovieDB_bots/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (self *Service) ValidateAdmin(context *gin.Context) {
	pass := context.Request.Header.Get("pass")
	if pass == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No admin password!"})
		return
	}

	if utils.CheckPasswordHash(pass, self.AdminPass) == false {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong admin password!"})
		return
	}

	context.Next()
}