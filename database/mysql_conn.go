package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"log"
)

var SqlDB *gorm.DB

func init() {
	//var err error
	//SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/python?parseTime=true")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//err = SqlDB.Ping()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	SqlDB, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/python?parseTime=true")
	if err != nil {
		panic("连接数据库失败")
	}
	defer SqlDB.Close()
}
