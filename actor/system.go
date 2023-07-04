package actor

import (
	"fmt"

	"github.com/google/uuid"
)

type ActorSystem struct {
	//Ovo treba promeniti u kanale, Evenutalno dodati Root Actora! i ovde samo Listu Actora
	//Root Actor da cuva mapu[PID->Channel] ili da to bude ovde u ActorSystem-u? A da mu Root pristupa direktno?
	//Jer on to jedini direktno i vidi
	environment map[uuid.UUID]*actor
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]*actor)

}

func (as *ActorSystem) InitActor(prop IActor) *Pid {
	a := actor{
		prop:     &prop,
		behavior: initBehavior(prop.Recieve),
	}

	uuid := a.birth()

	_, ok := as.environment[uuid]
	if ok {
		a.status = ActorEnd
		return nil
	}

	as.environment[uuid] = &a
	return &Pid{
		Value: uuid,
		// actor:   &a
		channel: a.channel,
	}
}

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}
