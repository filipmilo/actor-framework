package main

import (
	"fmt"
	"main/actor"
	"time"
)

func main() {
	fmt.Println("Hello World")

	system := actor.ActorSystem{}
	system.InitSystem()

	for i := 0; i < 3; i++ {
		system.InitActor()
	}

	system.PrintValues()

	time.Sleep(60 * time.Second)
	fmt.Println("FINISHED")
}
