package actor

import (
	ctx "context"
	"main/proto"

	"github.com/google/uuid"
)

type ActorContext struct {
	system   *ActorSystem
	behavior *behavior
	Pid      uuid.UUID
	Name     string
	Message  *Envelope
}

func (context *ActorContext) Become(newBehavior func(context *ActorContext)) {
	context.behavior.run = newBehavior
}

func (context *ActorContext) Send(reciever uuid.UUID, message IMessage) {
	envelope := Envelope{
		reciver: reciever,
		sender:  context.Pid,
		message: message,
	}

	context.system.ForwardMessage(envelope)
}

func (context *ActorContext) SendRemote(address string, reciever uuid.UUID, message IMessage) {
	envelope := Envelope{
		reciver: reciever,
		sender:  &context.Pid,
		message: message,
	}
	if reciever == uuid.Nil {
		return
	}
	msg := proto.Encode(envelope.message)
	msg.Target = reciever.String()

	go GetRemoteGrpcClient(address).SendMessage(ctx.Background(), msg)

}

func GetRemoteGrpcClient(address string) proto.RemoteClient {
	return proto.NewRemoteGrpcClient(address)
}
