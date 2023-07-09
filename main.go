package main

import (
	"fmt"
	"log"
	"main/actor"
	messages "main/messages/proto"
	"main/remote"
	"time"

	"github.com/google/uuid"
)

type Sender struct {
	adderPid uuid.UUID
}

type AdderMessage struct {
	add    bool
	amount int32
}

type SenderMessage struct {
	amount int32
<<<<<<< HEAD
=======
}

type RemoteMessage struct {
	message string
>>>>>>> c812ca6 (Implemented remote sending messages)
}

func (t *Sender) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case SenderMessage:
		context.Send(t.adderPid, AdderMessage{
			add:    true,
			amount: msg.amount,
		})

		context.Become(t.Subtract)
	default:
		log.Printf("Ivalid message type")
	}
}

func (t *Sender) Subtract(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case SenderMessage:
		context.Send(t.adderPid, AdderMessage{
			add:    false,
			amount: msg.amount,
		})

		context.Become(t.Recieve)
	default:
		log.Printf("Ivalid message type")
	}
}

type Adder struct {
	sum int32
}

type RemoteActor struct {
}

func (r *RemoteActor) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.Ping:
		fmt.Printf("Remote message recieved: %v\n", msg.Message)
	default:
		fmt.Printf("Invalid message type")
	}
}

func (a *Adder) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case AdderMessage:
		if msg.add {
			a.sum += msg.amount
		} else {
			a.sum -= msg.amount
		}
		fmt.Printf("Current sum is: %d\n", a.sum)
	default:
		fmt.Printf("Invalid message type")
	}
}

type ComplexValue struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	system := actor.NewSystem()
	context := system.Root

	adder := context.InitActor(&Adder{}, "Adder")
	sender := context.InitActor(&Sender{
		adderPid: *adder,
	}, "Sender")

	//If they are not initialized by this point it will throw or wont work

	context.Send(*sender, SenderMessage{amount: 6})
	context.Send(*sender, SenderMessage{amount: 1})
	context.Send(*sender, SenderMessage{amount: 8})
	context.Send(*sender, SenderMessage{amount: 1})
	context.Send(*sender, SenderMessage{amount: 4})
	context.Send(*sender, SenderMessage{amount: 10})
	context.Send(*sender, SenderMessage{amount: 89})
	context.Send(*sender, SenderMessage{amount: 6})

	system.PrintValues()
	time.Sleep(60 * time.Second)
	system.PrintValues()

}
