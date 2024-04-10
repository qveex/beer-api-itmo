package main

import (
	"log"
	pb "main/pkg/api"
	"main/pkg/catalog"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	pb.RegisterCatalogServer(s, &catalog.Server{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
