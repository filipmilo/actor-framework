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
  a := Actor {
    prop: &prop,
    Behavior: initBehavior(prop.Recieve),
  }

	pid := a.Birth()

	_, ok := as.environment[pid]
	if ok {
		a.Status = ActorEnd
		return
	}

	as.environment[pid] = &a
}

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}
