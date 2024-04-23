package domain

import (
	"main/internal/repository"
	pb "main/pkg/api"
)

func mapToDb(beer *pb.Beer) *repository.Beer {
	return &repository.Beer{
		BeerId: beer.BeerId,
		Name:   beer.Name,
		Brand:  beer.Brand,
		Price:  beer.Price,
		Type:   mapToDbType(&beer.Type),
		Degree: beer.Deg,
		Sweet:  beer.Sweet,
	}
}

func mapFromDb(beer *repository.Beer) *pb.Beer {
	return &pb.Beer{
		BeerId: beer.BeerId,
		Name:   beer.Name,
		Brand:  beer.Brand,
		Price:  beer.Price,
		Type:   mapFromDbType(beer.Type),
		Deg:    beer.Degree,
		Sweet:  beer.Sweet,
	}
}

func mapToDbType(beerType *pb.Type) int32 {
	var localType int32
	if beerType == nil {
		localType = -1
	} else {
		//localType = beerType.Descriptor().Index()
		switch *beerType {
		case pb.Type_ALE:
			localType = 0
		case pb.Type_LAGER:
			localType = 1
		case pb.Type_LAMBIC:
			localType = 2
		case pb.Type_STOUT:
			localType = 3
		case pb.Type_CRAFT:
			localType = 4
		case pb.Type_GAY:
			localType = 5
		default:
			localType = -1
		}
	}
	return localType
}

func mapFromDbType(beerType int32) pb.Type {
	var localType pb.Type
	switch beerType {
	case 0:
		localType = pb.Type_ALE
	case 1:
		localType = pb.Type_LAGER
	case 2:
		localType = pb.Type_LAMBIC
	case 3:
		localType = pb.Type_STOUT
	case 4:
		localType = pb.Type_CRAFT
	case 5:
		localType = pb.Type_GAY
	default:
		localType = pb.Type_GAY
	}
	return localType
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
