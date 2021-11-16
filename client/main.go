package main

import (
	"context"
	api "github.com/wcygan/pez/api/v1"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDataServiceClient(conn)

	response, err := c.PutRecord(context.Background(), &api.PutRequest{
		Key:   "Hello",
		Value: "World",
	})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s -> %s", response.Key, response.Value)
}
