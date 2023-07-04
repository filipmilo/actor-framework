package main

import (
	"fmt"
	"main/actor"
	"main/remote"
	"time"
)

type TestActor1 struct{}

func (t *TestActor1) Recieve(context actor.Context) {
	fmt.Printf("Hi im %s and this is my default state\n", context.Name)
	context.Become(t.SecondState)
}

func (t *TestActor1) SecondState(context actor.Context) {
	fmt.Printf("Hi im %s and this is my second state\n", context.Name)
	context.Become(t.Recieve)
}

type TestActor2 struct{}

func (t *TestActor2) Recieve(context actor.Context) {
	fmt.Printf("Hi im %s and this is my default state\n", context.Name)
	context.Become(t.SecondState)
}

func (t *TestActor2) SecondState(context actor.Context) {
	fmt.Printf("Hi im %s and this is my second state\n", context.Name)
	context.Become(t.Recieve)
}

type TestRemoteActor struct{}

func (t *TestRemoteActor) Recieve(context actor.Context) {
	fmt.Printf("Hi im %s and this is my default state\n", context.Name)
	context.Become(t.SecondState)
}

func (t *TestRemoteActor) SecondState(context actor.Context) {
	fmt.Printf("Hi im %s and this is my second state\n", context.Name)
	context.Become(t.Recieve)
}

type TestRemoteActor2 struct{}

func (t *TestRemoteActor2) Recieve(context actor.Context) {
	fmt.Printf("Hi im %s and this is my default state\n", context.Name)
	context.Become(t.SecondState)
}

func (t *TestRemoteActor2) SecondState(context actor.Context) {
	fmt.Printf("Hi im %s and this is my second state\n", context.Name)
	context.Become(t.Recieve)
}

func main() {
	system := actor.ActorSystem{}
	system.InitSystem()
	remoteActor := system.InitRemoteActor(&TestRemoteActor{}, "127.0.0.1:8000")
	remote.NewRemote(remoteActor, remote.NewConfig()).Start()
	remoteActor2 := system.InitRemoteActor(&TestRemoteActor2{}, "127.0.0.1:8001")
	remote.NewRemote(remoteActor2, remote.NewConfig()).Start()

	system.InitActor(&TestActor1{})
	system.InitActor(&TestActor2{})

	system.PrintValues()

	time.Sleep(60 * time.Second)
}
