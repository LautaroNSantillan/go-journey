package config

import (
	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialect/mysql"
)

var (
	db *gorm.DB
)

func ConnectToDB() {
	d, err := gorm.Open("mysql", "username:password@/nameoftable?charsrt=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
