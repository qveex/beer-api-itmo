package domain

import pb "main/pkg/api"

func (s *CatalogService) CreateNewBeer(beer *pb.Beer) (int64, error) {
	return s.repo.CreateNewBeer(beer)
}

func (s *CatalogService) GetAllBeers() ([]*pb.Beer, error) {
	return s.repo.GetAllBeers()
}
