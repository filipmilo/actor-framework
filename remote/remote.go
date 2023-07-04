package remote

import (
	"fmt"
	"main/actor"
	"net"

	"google.golang.org/grpc"
)

type Config struct {
	ServerOptions []grpc.ServerOption
}

func NewConfig(options ...grpc.ServerOption) *Config {
	return &Config{
		ServerOptions: options,
	}
}

type Remote struct {
	actor      *actor.Actor
	config     *Config
	grpcServer *grpc.Server
}

func NewRemote(actor *actor.Actor, config *Config) *Remote {
	return &Remote{
		actor:  actor,
		config: config,
	}
}

func (r *Remote) Start() {
	ln, err := net.Listen("tcp", r.actor.ListeningAddress)
	if err != nil {
		fmt.Println("Error starting server", err)
	}

	fmt.Println("Actor is listening on",
		ln.Addr(),
	)
	r.grpcServer = grpc.NewServer(r.config.ServerOptions...)
	go r.grpcServer.Serve(ln)
}
