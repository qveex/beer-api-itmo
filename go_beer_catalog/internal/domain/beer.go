package domain

import pb "main/pkg/api"

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
