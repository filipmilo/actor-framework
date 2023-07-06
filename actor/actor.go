package actor

import (
	"fmt"

	"github.com/google/uuid"
)

// Top level interface that is used for constructing valid props
type IActor interface {
	Recieve(context Context)
}

type ActorStatus int8

const (
	ActorLiving ActorStatus = 1
	ActorEnd    ActorStatus = 2
)

type actor struct {
	pid      uuid.UUID
	parentCh chan Envelope
	channel  chan Envelope
	name     string
	status   ActorStatus
	behavior *behavior
	prop     *IActor
}

func (a *actor) birth() uuid.UUID {
	a.pid = uuid.New()
	a.name = fmt.Sprintf("%s:%s", "BasicActor", a.pid.String())
	a.channel = make(chan Envelope, 100)

	fmt.Printf("I, %s am BORN!\n", a.name)

	go a.live()
	return a.pid
}

func (a *actor) live() {
	defer a.kill()

	for a.status = ActorLiving; a.status == ActorLiving; {

		select {
		case msg := <-a.channel:

			a.behavior.run(Context{
				Pid:      &a.pid,
				Name:     a.name,
				behavior: a.behavior,
				Envelope: msg,
			})
		default:

			a.behavior.run(Context{
				Pid:      &a.pid,
				Name:     a.name,
				behavior: a.behavior,
			})
		}

		// if rand.Intn(100) > 90 {
		// 	a.status = ActorEnd
		// }
	}
}

func (a *actor) kill() {
	fmt.Printf("I,%s have died... ARGHHHH!\n", a.name)
}
