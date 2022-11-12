package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (self *Service) GetHealth(context *gin.Context) {
	requestUrl := self.BaseUrl + "/public/health"
	_, err := http.Get(requestUrl)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Server is not responding"})
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}

func (self *Service) GetHealthSQL(context *gin.Context) {
	requestUrl := self.BaseUrl + "/public/health"
	_, err := http.Get(requestUrl)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Server is not responding"})
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}

func (self *Service) RunBackup(context *gin.Context) {
	requestUrl := self.BaseUrl + "/public/health"
	_, err := http.Get(requestUrl)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Server is not responding"})
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}

func (self *Service) RunServer(context *gin.Context) {
	requestUrl := self.BaseUrl + "/public/health"
	_, err := http.Get(requestUrl)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Server is not responding"})
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}

func (self *Service) StopServer(context *gin.Context) {
	requestUrl := self.BaseUrl + "/public/health"
	_, err := http.Get(requestUrl)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Server is not responding"})
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}
