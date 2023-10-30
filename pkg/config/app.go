package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type MysqlConfig struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Name      string `json:"name"`
	Port      string `json:"port"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc       string `json:"loc"`
}

func Connect() {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var cfg MysqlConfig
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = d
}

