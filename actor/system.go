package actor

import (
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]*Actor
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]*Actor)
}

func (as *ActorSystem) InitActor(prop IActor) {
	a := Actor{
		prop:     &prop,
		behavior: initBehavior(prop.Recieve),
	}

	pid := a.birth()

	_, ok := as.environment[pid]
	if ok {
		a.status = ActorEnd
		return
	}

	as.environment[pid] = &a
}

func (as *ActorSystem) InitRemoteActor(prop IActor, address string) *Actor {
	a := Actor{
		prop:             &prop,
		behavior:         initBehavior(prop.Recieve),
		ListeningAddress: address,
	}

	pid := a.birth()

	_, ok := as.environment[pid]
	if ok {
		a.status = ActorEnd
		return nil
	}

	as.environment[pid] = &a
	return &a
}

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}
