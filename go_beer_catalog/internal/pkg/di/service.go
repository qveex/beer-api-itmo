package di

import "main/internal/domain"

func (sp *ServiceProvider) GetCatalogService() *domain.CatalogService {
	if sp.catalogService != nil {
		return sp.catalogService
	}

	sp.catalogService = domain.NewCatalogService(sp.GetCatalogRepository())

	return sp.catalogService
}

func (sp *ServiceProvider) GetFavoriteService() *domain.FavoriteService {
	if sp.favService != nil {
		return sp.favService
	}

	sp.favService = domain.NewFavoriteService(sp.GetFavoriteRepository())

	return sp.favService
}

func (sp *ServiceProvider) GetAuthService() *domain.AuthService {
	if sp.catalogService != nil {
		return sp.authService
	}

	sp.authService = domain.NewAuthService(sp.GetAuthRepository())

	return sp.authService
}
