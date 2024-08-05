package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	pb "github.com/qdrant/go-client/qdrant"
)

func exampleMain() {
	//update here
	addr := "xxx"

	// Set up a connection to the server.
	config := &tls.Config{}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(credentials.NewTLS(config)), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collectionsClient := pb.NewCollectionsClient(conn)

	// Contact the server and print out its response.
	r, err := collectionsClient.List(ctx, &pb.ListCollectionsRequest{})
	if err != nil {
		log.Fatalf("could not get collections: %v", err)
	}
	log.Printf("List of collections: %s", r.GetCollections())
}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	apiKey := "xxx"

	newCtx := metadata.AppendToOutgoingContext(ctx, "api-key", apiKey)

	return invoker(newCtx, method, req, reply, cc, opts...)
}

// js
// const client = new QdrantClient({
//     url: 'xxx',
//     apiKey: 'xxx',
// });
