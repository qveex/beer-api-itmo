package repository

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	pb "main/pkg/api"
)

type Favorite struct {
	UserId int64 `gorm:"primaryKey;autoIncrement:false"`
	BeerId int64 `gorm:"primaryKey;autoIncrement:false"`
	Beer   *Beer `gorm:"foreignKey:BeerId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{
		db: db,
	}
}

func (r *FavoriteRepository) GetFavorites(userId int64) ([]*pb.Beer, error) {
	var beers []*pb.Beer

	r.db.
		Table("favorites").
		Where("user_id = ?", userId).
		Joins("left join beers on beers.beer_id = favorites.beer_id").
		Scan(&beers)

	return beers, nil
}

func (r *FavoriteRepository) SetFavorite(userId int64, beerId int64) error {
	var exists bool

	r.db.Raw("select true from favorites where user_id = ? and beer_id = ?", userId, beerId).Scan(&exists)
	if exists {
		return status.Error(codes.AlreadyExists, "favorite already exist")
	}

	r.db.Raw("select true from beers where beer_id = ?", beerId).Scan(&exists)
	if !exists {
		return status.Error(codes.NotFound, "beer not found")
	}

	result := r.db.Create(&Favorite{UserId: userId, BeerId: beerId})

	return result.Error
}

func (r *FavoriteRepository) DeleteFavorite(userId int64, beerId int64) error {
	var exists bool
	result := r.db.Raw("select true from favorites where user_id = ? and beer_id = ?", userId, beerId).Scan(&exists)

	if !exists {
		return status.Error(codes.NotFound, "favorite doesn't exist")
	}

	result = r.db.Delete(&Favorite{UserId: userId, BeerId: beerId})

	return result.Error
}
