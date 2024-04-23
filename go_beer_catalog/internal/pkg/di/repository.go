package di

import (
	"main/internal/repository"
)

func (sp ServiceProvider) GetCatalogRepository() *repository.CatalogRepository {
	if sp.catalogRepo != nil {
		return sp.catalogRepo
	}

	sp.catalogRepo = repository.NewCatalogRepository(sp.GetDb())

	return sp.catalogRepo
}

func (sp ServiceProvider) GetAuthRepository() *repository.AuthRepository {
	if sp.authRepo != nil {
		return sp.authRepo
	}

	sp.authRepo = repository.NewAuthRepository(sp.GetAuthClient())

	return sp.authRepo
}

func (sp ServiceProvider) GetFavoriteRepository() *repository.FavoriteRepository {
	if sp.favRepo != nil {
		return sp.favRepo
	}

	sp.favRepo = repository.NewFavoriteRepository(sp.GetDb())

	return sp.favRepo
}
