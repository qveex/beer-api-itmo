package di

import (
	ssov1 "github.com/Krutov777/protos/gen/go/sso"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log/slog"
	"main/internal/domain"
	"main/internal/repository"
)

type ServiceProvider struct {
	db             *gorm.DB
	server         *grpc.Server
	log            *slog.Logger
	authClient     *ssov1.AuthClient
	catalogService *domain.CatalogService
	authService    *domain.AuthService
	favService     *domain.FavoriteService
	catalogRepo    *repository.CatalogRepository
	authRepo       *repository.AuthRepository
	favRepo        *repository.FavoriteRepository
}

var Provider = &ServiceProvider{}
