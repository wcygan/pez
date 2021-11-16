package pez

import (
	"context"
	api "github.com/wcygan/pez/api/v1"
)

type DataService struct {
	api.UnimplementedDataServiceServer
}

func (s DataService) PutRecord(ctx context.Context, request *api.PutRequest) (*api.Record, error) {
	return &api.Record{
		Key:   request.Key,
		Value: request.Value + "!!!",
	}, nil
}

func (s DataService) GetRecord(ctx context.Context, request *api.GetRequest) (*api.Record, error) {
	panic("implement me")
}
