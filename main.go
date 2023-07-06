package main

import (
	"main/actor"
	"time"
)

type Sender struct{}

func (t *Sender) Recieve(context actor.Context) {
}

type Adder struct {
	sum int
}

func (a *Adder) Recieve(context actor.Context) {
}

type ComplexValue struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	system := actor.ActorSystem{}
	system.InitSystem()

	// pid1 := system.InitActor(&Sender{})
	// pid2 := system.InitActor(&Adder{})

	//Ako je recimo parent==nil onda u logici kreiranja stavimo da mu je parent ActorSystem il sta ti ja znam
	//Ako imas drugaciju ideju slobodno ispravi jer s ovim mojim pristupom ne mogu da vratim PID do Maina za programera
	system.CreateActor(nil, &Sender{})

	system.PrintValues()

	// fmt.Println(pid1, pid2)

	time.Sleep(60 * time.Second)
}
