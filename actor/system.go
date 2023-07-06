package actor

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]*actor

	hasStarted bool

	channel chan string
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]*actor)

	as.channel = make(chan string, 100)
}

func (as *ActorSystem) InitActor(prop IActor) {

	a := actor{
		prop:          &prop,
		behavior:      initBehavior(prop.Recieve),
		systemChannel: as.channel,
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

func recieveActorStatus(c chan string, actorPid uuid.UUID) {
	status := <-c
	fmt.Println(status)

	if status == "ActorEnd" {
		as := &ActorSystem{} // Create an instance of the ActorSystem struct
		err := as.StopActor(actorPid)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (as *ActorSystem) StopActor(actorPid uuid.UUID) error {

	if !as.hasStarted {
		return errors.New("actor system has not started yet")
	}
	retrivetActor, exist := as.environment[actorPid]
	if exist {

		delete(as.environment, actorPid)
		return fmt.Errorf("actor=%s is stopped", retrivetActor.pid)
	}
	return fmt.Errorf("actor=%s not found in the system", retrivetActor.pid)
}

// func (as *ActorSystem) RestartActor(ctx context.Context, actorPid uuid.UUID) error {

// 	if !as.hasStarted {
// 		return errors.New("actor system has not started yet")
// 	}
// 	retrivetActor, exist := as.environment[actorPid]
// 	// actor is found.
// 	if exist {

// 		return fmt.Errorf("actor=%s should be stopped", retrivetActor.pid)
// 	}
// 	return fmt.Errorf("actor=%s not found in the system", retrivetActor.pid)
// }
