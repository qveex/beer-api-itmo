package domain

import (
	"main/internal/repository"
	pb "main/pkg/api"
)

type FavoriteService struct {
	repo *repository.FavoriteRepository
}

func NewFavoriteService(repo *repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (s *FavoriteService) GetFavorites(userId int64) ([]*pb.Beer, error) {
	return s.repo.GetFavorites(userId)
}

func (s *FavoriteService) SetFavorite(userId int64, beerId int64) error {
	return s.repo.SetFavorite(userId, beerId)
}

func (s *FavoriteService) DeleteFavorite(userId int64, beerId int64) error {
	return s.repo.DeleteFavorite(userId, beerId)
}
