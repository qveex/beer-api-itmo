package domain

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"main/internal/repository"
	pb "main/pkg/api"
)

type CatalogService struct {
	repo *repository.CatalogRepository
	auth *repository.AuthRepository
}

func NewCatalogService(repo *repository.CatalogRepository, auth *repository.AuthRepository) *CatalogService {
	return &CatalogService{repo: repo, auth: auth}
}

func (s *CatalogService) CreateNewBeer(token string, beer *pb.Beer) (int64, error) {
	isAdmin, err := s.isAdmin(token)
	if err != nil || !isAdmin {
		return -1, err
	}
	return s.repo.CreateNewBeer(beer)
}

func (s *CatalogService) GetAllBeers(
	limit *int32,
	name *string,
	brand *string,
	fromPrice *float64,
	toPrice *float64,
	beerType *pb.Type,
	deg *int32,
	sweet *bool,
) ([]*pb.Beer, error) {

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

	return s.repo.GetAllBeers(localLimit, localName, localBrand, fromPrice, toPrice, localType, deg, sweet)
}

func (s *CatalogService) GetBeer(beerId int64) (*pb.Beer, error) {
	return s.repo.GetBeer(beerId)
}

func (s *CatalogService) UpdateBeer(token string, beerId int64, beer *pb.Beer) (*pb.Beer, error) {
	isAdmin, err := s.isAdmin(token)
	if err != nil || !isAdmin {
		return nil, err
	}
	return s.repo.UpdateBeer(beerId, beer)
}

func (s *CatalogService) DeleteBeer(token string, beerId int64) (*pb.Beer, error) {
	isAdmin, err := s.isAdmin(token)
	if err != nil || !isAdmin {
		return nil, err
	}
	return s.repo.DeleteBeer(beerId)
}

func (s *CatalogService) isAdmin(token string) (bool, error) {
	userId, err := s.auth.GetUserId(token)
	if err != nil {
		return false, err
	}
	isAdmin, err := s.auth.IsAdmin(userId)

	if err != nil {
		return false, err
	}
	if isAdmin {
		return true, nil
	} else {
		return false, status.Error(codes.PermissionDenied, "user is not admin")
	}
}
