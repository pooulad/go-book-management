package models

import (
	"github.com/pooulad/go-book-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `gorm:"" json:"publication"`
}

func init(){
	config.Connect()
	
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}