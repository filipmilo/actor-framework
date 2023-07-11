package proto

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRemoteGrpcClient(address string) RemoteClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to remote system: %v", err)
	}
	return NewRemoteClient(conn)
}
