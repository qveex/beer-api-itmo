package repository

import (
	"errors"
	pb "main/pkg/api"
)

// CreateNewBeer TODO domain entity && id
func (r *CatalogRepository) CreateNewBeer(beer *pb.Beer) (int64, error) {
	result := r.db.Create(&beer)

	if result.Error == nil {
		return -1, result.Error
	}

	var newBeer *pb.Beer
	r.db.Take(&newBeer)

	if newBeer == nil {
		return -1, errors.New("beer not created")
	}

	return newBeer.BeerId, nil
}

func (r *CatalogRepository) GetAllBeers() ([]*pb.Beer, error) {
	var beers []*pb.Beer
	r.db.Find(&beers)

	return beers, nil
}
