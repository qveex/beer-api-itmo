package controller

import (
	"context"
	"main/internal/domain"
	pb "main/pkg/api"
)

// Server ...
type Server struct {
	pb.UnimplementedCatalogServer
	Service *domain.CatalogService
}

// GetBeers ...
func (s *Server) GetBeers(ctx context.Context, req *pb.GetBeersRequest) (*pb.GetBeersResponse, error) {
	result, err := s.Service.GetAllBeers()
	return &pb.GetBeersResponse{Beers: result}, err
}

// GetBeer ...
func (s *Server) GetBeer(ctx context.Context, req *pb.GetBeerRequest) (*pb.GetBeerResponse, error) {
	result, err := s.Service.GetBeer(req.BeerId)
	return &pb.GetBeerResponse{Beer: result}, err
}

// CreateBeer ...
func (s *Server) CreateBeer(ctx context.Context, req *pb.CreateBeerRequest) (*pb.CreateBeerResponse, error) {
	result, err := s.Service.CreateNewBeer(req.Beer)
	return &pb.CreateBeerResponse{BeerId: result}, err
}

// UpdateBeer ...
func (s *Server) UpdateBeer(ctx context.Context, req *pb.UpdateBeerRequest) (*pb.UpdateBeerResponse, error) {
	result, err := s.Service.UpdateBeer(req.BeerId, req.Beer)
	return &pb.UpdateBeerResponse{Beer: result}, err
}

// DeleteBeer ...
func (s *Server) DeleteBeer(ctx context.Context, req *pb.DeleteBeerRequest) (*pb.DeleteBeerResponse, error) {
	result, err := s.Service.DeleteBeer(req.BeerId)
	return &pb.DeleteBeerResponse{Beer: result}, err
}
