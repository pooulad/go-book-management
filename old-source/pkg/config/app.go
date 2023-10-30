package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "poulad:poulad@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if (err) != nil {
		panic(err)
	}
	db = d
}
