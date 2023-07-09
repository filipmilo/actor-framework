package actor

import (
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]chan Envelope
	names       map[string]uuid.UUID
	Root        *RootActor
}

func NewSystem() *ActorSystem {
	as := ActorSystem{}
	as.environment = make(map[uuid.UUID]chan Envelope)
	as.Root = newRootActor(&as)

	return &as
}

func (as *ActorSystem) RegiserActor(newActor *CreateActorMessage) {
	as.environment[newActor.pid] = newActor.channel
	as.names[newActor.name] = newActor.pid
}

func (as *ActorSystem) ForwardMessage(message Envelope) {
	as.environment[message.reciver] <- message
}

func (as *ActorSystem) ActorPidByName(name string) uuid.UUID {
	return as.names[name]
}

func (as *ActorSystem) PrintValues() {
	fmt.Printf("Environment: %v\n", as.environment)
}

func (as *ActorSystem) StopActor(actorPid uuid.UUID) {

	// if !as.hasStarted {
	// 	return errors.New("actor system has not started yet")
	// }
	_, exist := as.environment[actorPid]
	if exist {

		delete(as.environment, actorPid)
		fmt.Printf("actor=%s is stopped\n", actorPid)
		return
	}
	fmt.Printf("actor=%s not found in the system\n", actorPid)
}
