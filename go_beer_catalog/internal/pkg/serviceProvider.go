package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/internal/domain"
	"main/internal/env"
	"main/internal/repository"
)

type ServiceProvider struct {
	db      *gorm.DB
	service *domain.CatalogService
	repo    *repository.CatalogRepository
}

var Provider = &ServiceProvider{}

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

func (sp *ServiceProvider) GetCatalogService() *domain.CatalogService {
	if sp.service != nil {
		return sp.service
	}

	sp.service = domain.NewCatalogService(sp.GetCatalogRepository())

	return sp.service
}

func (sp ServiceProvider) GetCatalogRepository() *repository.CatalogRepository {
	if sp.repo != nil {
		return sp.repo
	}

	sp.repo = repository.NewCatalogRepository(sp.GetDb())

	return sp.repo
}

const (
	dbUserEnv = "DB_USER"
	dbPassEnv = "DB_PASS"
	dbAddrEnv = "DB_ADDR"
	dbPortEnv = "DB_PORT"
	dbNameEnv = "DB_NAME"
)
