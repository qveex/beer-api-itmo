package repository

import (
	"context"
	ssov1 "github.com/Krutov777/protos/gen/go/sso"
)

type AuthRepository struct {
	api ssov1.AuthClient
}

func NewAuthRepository(api *ssov1.AuthClient) *AuthRepository {
	return &AuthRepository{
		api: *api,
	}
}

func (r *AuthRepository) GetUserId() (int64, error) {
	/*beer, err := r.api.GetBeer(context.Background(), &pb.GetBeerRequest{BeerId: 1})
	if err != nil {
		return -1, err
	}
	return beer.Beer.BeerId, nil*/
	return 0, nil
}

func (r *AuthRepository) IsAdmin(userId int64) (bool, error) {
	response, err := r.api.IsAdmin(context.Background(), &ssov1.IsAdminRequest{UserId: userId})
	if err != nil {
		return false, err
	}
	return response.IsAdmin, nil
}
