package main

import (
	"fmt"
	"main/internal/controller"
	"main/internal/env"
	pr "main/internal/pkg/di"
	"main/internal/repository"
	pb "main/pkg/api"
	"net"
)

const (
	serverPort = "PORT"
)

/*func init() {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Fatalf("No .env file found: %s", err.Error())
	}
}*/

func main() {
	initDb()
	initServer()
}

func initServer() {
	log := pr.Provider.GetLogger()
	server := pr.Provider.GetCatalogServer()
	catalogService := pr.Provider.GetCatalogService()
	favoriteService := pr.Provider.GetFavoriteService()

	port := fmt.Sprintf(":%s", env.GetEnv(serverPort, "8080"))
	pb.RegisterCatalogServer(
		server,
		&controller.Server{
			CatalogService:  catalogService,
			FavoriteService: favoriteService,
		},
	)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Error(err.Error())
	}

	log.Debug("Listening on port " + port)
	if err := server.Serve(l); err != nil {
		log.Error(err.Error())
	}
}

func initDb() {
	db := pr.Provider.GetDb()
	err := db.AutoMigrate(&repository.Favorite{}, &repository.Beer{})
	if err != nil {
		panic(err)
	}
}
