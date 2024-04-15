package repository

import (
	"gorm.io/gorm"
	pb "main/pkg/api"
)

type CatalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) *CatalogRepository {
	return &CatalogRepository{
		db: db,
	}
}

func (r *CatalogRepository) CreateNewBeer(beer *pb.Beer) (int64, error) {
	var lastBeer *pb.Beer
	r.db.Last(&lastBeer)
	beer.BeerId = lastBeer.BeerId + 1 // костылек, нужна подвязка к орм

	result := r.db.Create(&beer)

	return beer.BeerId, result.Error
}

func (r *CatalogRepository) GetAllBeers() ([]*pb.Beer, error) {
	var beers []*pb.Beer
	r.db.Find(&beers)

	return beers, nil
}

func (r *CatalogRepository) GetBeer(beerId int64) (*pb.Beer, error) {
	var beer *pb.Beer
	result := r.db.First(&beer, beerId)

	return beer, result.Error
}

func (r *CatalogRepository) UpdateBeer(beerId int64, beer *pb.Beer) (*pb.Beer, error) {
	var foundBeer *pb.Beer
	result := r.db.First(&foundBeer, beerId)

	if foundBeer == nil {
		return nil, result.Error
	}

	foundBeer = beer
	result = r.db.Save(foundBeer)

	return beer, result.Error
}

func (r *CatalogRepository) DeleteBeer(beerId int64) (*pb.Beer, error) {

	var beer *pb.Beer
	result := r.db.First(&beer, beerId)

	if beer == nil {
		return nil, result.Error
	}

	result = r.db.Delete(beer, beerId)

	return beer, result.Error
}
