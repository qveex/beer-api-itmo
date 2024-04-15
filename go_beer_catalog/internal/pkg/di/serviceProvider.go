package di

import (
	ssov1 "github.com/Krutov777/protos/gen/go/sso"
	"gorm.io/gorm"
	"main/internal/domain"
	"main/internal/repository"
)

type ServiceProvider struct {
	db             *gorm.DB
	authClient     *ssov1.AuthClient
	catalogService *domain.CatalogService
	authService    *domain.AuthService
	favService     *domain.FavoriteService
	catalogRepo    *repository.CatalogRepository
	authRepo       *repository.AuthRepository
	favRepo        *repository.FavoriteRepository
}

var Provider = &ServiceProvider{}
