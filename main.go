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
}

type RemoteMessage struct {
	message string
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

type MyActor1 struct {
	myActor2Pid uuid.UUID
}

func (t *MyActor1) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.Ping:
		fmt.Println("My actor1 received message: ", msg.Message)
		context.Send(t.myActor2Pid, msg)
	default:
		log.Printf("Ivalid message type")
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
	remote := remote.NewRemote(system, remote.NewConfig("127.0.0.1:8000"))
	remote.Start()
	adder := context.InitActor(&Adder{}, "Adder")
	sender := context.InitActor(&Sender{
		adderPid: *adder,
	}, "Sender")
	myActor2Pid, _ := remote.SpawnPid("MyActor2", "127.0.0.1:4200")
	context.InitActor(&MyActor1{
		myActor2Pid: *myActor2Pid,
	}, "MyActor1")

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
