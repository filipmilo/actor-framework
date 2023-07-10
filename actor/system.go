package actor

import (
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]Props
	Root        *RootActor
	Remoter     *Remoter
}

type Props struct {
	name    string
	channel chan Envelope
}

type Remoter struct {
	address string
}

func NewSystem() *ActorSystem {
	as := ActorSystem{}
	as.environment = make(map[uuid.UUID]Props)
	as.Root = newRootActor(&as)

	return &as
}

func (as *ActorSystem) WithRemoting(address string) {
	as.Remoter = &Remoter{address: address}
}

func (as *ActorSystem) RegiserActor(newActor *CreateActorMessage) {
	as.environment[newActor.pid] = Props{name: newActor.name, channel: newActor.channel}
}

func (as *ActorSystem) ForwardMessage(message Envelope) {
	as.environment[message.reciver].channel <- message
}

func (as *ActorSystem) ActorPidByName(name string) uuid.UUID {
	for uuid, props := range as.environment {
		if props.name == name {
			return uuid
		}
	}
	return uuid.Nil
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
