package domain

import "main/internal/repository"

type CatalogService struct {
	repo *repository.CatalogRepository
}

type AuthService struct {
	repo *repository.AuthRepository
}

func NewCatalogService(repo *repository.CatalogRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}
