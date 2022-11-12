package controllers

import "gorm.io/gorm"

type Service struct {
	AdminPass string
	Domain    string
	BaseUrl   string
	DB        *gorm.DB
}

func (self *Service) ConnectDB() {
	return
}
