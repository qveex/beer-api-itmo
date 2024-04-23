package controller

import (
	"main/internal/domain"
	pb "main/pkg/api"
)

// Server ...
type Server struct {
	pb.UnimplementedCatalogServer
	CatalogService  *domain.CatalogService
	FavoriteService *domain.FavoriteService
}
