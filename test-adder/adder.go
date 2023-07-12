package main

import (
	"fmt"
	"log"
	"main/actor"
	messages "main/messages/proto"
	"main/remote"
	"time"
)

type Adder struct {
	sum int32
}

func (a *Adder) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.AdderMessage:
		if msg.IsAdd {
			a.sum += msg.Amount
		} else {
			a.sum -= msg.Amount
		}
		fmt.Printf("Current sum is: %d\n", a.sum)
	default:
		log.Printf("Ivalid message type")
	}
}

func main() {

	system := actor.NewSystem()
	context := system.Root
	remote := remote.NewRemote(system, remote.NewConfig("192.168.1.125:8080"))
	remote.Start()

	context.InitActor(&Adder{}, "Adder")

	time.Sleep(60 * time.Second)
}
