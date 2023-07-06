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
	channel     chan Envelope
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]*actor)
	as.channel = make(chan Envelope, 1000)
	go as.doing()
}

func (as *ActorSystem) doing() {
	for {
		//Pa bih ovo onda da ceka na poruke i obradjuje ih sekvencionalno
		msg := <-as.channel
		fmt.Println(msg)
		switch msg.message.(type) {
		case CreateActor:
		//Kreiraj Aktora
		//OVO BI TREBALO DA RADI AL NISAM TESTIRAO!
		default:
			//Ako ne prepozna poruku da je recimo droppuje. Takodje mozemo interno
			//imati vise slucajeva, ili exposovati korisniku kroz Context
			//recimo Restart,Delete,SendMessage itd..
		}
	}
}

type CreateActor struct {
	props IActor
}

// S ovom metodom mogu svi da kreiraj
func (as *ActorSystem) CreateActor(parent *uuid.UUID, prop IActor) {
	as.channel <- Envelope{
		reciver: nil,
		sender:  parent,
		message: CreateActor{props: prop},
	}
}

func (as *ActorSystem) InitActor(prop IActor) *uuid.UUID {
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
	return &uuid
}

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}
