package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConn() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:root@tcp(127.0.0.1:60082)/course2?parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	return db
}
