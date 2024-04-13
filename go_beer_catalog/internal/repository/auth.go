package repository

import (
	"context"
	pb "main/pkg/api"
)

func (r *AuthRepository) GetUserId() (int64, error) {
	beer, err := r.api.GetBeer(context.Background(), &pb.GetBeerRequest{BeerId: 1})
	if err != nil {
		return -1, err
	}
	return beer.Beer.BeerId, nil
}
