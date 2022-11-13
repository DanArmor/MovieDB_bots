package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (self *Service) GetHealth(context *gin.Context) {
	requestUrl := strings.TrimSpace(self.Config.ServerURL + "/public/health")
	_, err := http.Get(requestUrl)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Server is not responding: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}

func (self *Service) GetHealthSQL(context *gin.Context) {
	db, err := gorm.Open(mysql.Open(strings.TrimSpace(self.Config.SqlUrl)))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Sql is not responding on connection: " + err.Error()})
		return
	}
	if db.Exec("SHOW STATUS").Error != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Sql is not responding on show status: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Up and running!"})
}

func (self *Service) RunBackup(context *gin.Context) {
	now := time.Now()
	fileName := now.Format("2006.1.2_15.4.5") + ".sql"
	output, err := exec.Command("mysqldump", "--opt", fmt.Sprintf("--user=%s", self.Config.SqlUser), fmt.Sprintf("--password=%s", self.Config.SqlPass), "movies").Output()
	println(output)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Error during backup: " + err.Error()})
		return
	}
	if err := os.WriteFile("./backups/" + fileName, output, 0644); err != nil{
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Error during backup: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : self.Config.ServerURL + "/bots/backups/" + fileName})
}

func (self *Service) RunServer(context *gin.Context) {
	requestUrl := strings.TrimSpace(self.Config.ServerURL + "/public/health")
	answ, _ := http.Get(requestUrl)
	if answ.StatusCode == 200 {
		context.JSON(http.StatusOK, gin.H{"status": "Ok", "desc" : "Server is already running"})
		return
	}
	cmd := exec.Command("./moviedb")
	cmd.Dir = self.Config.PathServer
	err := cmd.Start()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Error during backup: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Is running"})
}

func (self *Service) StopServer(context *gin.Context) {
	if err := exec.Command("pkill", "-SIGUSR1", "-f", "moviedb").Run(); err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error", "desc" : "Error during stop - can't kill: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "OK", "desc" : "Server was stopped"})
}
