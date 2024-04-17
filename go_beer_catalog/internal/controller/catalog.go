package controller

import (
	"context"
	pb "main/pkg/api"
)

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
	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}
	result, err := s.CatalogService.CreateNewBeer(token, req.Beer)
	return &pb.CreateBeerResponse{BeerId: result}, err
}

// UpdateBeer ...
func (s *Server) UpdateBeer(ctx context.Context, req *pb.UpdateBeerRequest) (*pb.UpdateBeerResponse, error) {
	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}
	result, err := s.CatalogService.UpdateBeer(token, req.BeerId, req.Beer)
	return &pb.UpdateBeerResponse{Beer: result}, err
}

// DeleteBeer ...
func (s *Server) DeleteBeer(ctx context.Context, req *pb.DeleteBeerRequest) (*pb.DeleteBeerResponse, error) {
	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}
	result, err := s.CatalogService.DeleteBeer(token, req.BeerId)
	return &pb.DeleteBeerResponse{Beer: result}, err
}
