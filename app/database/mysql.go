package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Mysql *gorm.DB

// Doc: https://gorm.io/zh_CN/docs/query.html
func init() {
	var err error
	Mysql, err = gorm.Open("mysql", "root:develop@tcp(127.0.0.1:3306)/api?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Mysql.Error != nil {
		fmt.Printf("database error %v", Mysql.Error)
	}

	Mysql.LogMode(true)
}
