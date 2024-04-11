package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"main/cmd/db"
	pb "main/pkg/api"
	"main/pkg/catalog"
	"net"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	//initDb()
	initServer()
}

func initServer() {
	s := grpc.NewServer()

	pb.RegisterCatalogServer(s, &catalog.Server{})

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initDb() {
	_, err := db.GetDB()
	if err != nil {
		return
	}
}
