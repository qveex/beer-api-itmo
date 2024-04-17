package controller

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/domain"
	pb "main/pkg/api"
	"strings"
)

// Server ...
type Server struct {
	pb.UnimplementedCatalogServer
	CatalogService  *domain.CatalogService
	FavoriteService *domain.FavoriteService
}

func getToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no token")
	}
	authHeader := md.Get("Authorization")
	if authHeader == nil {
		return "", status.Error(codes.Unauthenticated, "no token")
	}
	auth := authHeader[0]
	if auth == "" {
		return "", status.Error(codes.Unauthenticated, "token is invalid")
	}
	parts := strings.Split(auth, "Bearer ")
	if len(parts) == 2 {
		return parts[1], nil
	} else {
		return "", status.Error(codes.Unauthenticated, "token is invalid")
	}
}
