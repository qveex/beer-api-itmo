package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"main/internal/controller"
	"main/internal/env"
	pr "main/internal/pkg/di"
	"main/internal/repository"
	pb "main/pkg/api"
	"net"
)

func init() {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Fatalf("No .env file found: %s", err.Error())
	}
}

func main() {
	initDb()
	initServer()
}

func initServer() {

	port := fmt.Sprintf(":%s", env.GetEnv(serverPort, "8080"))
	s := grpc.NewServer()

	catalogService := pr.Provider.GetCatalogService()
	favoriteService := pr.Provider.GetFavoriteService()

	pb.RegisterCatalogServer(
		s,
		&controller.Server{
			CatalogService:  catalogService,
			FavoriteService: favoriteService,
		},
	)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port %s", port)
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initDb() {
	db := pr.Provider.GetDb()
	err := db.AutoMigrate(&pb.Beer{}, repository.Favorite{})
	if err != nil {
		panic(err)
	}
}

const (
	serverPort = "PORT"
)
