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

func (a *ActorSystem) StopActor(ctx context.Context, actorPid uuid.UUID) error {
	// // add a span context
	// ctx, span := telemetry.SpanContext(ctx, "StopActor")
	// defer span.End()
	if !a.hasStarted {
		return errors.New("actor system has not started yet")
	}
	// actorPath := NewPath(name, NewAddress(protocol, a.name, "", -1))
	// if a.remotingEnabled {
	// 	// get the path of the given actor
	// 	actorPath = NewPath(name, NewAddress(protocol, a.name, a.remotingHost, int(a.remotingPort)))
	// }
	// check whether the given actor already exist in the system or not
	retrivetActor, exist := a.environment[actorPid]
	// actor is found.
	if exist {
		// stop the given actor
		retrivetActor.kill()
		return fmt.Errorf("actor=%s should be stopped", retrivetActor.pid)
	}
	return fmt.Errorf("actor=%s not found in the system", retrivetActor.pid)
}
