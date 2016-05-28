package main

import (
	"crypto/tls"
	"log"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func BenchmarkGRPCSetInfo(b *testing.B) {

	config := &tls.Config{}
	config.InsecureSkipVerify = true
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:4443", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewInfoServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		client.SetInfo(context.Background(), &InfoRequest{
			Name:   "test",
			Age:    1,
			Height: 1,
		})
	}

}
