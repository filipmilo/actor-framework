package remote

import (
	"context"
	"fmt"
	"main/actor"
	"main/proto"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Config struct {
	ListeningAddress string
	ServerOptions    []grpc.ServerOption
}

func NewConfig(address string, options ...grpc.ServerOption) *Config {
	return &Config{
		ListeningAddress: address,
		ServerOptions:    options,
	}
}

type Remote struct {
	system       *actor.ActorSystem
	config       *Config
	grpcServer   *grpc.Server
	remoteReader *RemoteReader
}

func NewRemote(system *actor.ActorSystem, config *Config) *Remote {
	return &Remote{
		system: system,
		config: config,
	}
}

func (r *Remote) Start() {
	ln, err := net.Listen("tcp", r.config.ListeningAddress)
	if err != nil {
		fmt.Println("Error starting server", err)
	}

	fmt.Println("Actor is listening on",
		ln.Addr(),
	)
	r.grpcServer = grpc.NewServer(r.config.ServerOptions...)
	proto.RegisterRemoteServer(r.grpcServer, NewRemoteReader(r))
	go r.grpcServer.Serve(ln)
}

func (r *Remote) SpawnRemotingActor(name, address string) uuid.UUID {
	response, _ := GetRemoteGrpcClient(address).GetRemotingActor(context.Background(), &proto.RemotingActorRequest{
		Name: name,
	})

	return uuid.MustParse(response.Pid)
}

func GetRemoteGrpcClient(address string) proto.RemoteClient {
	return proto.NewRemoteGrpcClient(address)
}
