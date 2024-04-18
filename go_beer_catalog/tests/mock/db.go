package mock

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/env"
	"main/internal/repository"
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
		env.GetEnv(dbUserEnv, "postgres"),
		env.GetEnv(dbPassEnv, "postgres"),
		env.GetEnv(dbAddrEnv, "localhost"),
		env.GetEnv(dbPortEnv, "5432"),
		env.GetEnv(dbNameEnv, "test_catalog"),
	)

	var err error
	sp.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = sp.db.AutoMigrate(&repository.Favorite{}, &repository.Beer{})

	if err != nil {
		panic(err)
	}

	return sp.db
}

func (sp ServiceProvider) GetCatalogRepository() *repository.CatalogRepository {
	if sp.catalogRepo != nil {
		return sp.catalogRepo
	}

	sp.catalogRepo = repository.NewCatalogRepository(sp.GetDb())

	return sp.catalogRepo
}
