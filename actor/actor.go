package actor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Top level interface that is used for constructing valid props
type IActor interface {
	Recieve(context Context)
}

type RemoteActor struct {
	actor   Actor
	address string
}

type ActorStatus int8

const (
	ActorLiving ActorStatus = 1
	ActorEnd    ActorStatus = 2
)

type Actor struct {
	pid              uuid.UUID
	name             string
	status           ActorStatus
	behavior         *behavior
	prop             *IActor
	ListeningAddress string
}

func (a *Actor) birth() uuid.UUID {
	a.pid = uuid.New()

	a.name = fmt.Sprintf("%s:%s", "BasicActor", a.pid.String())

	fmt.Printf("I, %s am BORN!\n", a.name)

	go a.live()
	return a.pid
}

func (a *Actor) live() {
	defer a.kill()

	for a.status = ActorLiving; a.status == ActorLiving; {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		a.behavior.run(Context{
			Name:     a.name,
			behavior: a.behavior,
		})

		if rand.Intn(100) > 90 {
			a.status = ActorEnd
		}

		fmt.Printf("%s waiting for message\n", a.name)
	}
}

func (a *Actor) kill() {
	fmt.Printf("I,%s have died... ARGHHHH!\n", a.name)
}
