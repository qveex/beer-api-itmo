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

func (r *AuthRepository) GetUserId(token string) (int64, error) {
	info, err := r.api.GetUserInfo(context.Background(), &ssov1.GetUserInfoRequest{Token: token})
	if err != nil {
		return -1, err
	}
	return info.UserId, nil
}

func (r *AuthRepository) IsAdmin(userId int64) (bool, error) {
	response, err := r.api.IsAdmin(context.Background(), &ssov1.IsAdminRequest{UserId: userId})
	if err != nil {
		return false, err
	}
	return response.IsAdmin, nil
}
