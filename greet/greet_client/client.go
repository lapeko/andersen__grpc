package main

import (
	"andersen__grpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	grpcClient, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer func(grpcClient *grpc.ClientConn) {
		err := grpcClient.Close()
		if err != nil {
			log.Fatalf("%v", err)
		}
	}(grpcClient)

	greetServiceClient := greetpb.NewGreetServiceClient(grpcClient)
	log.Printf("%v", greetServiceClient)
}
