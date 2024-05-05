package models

import "gorm.io/gorm"

type App struct {
	DB *gorm.DB
}

type DatabaseCreadentials struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	User     string `json:"user"`
}
