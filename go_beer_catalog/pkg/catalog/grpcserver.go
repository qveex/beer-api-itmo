package catalog

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "main/pkg/api"
	"slices"
)

// Server ...
type Server struct {
	pb.UnimplementedCatalogServer
}

var beers = []*pb.Beer{
	{BeerId: 1, Name: "Baltika 9", Brand: "Baltika", Type: pb.Type_STOUT, Deg: 9, Sweet: false},
	{BeerId: 2, Name: "Baltika 8", Brand: "Baltika", Type: pb.Type_STOUT, Deg: 8, Sweet: false},
	{BeerId: 3, Name: "Baltika 7", Brand: "Baltika", Type: pb.Type_STOUT, Deg: 7, Sweet: false},
}

// GetBeers ...
func (s *Server) GetBeers(ctx context.Context, req *pb.GetBeersRequest) (*pb.GetBeersResponse, error) {
	/*tmp, _ := db.GetDB()
	var tmpBeers []*pb.Beer
	tmp.Find(&tmpBeers)

	return &pb.GetBeersResponse{Beers: tmpBeers}, nil*/
	return &pb.GetBeersResponse{Beers: beers}, nil
}

// GetBeer ...
func (s *Server) GetBeer(ctx context.Context, req *pb.GetBeerRequest) (*pb.GetBeerResponse, error) {
	beerIndex := slices.IndexFunc(beers, func(beer *pb.Beer) bool {
		return req.BeerId == beer.BeerId
	})

	if beerIndex != -1 {
		return &pb.GetBeerResponse{Beer: beers[beerIndex]}, nil
	} else {
		return nil, status.Error(codes.NotFound, "Beer not found")
	}
}

// CreateBeer ...
func (s *Server) CreateBeer(ctx context.Context, req *pb.CreateBeerRequest) (*pb.CreateBeerResponse, error) {
	beers = append(beers, req.Beer)
	return &pb.CreateBeerResponse{}, nil
}

// UpdateBeer ...
func (s *Server) UpdateBeer(ctx context.Context, req *pb.UpdateBeerRequest) (*pb.UpdateBeerResponse, error) {
	found := false
	for i := range beers {
		if beers[i].BeerId == req.BeerId {
			beers[i].Name = req.Beer.Name
			beers[i].Brand = req.Beer.Brand
			beers[i].Type = req.Beer.Type
			beers[i].Deg = req.Beer.Deg
			beers[i].Sweet = req.Beer.Sweet

			found = true
			break
		}
	}

	if found == true {
		return &pb.UpdateBeerResponse{}, nil
	} else {
		return nil, status.Error(codes.NotFound, "Beer not found")
	}
}

// DeleteBeer ...
func (s *Server) DeleteBeer(ctx context.Context, req *pb.DeleteBeerRequest) (*pb.DeleteBeerResponse, error) {

	initSize := len(beers)
	beers = slices.DeleteFunc(beers, func(beer *pb.Beer) bool {
		return beer.BeerId == req.BeerId
	})
	newSize := len(beers)

	if newSize != initSize {
		return &pb.DeleteBeerResponse{}, nil
	} else {
		return nil, status.Error(codes.NotFound, "Beer not found")
	}
}
