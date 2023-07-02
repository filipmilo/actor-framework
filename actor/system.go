package actor

import (
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]*actor
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]*actor)
}

func (as *ActorSystem) InitActor(prop IActor) {
  a := actor {
    prop: &prop,
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

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}
