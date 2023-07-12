package actor

import (
	ctx "context"
	"main/proto"
)

type remotingActor struct {
	from                 string
	grpcClientConnection proto.RemoteClient
}

func RemotingActorProp(from string) IActor {
	return &remotingActor{
		from:                 from,
		grpcClientConnection: proto.NewRemoteGrpcClient(from),
	}
}

func (a *remotingActor) Recieve(context *ActorContext) {
	switch context.Message.message.(type) {
	case interface{}:
		envelope := proto.Encode(context.Message.message)
		envelope.Target = context.Message.reciver.String()
		a.grpcClientConnection.SendMessage(ctx.Background(), envelope)
	}
}
