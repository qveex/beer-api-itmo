package domain

import "main/internal/repository"

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s AuthService) GetUserId(token string) (int64, error) {
	return s.repo.GetUserId(token)
}

func (s AuthService) IsAdmin(userId int64) (bool, error) {
	return s.repo.IsAdmin(userId)
}
