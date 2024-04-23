package domain

import (
	"main/internal/repository"
	pb "main/pkg/api"
)

type FavoriteService struct {
	repo *repository.FavoriteRepository
	auth *repository.AuthRepository
}

func NewFavoriteService(repo *repository.FavoriteRepository, auth *repository.AuthRepository) *FavoriteService {
	return &FavoriteService{repo: repo, auth: auth}
}

func (s *FavoriteService) GetFavorites(token string) ([]*pb.Beer, error) {
	userId, err := s.auth.GetUserId(token)
	if err != nil {
		return nil, err
	}
	return s.repo.GetFavorites(userId)
}

func (s *FavoriteService) SetFavorite(token string, beerId int64) error {
	userId, err := s.auth.GetUserId(token)
	if err != nil {
		return err
	}
	return s.repo.SetFavorite(userId, beerId)
}

func (s *FavoriteService) DeleteFavorite(token string, beerId int64) error {
	userId, err := s.auth.GetUserId(token)
	if err != nil {
		return err
	}
	return s.repo.DeleteFavorite(userId, beerId)
}
