package mock

import (
	"gorm.io/gorm"
	"main/internal/domain"
	"main/internal/repository"
)

type ServiceProvider struct {
	db             *gorm.DB
	catalogService *domain.CatalogService
	catalogRepo    *repository.CatalogRepository
}

var Provider = &ServiceProvider{}
