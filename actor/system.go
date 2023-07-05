package actor

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	environment map[uuid.UUID]*actor

	hasStarted bool
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]*actor)
}

func (as *ActorSystem) InitActor(prop IActor) {
	a := actor{
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

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}

func (as *ActorSystem) StopActor(ctx context.Context, actorPid uuid.UUID) error {

	if !as.hasStarted {
		return errors.New("actor system has not started yet")
	}
	retrivetActor, exist := as.environment[actorPid]
	// actor is found.
	if exist {
		// stop the given actor
		retrivetActor.kill()
		return fmt.Errorf("actor=%s should be stopped", retrivetActor.pid)
	}
	return fmt.Errorf("actor=%s not found in the system", retrivetActor.pid)
}

func (as *ActorSystem) RestartActor(ctx context.Context, actorPid uuid.UUID) error {

	if !as.hasStarted {
		return errors.New("actor system has not started yet")
	}
	retrivetActor, exist := as.environment[actorPid]
	// actor is found.
	if exist {
		// stop the given actor
		retrivetActor.kill()
		retrivetActor.birth()
		return fmt.Errorf("actor=%s should be stopped", retrivetActor.pid)
	}
	return fmt.Errorf("actor=%s not found in the system", retrivetActor.pid)
}
