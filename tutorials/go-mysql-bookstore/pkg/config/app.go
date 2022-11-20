package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// “mysql”, “user:password@/dbname?charset=utf8&parseTime=True&loc=Local”
func Connect() {
	d, err := gorm.Open("mysql", "demo:pass@tcp(127.0.0.1:13306)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to mysql ...")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
