package config

import (
	"fmt"
	"log"
	"main/prototypes/gen/go/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GrpcClient pb.IEventControllerClient

func GrpcConnect() {
	connection, err := grpc.Dial(EnvirontmentVariables.EventGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Printf("connection: %v\n", connection)

	GrpcClient = pb.NewIEventControllerClient(connection)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
}
