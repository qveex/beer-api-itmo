package controller

import (
	"context"
	"main/internal/domain"
	pb "main/pkg/api"
)

// Server ...
type Server struct {
	pb.UnimplementedCatalogServer
	CatalogService  *domain.CatalogService
	FavoriteService *domain.FavoriteService
}

// GetBeers ...
func (s *Server) GetBeers(ctx context.Context, req *pb.GetBeersRequest) (*pb.GetBeersResponse, error) {
	result, err := s.CatalogService.GetAllBeers(
		req.Limit,
		req.Name,
		req.Brand,
		req.Type,
		req.Deg,
		req.Sweet,
	)
	return &pb.GetBeersResponse{Beers: result}, err
}

// GetBeer ...
func (s *Server) GetBeer(ctx context.Context, req *pb.GetBeerRequest) (*pb.GetBeerResponse, error) {
	result, err := s.CatalogService.GetBeer(req.BeerId)
	return &pb.GetBeerResponse{Beer: result}, err
}

// CreateBeer ...
func (s *Server) CreateBeer(ctx context.Context, req *pb.CreateBeerRequest) (*pb.CreateBeerResponse, error) {
	result, err := s.CatalogService.CreateNewBeer(req.Beer)
	return &pb.CreateBeerResponse{BeerId: result}, err
}

// UpdateBeer ...
func (s *Server) UpdateBeer(ctx context.Context, req *pb.UpdateBeerRequest) (*pb.UpdateBeerResponse, error) {
	result, err := s.CatalogService.UpdateBeer(req.BeerId, req.Beer)
	return &pb.UpdateBeerResponse{Beer: result}, err
}

// DeleteBeer ...
func (s *Server) DeleteBeer(ctx context.Context, req *pb.DeleteBeerRequest) (*pb.DeleteBeerResponse, error) {
	result, err := s.CatalogService.DeleteBeer(req.BeerId)
	return &pb.DeleteBeerResponse{Beer: result}, err
}

// SetFavorite ...
func (s *Server) SetFavorite(ctx context.Context, req *pb.SetFavoriteRequest) (*pb.SetFavoriteResponse, error) {
	err := s.FavoriteService.SetFavorite(req.UserId, req.BeerId)
	return &pb.SetFavoriteResponse{}, err
}

// GetFavorites ...
func (s *Server) GetFavorites(ctx context.Context, req *pb.GetFavoritesRequest) (*pb.GetFavoritesResponse, error) {
	result, err := s.FavoriteService.GetFavorites(req.UserId)
	return &pb.GetFavoritesResponse{Beers: result}, err
}

func (s *Server) DeleteFavorite(ctx context.Context, req *pb.DeleteFavoriteRequest) (*pb.DeleteFavoriteResponse, error) {
	err := s.FavoriteService.DeleteFavorite(req.UserId, req.BeerId)
	return &pb.DeleteFavoriteResponse{}, err
}
