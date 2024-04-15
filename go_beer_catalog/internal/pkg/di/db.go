package di

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/internal/env"
)

const (
	dbUserEnv = "DB_USER"
	dbPassEnv = "DB_PASS"
	dbAddrEnv = "DB_ADDR"
	dbPortEnv = "DB_PORT"
	dbNameEnv = "DB_NAME"
)

func (sp *ServiceProvider) GetDb() *gorm.DB {
	if sp.db != nil {
		return sp.db
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		env.GetEnv(dbUserEnv, ""),
		env.GetEnv(dbPassEnv, ""),
		env.GetEnv(dbAddrEnv, ""),
		env.GetEnv(dbPortEnv, ""),
		env.GetEnv(dbNameEnv, ""),
	)

	var err error
	sp.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return sp.db
}
