package main

import (
	"fmt"
	"log"
	"main/actor"
	messages "main/messages/proto"
	"main/remote"
	"time"
)

type MyActor struct{}

func (t *MyActor) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.Ping:
		fmt.Println("My actor received message: ", msg.Message)
	default:
		log.Printf("Ivalid message type")
	}
}

func main() {

	system := actor.NewSystem()
	context := system.Root
	remote := remote.NewRemote(system, remote.NewConfig("127.0.0.1:4200"))
	remote.Start()
	context.InitActor(&MyActor{}, "MyActor")

	time.Sleep(60 * time.Second)
}
