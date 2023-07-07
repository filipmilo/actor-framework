package actor

import (
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]chan Envelope

	Root *RootActor
}


func NewSystem() *ActorSystem {
  as := ActorSystem{}
	as.environment = make(map[uuid.UUID]chan Envelope)
  as.Root = newRootActor(&as)

  return &as
}


func (as *ActorSystem) RegiserActor(newActor *CreateActorMessage) {
  as.environment[newActor.pid] = newActor.channel
}

func (as *ActorSystem) ForwardMessage(message Envelope, actor string) {
  fmt.Println(actor)
  as.environment[message.reciver] <- message
}

func (as *ActorSystem) PrintValues() {
  fmt.Printf("Environment: %v\n", as.environment)
}
