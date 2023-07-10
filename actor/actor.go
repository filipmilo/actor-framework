package actor

import (
	"fmt"

	"github.com/google/uuid"
)

// Top level interface that is used for constructing valid props
type IActor interface {
	Recieve(context *ActorContext)
}

type ActorStatus int8

const (
	ActorStarting = iota
	ActorLiving
	ActorEnd
)

type actor struct {
	system   *ActorSystem
	pid      uuid.UUID
	channel  chan Envelope
	name     string
	status   ActorStatus
	behavior *behavior
	prop     *IActor
}

func (a *actor) birth() uuid.UUID {
	a.pid = uuid.New()
	a.setup()
	go a.live()
	return a.pid
}

func (a *actor) setup() {
	a.channel = make(chan Envelope, 100)

	fmt.Printf("I, %s am BORN!\n", a.name)

	a.onCreateSignal()
	msg := <-a.channel
	fmt.Printf("I, %s am %s\n", a.name, msg.message)
}

func (a *actor) onCreateSignal() {
	a.system.Root.in <- Envelope{
		reciver: a.system.Root.pid,
		sender:  a.pid,
		message: &CreateActorMessage{
			pid:     a.pid,
			channel: a.channel,
			name:    a.name,
		},
	}
}

func (a *actor) onDeleteSignal() {
	a.system.Root.in <- Envelope{
		reciver: a.system.Root.pid,
		sender:  a.pid,
		message: &DeleteActorMessage{
			pid:     a.pid,
			channel: a.channel,
		},
	}
}

func (a *actor) createContext(msg *Envelope) *ActorContext {
	return &ActorContext{
		system:   a.system,
		behavior: a.behavior,
		Pid:      a.pid,
		Name:     a.name,
		Message:  msg,
	}
}

func (a *actor) live() {
	defer a.kill()
	a.status = ActorLiving

	for a.status == ActorLiving {
		select {

		case msg := <-a.channel:
			a.behavior.run(a.createContext(&msg))

		}
	}

}

func (a *actor) kill() {

	a.onDeleteSignal()

	fmt.Printf("I,%s have died... ARGHHHH!\n", a.name)
}
