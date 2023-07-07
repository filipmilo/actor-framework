package actor

import (
	"errors"
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

func (as *ActorSystem) ForwardMessage(message Envelope) {
  as.environment[message.reciver] <- message
}

func (as *ActorSystem) PrintValues() {
  fmt.Printf("Environment: %v\n", as.environment)
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

	// if !as.hasStarted {
	// 	return errors.New("actor system has not started yet")
	// }
	_, exist := as.environment[actorPid]
	if exist {

		delete(as.environment, actorPid)
		return fmt.Errorf("actor=%s is stopped", actorPid)
	}
	return fmt.Errorf("actor=%s not found in the system", actorPid)
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
}