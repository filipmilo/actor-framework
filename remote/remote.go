package remote

import (
	"context"
	"fmt"
	"main/actor"
	"main/proto"
	"net"
	"strings"

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

	fmt.Println("Remote server is starting on: ",
		ln.Addr(),
	)
	r.grpcServer = grpc.NewServer(r.config.ServerOptions...)
	proto.RegisterRemoteServer(r.grpcServer, NewRemoteReader(r))
	go r.grpcServer.Serve(ln)
}

func (r *Remote) SpawnPid(name, address string) (*uuid.UUID, error) {
	response, err := GetRemoteGrpcClient(address).GetRemotingActor(context.Background(), &proto.RemotingActorRequest{
		Name: name,
	})
	if err != nil {
		fmt.Print(strings.Split(err.Error(), "=")[2])
		return &uuid.Nil, err
	}

	return r.system.Root.InitRemoteActor(uuid.MustParse(response.Pid), address), nil
}

func GetRemoteGrpcClient(address string) proto.RemoteClient {
	return proto.NewRemoteGrpcClient(address)
}
