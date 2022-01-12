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
package client

import (
	"context"
	"fmt"
	api "github.com/wcygan/pez/api/v1"
	"google.golang.org/grpc"
	"log"
)

func Put(key, value string, port int) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDataServiceClient(conn)

	response, err := c.PutRecord(context.Background(), &api.PutRequest{
		Key:   key,
		Value: value,
	})

	if err != nil {
		log.Fatalf("Error when calling PutRecord: %s", err)
	}
	log.Printf("Assigned (%s, %s)", response.Key, response.Value)
}

func Get(key string, port int) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDataServiceClient(conn)

	response, err := c.GetRecord(context.Background(), &api.GetRequest{
		Key: key,
	})

	if err != nil {
		log.Fatalf("Error when calling GetRecord: %s", err)
	}
	log.Printf("Retrieved (%s, %s)", response.Key, response.Value)
}
