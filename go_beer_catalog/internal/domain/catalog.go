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
	return s.repo.CreateNewBeer(mapToDb(beer))
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

	result, err := s.repo.GetAllBeers(localLimit, localName, localBrand, fromPrice, toPrice, mapToDbType(beerType), deg, sweet)
	return Map(result, mapFromDb), err
}

func (s *CatalogService) GetBeer(beerId int64) (*pb.Beer, error) {
	result, err := s.repo.GetBeer(beerId)
	return mapFromDb(result), err
}

func (s *CatalogService) UpdateBeer(token string, beerId int64, beer *pb.Beer) (*pb.Beer, error) {
	isAdmin, err := s.isAdmin(token)
	if err != nil || !isAdmin {
		return nil, err
	}
	result, err := s.repo.UpdateBeer(beerId, mapToDb(beer))
	return mapFromDb(result), err
}

func (s *CatalogService) DeleteBeer(token string, beerId int64) (*pb.Beer, error) {
	isAdmin, err := s.isAdmin(token)
	if err != nil || !isAdmin {
		return nil, err
	}
	result, err := s.repo.DeleteBeer(beerId)
	return mapFromDb(result), err
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
