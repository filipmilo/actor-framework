package main

import (
	"fmt"
	"main/actor"
	"time"
)

type TestActor struct {}

func(t *TestActor) Recieve(context actor.Context) {
  fmt.Printf("This is my default state\n")
  context.Become(t.SecondState)
}


func(t *TestActor) SecondState(context actor.Context) {
  fmt.Printf("This is my second state\n")
  context.Become(t.Recieve)
}

func main() {
	system := actor.ActorSystem{}
	system.InitSystem()

	system.InitActor(&TestActor{})

	system.PrintValues()

	time.Sleep(60 * time.Second)
}
