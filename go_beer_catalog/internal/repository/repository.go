package repository

import (
	"gorm.io/gorm"
	pb "main/pkg/api"
)

type CatalogRepository struct {
	db *gorm.DB
}

type AuthRepository struct {
	api pb.CatalogClient // TODO Auth
}

func NewCatalogRepository(db *gorm.DB) *CatalogRepository {
	return &CatalogRepository{
		db: db,
	}
}

func NewAuthRepository(api *pb.CatalogClient) *AuthRepository {
	return &AuthRepository{
		api: *api,
	}
}
