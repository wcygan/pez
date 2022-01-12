/*
Copyright © 2022 William Cygan <wcygan.io@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package server

import (
	"context"
	"fmt"
	api "github.com/wcygan/pez/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("pez is listening on port %d\n", port)
	}

	dataService := DataService{
		UnimplementedDataServiceServer: api.UnimplementedDataServiceServer{},
		datastore:                      Datastore{backingStore: make(map[string]string)},
	}

	grpcServer := grpc.NewServer()

	api.RegisterDataServiceServer(grpcServer, dataService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type DataService struct {
	api.UnimplementedDataServiceServer
	datastore Datastore
}

func (s DataService) PutRecord(ctx context.Context, request *api.PutRequest) (*api.Record, error) {
	err := s.datastore.Put(request.Key, request.Value)
	if err != nil {
		return nil, err
	}

	return &api.Record{
		Key:   request.Key,
		Value: request.Value,
	}, nil
}

func (s DataService) GetRecord(ctx context.Context, request *api.GetRequest) (*api.Record, error) {
	value, err := s.datastore.Get(request.Key)
	if err != nil {
		return nil, err
	}

	return &api.Record{
		Key:   request.Key,
		Value: value,
	}, nil
}
