package repository

import (
	"main/tests/mock"
	"testing"
)

func TestCreateBeer(t *testing.T) {
	repo := mock.Provider.GetCatalogRepository()

	result, err := repo.GetAllBeers(100, "", "", nil, nil, -1, nil, nil)
	if err != nil {
		t.Error(err)
	}

	lenBefore := len(result)

	_, err = repo.CreateNewBeer(mock.Beers[0])
	if err != nil {
		t.Error(err)
	}

	result, err = repo.GetAllBeers(100, "", "", nil, nil, -1, nil, nil)

	if err != nil {
		t.Error(err)
	}

	if lenBefore+1 != len(result) {
		t.Error("result len should be more than init")
	}
}

func TestGetBeerById(t *testing.T) {
	repo := mock.Provider.GetCatalogRepository()

	beerId, err := repo.CreateNewBeer(mock.Beers[0])
	if err != nil {
		t.Error(err)
	}

	beer, err := repo.GetBeer(beerId)

	if err != nil {
		t.Error(err)
	}

	if beer.Name != mock.Beers[0].Name {
		t.Error("beer should be mock beer")
	}
}

func TestDeleteBeer(t *testing.T) {
	repo := mock.Provider.GetCatalogRepository()

	beerId, err := repo.CreateNewBeer(mock.Beers[0])
	if err != nil {
		t.Error(err)
	}

	beer, err := repo.GetBeer(beerId)

	if err != nil {
		t.Error(err)
	}

	if beer == nil {
		t.Error("beer should not be nil")
	}

	beer, err = repo.DeleteBeer(beerId)

	if err != nil {
		t.Error(err)
	}

	beer, err = repo.GetBeer(beerId)

	if err == nil {
		t.Error("beer should not be found")
	}

}
