package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectToDB() {
	username := os.Getenv("DB_USERNAME")
	pwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	connectionString := username + ":" + pwd + "@/" + dbName + "?charset=utf8&parseTime=true&loc=Local"

	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
