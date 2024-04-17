package controller

import (
	"context"
	pb "main/pkg/api"
)

// SetFavorite ...
func (s *Server) SetFavorite(ctx context.Context, req *pb.SetFavoriteRequest) (*pb.SetFavoriteResponse, error) {
	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.FavoriteService.SetFavorite(token, req.BeerId)
	return &pb.SetFavoriteResponse{}, err
}

// GetFavorites ...
func (s *Server) GetFavorites(ctx context.Context, req *pb.GetFavoritesRequest) (*pb.GetFavoritesResponse, error) {
	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}
	result, err := s.FavoriteService.GetFavorites(token)
	return &pb.GetFavoritesResponse{Beers: result}, err
}

func (s *Server) DeleteFavorite(ctx context.Context, req *pb.DeleteFavoriteRequest) (*pb.DeleteFavoriteResponse, error) {
	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.FavoriteService.DeleteFavorite(token, req.BeerId)
	return &pb.DeleteFavoriteResponse{}, err
}
