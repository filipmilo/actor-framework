package main

import (
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

func (s *Sender) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.AdderMessage:
		msg.IsAdd = true
		context.Send(s.adderPid, msg)
		context.Become(s.Subtract)
	default:
		log.Printf("Ivalid message type")
	}
}

func (s *Sender) Subtract(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.AdderMessage:
		msg.IsAdd = false
		context.Send(s.adderPid, msg)

		context.Become(s.Recieve)
	default:
		log.Printf("Ivalid message type")
	}
}

func main() {
	system := actor.NewSystem()
	context := system.Root
	remote := remote.NewRemote(system, remote.NewConfig("192.168.1.105:8000"))
	remote.Start()

	adder, _ := remote.SpawnPid("Adder", "192.168.1.125:8080")
	sender := context.InitActor(&Sender{
		adderPid: *adder,
	}, "Sender")

	context.Send(*sender, &messages.AdderMessage{Amount: 6})
	context.Send(*sender, &messages.AdderMessage{Amount: 1})
	context.Send(*sender, &messages.AdderMessage{Amount: 8})
	context.Send(*sender, &messages.AdderMessage{Amount: 1})
	context.Send(*sender, &messages.AdderMessage{Amount: 4})
	context.Send(*sender, &messages.AdderMessage{Amount: 10})
	context.Send(*sender, &messages.AdderMessage{Amount: 89})
	context.Send(*sender, &messages.AdderMessage{Amount: 6})

	system.PrintValues()
	time.Sleep(60 * time.Second)

}
