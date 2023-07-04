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
	parentCh chan IMessage
	channel  chan IMessage
	name     string
	status   ActorStatus
	behavior *behavior
	prop     *IActor
}

func (a *actor) birth() uuid.UUID {
	a.pid = uuid.New()
	a.name = fmt.Sprintf("%s:%s", "BasicActor", a.pid.String())
	a.channel = make(chan IMessage, 100)

	fmt.Printf("I, %s am BORN!\n", a.name)

	go a.live()
	return a.pid
}

func (a *actor) live() {
	defer a.kill()

	for a.status = ActorLiving; a.status == ActorLiving; {
		// fmt.Printf("ENTERED FOR, actor %s \n", a.name)

		select {
		case msg := <-a.channel:
			// fmt.Println("Got Message!")
			// fmt.Println(msg)

			// time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

			a.behavior.run(Context{
				Pid:      Pid{Value: a.pid, channel: a.channel},
				Name:     a.name,
				behavior: a.behavior,
				Message:  msg,
			})
		default:
			// fmt.Println("No Message! Do Usual Stuff")
			// time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

			a.behavior.run(Context{
				Pid:      Pid{Value: a.pid, channel: a.channel},
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
