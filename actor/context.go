package actor

import (
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
