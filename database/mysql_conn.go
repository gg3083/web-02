package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var SqlDB *gorm.DB

func init() {
	var err error
	SqlDB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/python?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err.Error())
	}

}

func closeDB() {
	defer SqlDB.Close()
}
