package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/cmd/env"
)

const (
	dbUserEnv = "DB_USER"
	dbPassEnv = "DB_PASS"
	dbConnEnv = "DB_CONN"
	dbNameEnv = "DB_NAME"
)

var db *gorm.DB

func GetDB() (*gorm.DB, error) {

	if db != nil {
		return db, nil
	}

	cfg := config{
		user: env.GetEnv(dbUserEnv, ""),
		pass: env.GetEnv(dbPassEnv, ""),
		conn: env.GetEnv(dbConnEnv, ""),
		name: env.GetEnv(dbNameEnv, ""),
	}

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.user, cfg.pass, cfg.conn, cfg.name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
