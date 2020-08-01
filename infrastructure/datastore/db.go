package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMySqlDb() *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"user",
		"password",
		"127.0.0.1",
		"3306",
		"go-rest-api",
	)

	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
