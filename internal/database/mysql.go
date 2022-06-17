package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnDatabase() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:60082)/course2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect db")
	}
	return db
}
