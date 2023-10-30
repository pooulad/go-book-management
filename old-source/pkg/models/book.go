package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pooulad/go-bookstore/pkg/config"
)

var db *gorm.DB


type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db.AutoMigrate(&Book{})
}