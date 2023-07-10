package actor

import (
<<<<<<< HEAD
=======
	ctx "context"
	"main/proto"

>>>>>>> 768d97b (small changes)
	"github.com/google/uuid"
)

type RootActor struct {
	pid    uuid.UUID
	system *ActorSystem
	in     chan Envelope
}

type CreateActorMessage struct {
	pid     uuid.UUID
	channel chan Envelope
	name    string
}

type DeleteActorMessage struct {
	pid     uuid.UUID
	channel chan Envelope
}

func newRootActor(system *ActorSystem) *RootActor {
	root := &RootActor{
		pid:    uuid.New(),
		system: system,
		in:     make(chan Envelope, 1000),
	}

	go root.events()
	return root
}

func (as *RootActor) events() {
	for {
		msg := <-as.in
		switch msg.message.(type) {
		case *CreateActorMessage:
			as.system.RegiserActor((msg.message).(*CreateActorMessage))
			as.system.ForwardMessage(Envelope{
				reciver: msg.sender,
				sender:  msg.reciver,
				message: "Successfully Registered!",
			})
		case *DeleteActorMessage:
			help := msg.message.(*DeleteActorMessage)
			as.system.StopActor(help.pid)
		default:
		}
	}
}

func (as *RootActor) InitActor(prop IActor, name string) *uuid.UUID {
	a := actor{
		system:   as.system,
		prop:     &prop,
		behavior: initBehavior(prop.Recieve),
		status:   ActorStarting,
		name:     name,
	}

	uuid := a.birth()
	return &uuid
}

func (as *RootActor) Send(reciever uuid.UUID, message IMessage) {
	envelope := Envelope{
		reciver: reciever,
		sender:  as.pid,
		message: message,
	}

	as.system.ForwardMessage(envelope)
}

func (as *RootActor) InitRemotingActor(prop IActor, name string, pid uuid.UUID) {
	a := actor{
		system:   as.system,
		prop:     &prop,
		behavior: initBehavior(prop.Recieve),
		status:   ActorStarting,
		name:     name,
	}

	a.birthRemote(pid)
}
