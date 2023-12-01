package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gDb *gorm.DB
)

func Init() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("connect to sqlite")
	}
	_ = db.Callback().Row().After("all").Register("all", func(db *gorm.DB) {
		fmt.Printf("db_name:%s, db.sm_name:%s\n", db.Name(), db.Statement.Name())
	})
	db.Callback().Query().After("all").Register("all", func(db *gorm.DB) {
		fmt.Printf("db_name:%s, db.sm_name:%s\n", db.Name(), db.Statement.Name())
	})
	return nil
}

func GetDB() *gorm.DB {
	return gDb
}
