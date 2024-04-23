package di

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
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
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		env.GetEnv(dbUserEnv, ""),
		env.GetEnv(dbPassEnv, ""),
		env.GetEnv(dbAddrEnv, ""),
		env.GetEnv(dbPortEnv, ""),
		env.GetEnv(dbNameEnv, ""),
	)

	var err error
	sp.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return sp.db
}
