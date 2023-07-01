package actor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type IActor interface {
  Recieve(context Context)
}

type ActorStatus int8

const (
	ActorLiving ActorStatus = 1
	ActorEnd    ActorStatus = 2
)

type Actor struct {
	Pid    uuid.UUID
	Name   string
	Status ActorStatus
  Behavior *behavior
  prop *IActor
}

func (a *Actor) Birth() uuid.UUID {
	a.Pid = uuid.New()
	a.Name = fmt.Sprintf("%s%d", "BasicActor", a.Pid)

	fmt.Printf("I, %s am BORN!\n", a.Name)

	go a.Live()
	return a.Pid
}

func (a *Actor) Live() {
	defer a.Kill()

	for a.Status = ActorLiving; a.Status == ActorLiving; {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
    a.Behavior.run(Context{
      behavior: a.Behavior,
    })

		if rand.Intn(100) > 90 {
			a.Status = ActorEnd
		}

		fmt.Printf("%s waiting for message\n", a.Name)
	}
}

func (a *Actor) Kill() {
	fmt.Printf("I,%s have died... ARGHHHH!\n", a.Name)
}

