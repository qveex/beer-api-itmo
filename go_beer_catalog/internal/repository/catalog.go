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

func (r *CatalogRepository) GetAllBeers(
	limit int,
	name string,
	brand string,
	fromPrice *float64,
	toPrice *float64,
	beerType int,
	deg *int32,
	sweet *bool,
) ([]*pb.Beer, error) {
	var beers []*pb.Beer

	result := r.db.
		Limit(limit).
		Table("beers").
		Where("name like ? and brand like ?", "%"+name+"%", "%"+brand+"%").
		Find(&beers)

	if beerType != -1 {
		result = result.Where("type = ?", beerType)
	}

	if deg != nil {
		result = result.Where("deg = ?", *deg)
	}

	if sweet != nil {
		result = result.Where("sweet = ?", *sweet)
	}

	if fromPrice != nil {
		result = result.Where("price >= ?", *fromPrice)
	}

	if toPrice != nil {
		result = result.Where("price <= ?", *toPrice)
	}

	result.Find(&beers)

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

	if foundBeer == nil || result.Error != nil {
		return nil, result.Error
	}

	beer.BeerId = foundBeer.BeerId
	foundBeer = beer
	result = r.db.Where("beer_id = ?", beerId).Updates(foundBeer)

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
