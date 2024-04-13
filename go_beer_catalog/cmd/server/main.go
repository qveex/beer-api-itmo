package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"main/internal/env"
	pr "main/internal/pkg"
	pb "main/pkg/api"
	"main/pkg/controller"
	"net"
)

func init() {
	if err := godotenv.Load(); err != nil {
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

	service := pr.Provider.GetCatalogService()

	pb.RegisterCatalogServer(s, &controller.Server{Service: service})
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
	err := db.AutoMigrate(&pb.Beer{})
	if err != nil {
		panic(err)
	}
}

const (
	serverPort = "PORT"
)
