package di

import (
	"fmt"
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
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.GetEnv(dbAddrEnv, ""),
		env.GetEnv(dbUserEnv, ""),
		env.GetEnv(dbPassEnv, ""),
		env.GetEnv(dbNameEnv, ""),
		env.GetEnv(dbPortEnv, ""),
	)

	var err error
	sp.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return sp.db
}
