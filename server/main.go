package main

import (
	api "github.com/wcygan/pez/api/v1"
	"github.com/wcygan/pez/server/pez"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dataService := pez.DataService{}

	grpcServer := grpc.NewServer()

	api.RegisterDataServiceServer(grpcServer, dataService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
