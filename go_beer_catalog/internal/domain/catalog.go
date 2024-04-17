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

func (s *CatalogService) GetAllBeers(limit *int32, name *string, brand *string, beerType *pb.Type, deg *int32, sweet *bool) ([]*pb.Beer, error) {
	var localLimit int
	var localName string
	var localBrand string
	var localType int

	if limit == nil {
		localLimit = 50
	} else {
		localLimit = int(*limit)
	}
	if name == nil {
		localName = ""
	} else {
		localName = *name
	}
	if brand == nil {
		localBrand = ""
	} else {
		localBrand = *brand
	}
	if beerType == nil {
		localType = -1
	} else {
		//localType = beerType.Descriptor().Index()
		switch *beerType {
		case pb.Type_ALE:
			localType = 0
		case pb.Type_LAGER:
			localType = 1
		case pb.Type_LAMBIC:
			localType = 2
		case pb.Type_STOUT:
			localType = 3
		case pb.Type_CRAFT:
			localType = 4
		case pb.Type_GAY:
			localType = 5
		default:
			localType = -1
		}
	}

	return s.repo.GetAllBeers(localLimit, localName, localBrand, localType, deg, sweet)
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
