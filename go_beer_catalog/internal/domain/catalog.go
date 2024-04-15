package domain

import (
	"main/internal/repository"
	pb "main/pkg/api"
)

type CatalogService struct {
	repo *repository.CatalogRepository
}

func NewCatalogService(repo *repository.CatalogRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

func (s *CatalogService) CreateNewBeer(beer *pb.Beer) (int64, error) {
	return s.repo.CreateNewBeer(beer)
}

func (s *CatalogService) GetAllBeers() ([]*pb.Beer, error) {
	return s.repo.GetAllBeers()
}

func (s *CatalogService) GetBeer(beerId int64) (*pb.Beer, error) {
	return s.repo.GetBeer(beerId)
}

func (s *CatalogService) UpdateBeer(beerId int64, beer *pb.Beer) (*pb.Beer, error) {
	return s.repo.UpdateBeer(beerId, beer)
}

func (s *CatalogService) DeleteBeer(beerId int64) (*pb.Beer, error) {
	return s.repo.DeleteBeer(beerId)
}