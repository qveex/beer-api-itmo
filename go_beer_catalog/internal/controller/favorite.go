package controller

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	pb "main/pkg/api"
)

// SetFavorite ...
func (s *Server) SetFavorite(ctx context.Context, req *pb.SetFavoriteRequest) (*pb.SetFavoriteResponse, error) {
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	err = s.FavoriteService.SetFavorite(token, req.BeerId)
	return &pb.SetFavoriteResponse{}, err
}

// GetFavorites ...
func (s *Server) GetFavorites(ctx context.Context, req *pb.GetFavoritesRequest) (*pb.GetFavoritesResponse, error) {
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	result, err := s.FavoriteService.GetFavorites(token)
	return &pb.GetFavoritesResponse{Beers: result}, err
}

func (s *Server) DeleteFavorite(ctx context.Context, req *pb.DeleteFavoriteRequest) (*pb.DeleteFavoriteResponse, error) {
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	err = s.FavoriteService.DeleteFavorite(token, req.BeerId)
	return &pb.DeleteFavoriteResponse{}, err
}
