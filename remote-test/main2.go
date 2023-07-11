package main

import (
	"fmt"
	"log"
	"main/actor"
	messages "main/messages/proto"
	"main/remote"
	"time"
)

type MyActor2 struct{}

func (t *MyActor2) Recieve(context *actor.ActorContext) {
	switch msg := context.Message.Message().(type) {
	case *messages.Ping:
		fmt.Println("My actor2 received message: ", msg.Message)
	default:
		log.Printf("Ivalid message type")
	}
}

func main() {

	system := actor.NewSystem()
	context := system.Root
	remote := remote.NewRemote(system, remote.NewConfig("127.0.0.1:4200"))
	remote.Start()
	context.InitActor(&MyActor2{}, "MyActor2")
	myActor1Pid, _ := remote.SpawnPid("MyActor1", "127.0.0.1:8000")
	fmt.Print(myActor1Pid)
	time.Sleep(3 * time.Second)
	context.Send(*myActor1Pid, &messages.Ping{Message: "Hello from remote system"})
	time.Sleep(60 * time.Second)
}
