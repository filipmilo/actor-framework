package main

import (
	"fmt"
	"main/actor"
	"time"
)

type TestActor1 struct {}

func(t *TestActor1) Recieve(context actor.Context) {
  fmt.Printf("Hi im %s and this is my default state\n", context.Name)
  context.Become(t.SecondState)
}


func(t *TestActor1) SecondState(context actor.Context) {
  fmt.Printf("Hi im %s and this is my second state\n", context.Name)
  context.Become(t.Recieve)
}

type TestActor2 struct {}

func(t *TestActor2) Recieve(context actor.Context) {
  fmt.Printf("Hi im %s and this is my default state\n", context.Name)
  context.Become(t.SecondState)
}


func(t *TestActor2) SecondState(context actor.Context) {
  fmt.Printf("Hi im %s and this is my second state\n", context.Name)
  context.Become(t.Recieve)
}

func main() {
	system := actor.ActorSystem{}
	system.InitSystem()

	system.InitActor(&TestActor1{})
	system.InitActor(&TestActor2{})

	system.PrintValues()

	time.Sleep(60 * time.Second)
}
